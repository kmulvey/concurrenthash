package concurrenthash

import (
	"testing"

	"crypto/sha256"

	"github.com/stretchr/testify/assert"
)

func TestEverything(t *testing.T) {
	t.Parallel()

	var ctx = t.Context()
	var cs = NewConcurrentHash(2, 10, sha256.New)
	var sum, err = cs.HashFile(ctx, "./rand-file.txt")
	assert.NoError(t, err)
	assert.Equal(t, "4870bc3d9a751543ee66685fde1a81a78265bed8e532878fb45da0cb08aa5f3c", sum)

	sum, err = cs.HashFile(ctx, "./sdfsdfsf.txt")
	assert.Contains(t, err.Error(), "./sdfsdfsf.txt:") // this is quite loose because windows
	assert.Equal(t, "", sum)
}
