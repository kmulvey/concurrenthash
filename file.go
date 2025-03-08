package concurrenthash

import (
	"bufio"
	"errors"
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
		return fmt.Errorf("error opening file: %s, err: %w", filePath, err)
	}

	var r = bufio.NewReader(file)
	var index int
	for {
		var data = make([]byte, c.BlockSize)
		n, err := io.ReadFull(r, data)
		data = data[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			var closeErr = file.Close()
			if closeErr != nil {
				return fmt.Errorf("error closing file: %s, err: %w", filePath, err)
			}
			return fmt.Errorf("error reading file: %s, err: %w", filePath, err)
		}

		blocks <- block{Index: index, Data: data}
		index++

		if err != nil {
			if err == io.EOF {
				break
			}
			// assuming it is not an error
			// if the last file block is short
			if errors.Is(err, io.ErrUnexpectedEOF) {
				break
			}
			return fmt.Errorf("error reading file: %s, err: %w", filePath, err)
		}
	}

	if err := file.Close(); err != nil {
		return fmt.Errorf("error closing file: %s, err: %w", filePath, err)
	}

	return nil
}
