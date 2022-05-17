package concurrenthash

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"hash"
	"os"

	"golang.org/x/sync/errgroup"
)

type block struct {
	Index int
	Data  []byte
}
type sum struct {
	Index int
	Hash  []byte
}

// ConcurrentHash is basically a https://en.wikipedia.org/wiki/Merkle_tree
type ConcurrentHash struct {
	Concurrency     int
	BlockSize       int64
	HashConstructor func() hash.Hash

	// internal
	Context context.Context
	Cancel  context.CancelFunc
	Hashes  [][]byte
}

// NewConcurrentHash is the constructor and entrypoint
func NewConcurrentHash(concurrency int, blockSize int64, hashFunc func() hash.Hash) *ConcurrentHash {
	var ctx, cancel = context.WithCancel(context.Background())

	return &ConcurrentHash{
		Concurrency:     concurrency,
		BlockSize:       blockSize,
		HashConstructor: hashFunc,
		Context:         ctx,
		Cancel:          cancel,
	}
}

// HashFile is a coordination func that fans out to hash workers,
// collects their output and hashes the final array
func (c *ConcurrentHash) HashFile(file string) (string, error) {

	// make sure the file even exists first
	var stat, err = os.Stat(file)
	if err != nil {
		c.Cancel()
		return "", err
	}
	c.Hashes = make([][]byte, (stat.Size()+c.BlockSize-1)/c.BlockSize)

	// startup all our routines, readers first
	var sumChans = make([]chan sum, c.Concurrency)
	var blocks = make(chan block)
	var errGroup = new(errgroup.Group)

	for i := 0; i < c.Concurrency; i++ {
		var sums = make(chan sum)
		sumChans[i] = sums
		errGroup.Go(func() error {
			return c.hashBlock(blocks, sums)
		})
	}
	errGroup.Go(func() error {
		c.collectSums(mergeSums(sumChans...))
		return nil
	})
	errGroup.Go(func() error {
		return c.streamFile(file, blocks)
	})

	// wait for all ^^ to finish
	if err := errGroup.Wait(); err != nil {
		return "", err
	}

	// hash the hashes
	var buf bytes.Buffer
	var enc = gob.NewEncoder(&buf)

	if err := enc.Encode(c.Hashes); err != nil {
		c.Cancel()
		return "", err
	}

	var h = c.HashConstructor()
	h.Reset()
	h.Write(buf.Bytes())
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
