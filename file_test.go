package concurrenthash

import (
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testFileBytes = [][]byte{
	{0x32, 0x96, 0xd0, 0x3, 0x2b, 0x56, 0x72, 0x2b, 0xaf, 0x39},
	{0xdf, 0x69, 0x0, 0xb3, 0x9d, 0x59, 0x5f, 0x55, 0xfe, 0x50},
	{0x62, 0x9a, 0x31, 0x50, 0xba, 0x82, 0xd0, 0xaf, 0x4f, 0xd2},
	{0xf1, 0xae, 0xd, 0x51, 0x51, 0x34, 0xd6, 0xb2, 0xa0, 0xec},
	{0xd7, 0x78, 0x60, 0x27, 0x4e, 0x92, 0x32, 0x69, 0x8b, 0xc},
	{0xff, 0x24, 0x77, 0xfa, 0x13, 0x9e, 0x50, 0xef, 0x54, 0x5f},
	{0x2, 0x68, 0x61, 0xd0, 0xde, 0x5a, 0xb1, 0x7b, 0x6d, 0x61},
	{0x66, 0x91, 0x20, 0xf, 0x47, 0x98, 0x1f, 0x16, 0xba, 0x13},
	{0x68, 0xe7, 0x22, 0x9d, 0xf3, 0x5f, 0xf4, 0x81, 0xe5, 0x26},
	{0xfa, 0x9e, 0x1, 0x12, 0x98, 0x4f, 0x6, 0xb9, 0x8, 0x1f},
}

func TestReadFile(t *testing.T) {
	t.Parallel()

	var blocks = make(chan block)
	var cs = NewConcurrentHash(1, 10, sha256.New)

	var done = make(chan struct{})
	go func() {
		for block := range blocks {
			assert.ElementsMatch(t, testFileBytes[block.Index], block.Data)
		}
		close(done)
	}()

	go func() {
		assert.NoError(t, cs.streamFile("./rand-file.txt", blocks))
	}()
	<-done
}
