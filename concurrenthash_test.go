package concurrenthash

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twmb/murmur3"
)

func TestEverything(t *testing.T) {
	var cs = NewConcurrentHash(2, 10, murmur3.New64())
	var sum, err = cs.HashFile("./rand-file.txt")
	assert.NoError(t, err)
	assert.Equal(t, "8434139902547509435", sum)
}
