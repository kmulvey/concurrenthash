package concurrenthash

import (
	"bytes"
	"context"

	//nolint:gosec
	"crypto/md5"
	//nolint:gosec
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/gob"
	"fmt"
	"hash"
	"os"
	"sync"

	"github.com/kmulvey/goutils"
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

var HashNamesToHashFuncs = map[string]func() hash.Hash{
	"adler32":         WrapAdler32,
	"crc32IEEE":       WrapCrc32IEEE,
	"crc32Castagnoli": WrapCrc32Castagnoli,
	"crc32Koopman":    WrapCrc32Koopman,
	"crc64ISO":        WrapCrc64ISO,
	"crc64ECMA":       WrapCrc64ECMA,
	"fnv32":           WrapFnv32,
	"fnv32a":          WrapFnv32a,
	"fnv64":           WrapFnv64,
	"fnv64a":          WrapFnv64a,
	"sha256":          sha256.New,
	"md5":             md5.New,
	"sha1":            sha1.New,
	"sha512":          sha512.New,
	"murmur32":        WrapMurmur32,
	"murmur64":        WrapMurmur64,
}

// ConcurrentHash is basically a https://en.wikipedia.org/wiki/Merkle_tree
type ConcurrentHash struct {
	Concurrency     int
	BlockSize       int64
	HashConstructor func() hash.Hash

	// internal
	Hashes     [][]byte
	HashesLock sync.RWMutex
}

// NewConcurrentHash is the constructor and entrypoint
func NewConcurrentHash(concurrency int, blockSize int64, hashFunc func() hash.Hash) ConcurrentHash {
	return ConcurrentHash{
		Concurrency:     concurrency,
		BlockSize:       blockSize,
		HashConstructor: hashFunc,
	}
}

// HashFile is a coordination func that fans out to hash workers,
// collects their output and hashes the final array
func (c *ConcurrentHash) HashFile(ctx context.Context, file string) (string, error) {

	// make sure the file even exists first
	var stat, err = os.Stat(file)
	if err != nil {
		return "", err
	}
	c.HashesLock.Lock()
	c.Hashes = make([][]byte, (stat.Size()+c.BlockSize-1)/c.BlockSize)
	c.HashesLock.Unlock()

	// startup all our routines, readers first
	var sumChans = make([]chan sum, c.Concurrency)
	var blocks = make(chan block)
	var errGroup = new(errgroup.Group)

	for i := 0; i < c.Concurrency; i++ {
		var sums = make(chan sum)
		sumChans[i] = sums
		errGroup.Go(func() error {
			return c.hashBlock(ctx, blocks, sums)
		})
	}
	errGroup.Go(func() error {
		c.collectSums(ctx, goutils.MergeChannels(sumChans...))
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

	c.HashesLock.Lock()
	if err := enc.Encode(c.Hashes); err != nil {
		return "", err
	}
	c.HashesLock.Unlock()

	var h = c.HashConstructor()
	h.Reset()
	_, err = h.Write(buf.Bytes())
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
