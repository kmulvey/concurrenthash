package concurrenthash

import (
	"fmt"
	"hash"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twmb/murmur3"
)

func TestHash(t *testing.T) {
	var b = []byte{0x0, 0x0, 0x0, 0x20, 0x66, 0x74, 0x79, 0x70, 0x69, 0x73}

	var h64 hash.Hash64 = murmur3.New64()
	var _, err = h64.Write(b)
	assert.NoError(t, err)
	fmt.Println(h64.Sum64())
	h64.Reset()
}
