package concurrenthash

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

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
