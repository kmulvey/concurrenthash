package concurrenthash

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
