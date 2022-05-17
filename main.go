package concurrenthash

import (
	"bufio"
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
	"sync"

	"github.com/twmb/murmur3"
)

type block struct {
	Index int
	Data  []byte
}
type sum struct {
	Index int
	Hash  uint64
}

type ConcurrentHash struct {
	Concurrency int
	BlockSize   int64
	HashFunc    hash.Hash64

	// internal
	Context context.Context
	Cancel  context.CancelFunc
	Hashes  []uint64
}

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

func (c *ConcurrentHash) HashFile(file string) (string, error) {
	var stat, err = os.Stat(file)
	if err != nil {
		c.Cancel()
		return "", err
	}
	c.Hashes = make([]uint64, (stat.Size()+c.BlockSize-1)/c.BlockSize)

	var sums = make(chan sum)
	var blocks = make(chan block)

	go c.collectSums(sums)
	var wg sync.WaitGroup
	for i := 0; i < c.Concurrency; i++ {
		wg.Add(1)
		go func() {
			if err := c.hashBlock(blocks, sums, &wg); err != nil {
				log.Fatal(err)
			}
		}()
	}
	go func() {
		if err := c.streamFile(file, blocks); err != nil {
			log.Fatal(err)
		}
	}()
	wg.Wait()

	// hash the hashes
	fmt.Printf("hashes: %+v\n", c.Hashes)
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
			file.Close() // err: too bad
			return err
		}
		if err != nil && err != io.EOF {
			close(blocks)
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

func (c *ConcurrentHash) hashBlock(blocks <-chan block, sums chan<- sum, wg *sync.WaitGroup) error {
	defer wg.Done()
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
