package concurrenthash

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	t.Parallel()

	var blocks = make(chan block)
	var sums = make(chan sum)
	var ctx, cancel = context.WithCancel(context.Background())
	var cs = ConcurrentHash{Context: ctx}

	go func() {
		for sum := range sums {
			assert.Equal(t, 0, sum.Index)
			assert.Equal(t, uint64(5931046006978006171), sum.Hash)
		}
	}()

	go cs.hashBlock(blocks, sums)

	blocks <- block{Index: 0, Data: []byte{0x32, 0x96, 0xd0, 0x3, 0x2b, 0x56, 0x72, 0x2b, 0xaf, 0x39}}
	time.Sleep(time.Millisecond * 100)
	cancel()
}
