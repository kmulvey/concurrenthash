package concurrenthash

import (
	"hash"

	"github.com/twmb/murmur3"
)

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
