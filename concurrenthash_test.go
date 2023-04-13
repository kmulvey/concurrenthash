package concurrenthash

import (
	"context"
	"testing"

	"crypto/sha256"

	"github.com/stretchr/testify/assert"
)

func TestEverything(t *testing.T) {
	t.Parallel()

	var ctx = context.Background()
	var cs = NewConcurrentHash(2, 10, sha256.New)
	var sum, err = cs.HashFile(ctx, "./rand-file.txt")
	assert.NoError(t, err)
	assert.Equal(t, "bf842e96b246556052bc7e518de1fdf7c4a5a859ad104a201880074bece30b82", sum)

	sum, err = cs.HashFile(ctx, "./sdfsdfsf.txt")
	assert.Equal(t, "stat ./sdfsdfsf.txt: no such file or directory", err.Error())
	assert.Equal(t, "", sum)
}
