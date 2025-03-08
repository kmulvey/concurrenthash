package concurrenthash

import (
	"context"
	"fmt"
)

// hashBlock runs the hash func on each block of bytes.
func (c *ConcurrentHash) hashBlock(ctx context.Context, blocks <-chan block, sums chan<- sum) error {
	defer close(sums)
	var h = c.HashConstructor()
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			select {
			case b, open := <-blocks:
				if !open {
					return nil
				}
				h.Reset()
				var _, err = h.Write(b.Data)
				if err != nil {
					return fmt.Errorf("error writing data: %w", err)
				}
				sums <- sum{Index: b.Index, Hash: h.Sum(nil)}
			default:
			}
		}
	}
}
