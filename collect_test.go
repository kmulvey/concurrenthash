package concurrenthash

import (
	"context"
	"testing"

	"crypto/sha256"
)

func TestCollectSums(t *testing.T) {
	t.Parallel()

	var ctx, cancel = context.WithCancel(context.Background())
	var cs = NewConcurrentHash(2, 10, sha256.New)
	cs.Hashes = make([][]byte, 2)
	var sums = make(chan sum)
	go cs.collectSums(ctx, sums)

	sums <- sum{
		Index: 1,
		Hash:  []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f},
	}
	cancel()
}
