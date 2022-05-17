package concurrenthash

import (
	"bufio"
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"hash"
	"io"
	"os"

	"github.com/twmb/murmur3"
	"golang.org/x/sync/errgroup"
)

type block struct {
	Index int
	Data  []byte
}
type sum struct {
	Index int
	Hash  uint64
}

// ConcurrentHash is basically a https://en.wikipedia.org/wiki/Merkle_tree
type ConcurrentHash struct {
	Concurrency int
	BlockSize   int64
	HashFunc    hash.Hash64

	// internal
	Context context.Context
	Cancel  context.CancelFunc
	Hashes  []uint64
}

// NewConcurrentHash is the constructor and entrypoint
func NewConcurrentHash(concurrency int, blockSize int64, hashFunc hash.Hash64) *ConcurrentHash {
	var ctx, cancel = context.WithCancel(context.Background())

	return &ConcurrentHash{
		Concurrency: concurrency,
		BlockSize:   blockSize,
		HashFunc:    hashFunc,
		Context:     ctx,
		Cancel:      cancel,
	}
}

// HashFile is a coordination func that fans out to hash workers,
// collects their output and hashes the final array
func (c *ConcurrentHash) HashFile(file string) (string, error) {
	var stat, err = os.Stat(file)
	if err != nil {
		c.Cancel()
		return "", err
	}
	c.Hashes = make([]uint64, (stat.Size()+c.BlockSize-1)/c.BlockSize)

	var sums = make(chan sum)
	var blocks = make(chan block)
	var errGroup = new(errgroup.Group)

	go c.collectSums(sums)
	for i := 0; i < c.Concurrency; i++ {
		errGroup.Go(func() error {
			return c.hashBlock(blocks, sums)
		})
	}
	errGroup.Go(func() error {
		return c.streamFile(file, blocks)
	})

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

	var h64 hash.Hash64 = murmur3.New64()
	_, err = h64.Write(buf.Bytes())
	if err != nil {
		return "", err
	}
	return fmt.Sprint(h64.Sum64()), nil
}

// collectSums is a fan in func to get the hashes and write them to an array
func (c *ConcurrentHash) collectSums(sums <-chan sum) {
	for {
		select {
		case <-c.Context.Done():
			return
		default:
			select {
			case sum, open := <-sums:
				if !open {
					return
				}
				c.Hashes[sum.Index] = sum.Hash
			default:
			}
		}
	}
}

// streamFile reads the file in blocks given a block size in ConcurrentHash and
// writes them to a given channel: blocks
func (c *ConcurrentHash) streamFile(filePath string, blocks chan<- block) error {
	defer close(blocks)

	var file, err = os.Open(filePath)
	if err != nil {
		return err
	}

	var r = bufio.NewReader(file)
	var buf = make([]byte, 0, c.BlockSize)
	var index int
	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			var closeErr = file.Close()
			if closeErr != nil {
				return fmt.Errorf("close file err: %w, buf.Read err: %s", closeErr, err.Error()) // cant have two %w
			}
			return err
		}
		if err != nil && err != io.EOF {
			return err
		}

		// Write must not modify the slice data, even temporarily. Implementations must not retain p
		// https://pkg.go.dev/io#Writer
		var transportArr = make([]byte, len(buf))
		copy(transportArr, buf)
		blocks <- block{Index: index, Data: transportArr}
		index++
	}

	return file.Close()
}

// hashBlock runs the hash func on each block of bytes
func (c *ConcurrentHash) hashBlock(blocks <-chan block, sums chan<- sum) error {
	for b := range blocks {
		var h64 hash.Hash64 = murmur3.New64()
		var _, err = h64.Write(b.Data)
		if err != nil {
			return err
		}
		sums <- sum{Index: b.Index, Hash: h64.Sum64()}
		h64.Reset()
	}
	return nil
}
