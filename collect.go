package concurrenthash

import "context"

// collectSums is a fan in func to get the hashes and write them to an array.
func (c *ConcurrentHash) collectSums(ctx context.Context, sums <-chan sum) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			sum, open := <-sums
			if !open {
				return
			}
			c.HashesLock.Lock()
			c.Hashes[sum.Index] = sum.Hash
			c.HashesLock.Unlock()
		}
	}
}
