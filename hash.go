package concurrenthash

// hashBlock runs the hash func on each block of bytes
func (c *ConcurrentHash) hashBlock(blocks <-chan block, sums chan<- sum) error {
	defer close(sums)
	var h = c.HashConstructor()
	for {
		select {
		case <-c.Context.Done():
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
					return err
				}
				sums <- sum{Index: b.Index, Hash: h.Sum(nil)}
			default:
			}
		}
	}
}
