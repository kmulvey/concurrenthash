package concurrenthash

import (
	"context"
	"crypto/sha256"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	t.Parallel()

	var blocks = make(chan block)
	var sums = make(chan sum)
	var ctx, cancel = context.WithCancel(context.Background())
	var cs = ConcurrentHash{Context: ctx, HashConstructor: sha256.New}

	go func() {
		for sum := range sums {
			assert.Equal(t, 0, sum.Index)
			assert.Equal(t, "612dc7d14dac2825e09c75574273db51122fa815657ebc0581a19ada5a5606c2", fmt.Sprintf("%x", sum.Hash))
		}
	}()

	go cs.hashBlock(blocks, sums)

	blocks <- block{Index: 0, Data: []byte{0x32, 0x96, 0xd0, 0x3, 0x2b, 0x56, 0x72, 0x2b, 0xaf, 0x39}}
	time.Sleep(time.Millisecond * 100)
	cancel()
}
