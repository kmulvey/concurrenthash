package concurrenthash

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrappers(t *testing.T) {
	t.Parallel()

	var cs = NewConcurrentHash(context.Background(), 2, 10, WrapAdler32)
	var sum, err = cs.HashFile("./rand-file.txt")
	assert.NoError(t, err)
	assert.Equal(t, "6815116b", sum)

	cs = NewConcurrentHash(context.Background(), 2, 10, WrapCrc32IEEE)
	sum, err = cs.HashFile("./rand-file.txt")
	assert.NoError(t, err)
	assert.Equal(t, "6f81847d", sum)

	cs = NewConcurrentHash(context.Background(), 2, 10, WrapCrc32Castagnoli)
	sum, err = cs.HashFile("./rand-file.txt")
	assert.NoError(t, err)
	assert.Equal(t, "1e094105", sum)

	cs = NewConcurrentHash(context.Background(), 2, 10, WrapCrc32Koopman)
	sum, err = cs.HashFile("./rand-file.txt")
	assert.NoError(t, err)
	assert.Equal(t, "028f318a", sum)

	cs = NewConcurrentHash(context.Background(), 2, 10, WrapCrc64ISO)
	sum, err = cs.HashFile("./rand-file.txt")
	assert.NoError(t, err)
	assert.Equal(t, "e7d62d1e76bd73da", sum)

	cs = NewConcurrentHash(context.Background(), 2, 10, WrapCrc64ECMA)
	sum, err = cs.HashFile("./rand-file.txt")
	assert.NoError(t, err)
	assert.Equal(t, "621d482973c9cc9c", sum)

	cs = NewConcurrentHash(context.Background(), 2, 10, WrapFnv32)
	sum, err = cs.HashFile("./rand-file.txt")
	assert.NoError(t, err)
	assert.Equal(t, "8229dab3", sum)

	cs = NewConcurrentHash(context.Background(), 2, 10, WrapFnv32a)
	sum, err = cs.HashFile("./rand-file.txt")
	assert.NoError(t, err)
	assert.Equal(t, "6bfaa839", sum)

	cs = NewConcurrentHash(context.Background(), 2, 10, WrapFnv64)
	sum, err = cs.HashFile("./rand-file.txt")
	assert.NoError(t, err)
	assert.Equal(t, "faf332cf3012c68c", sum)

	cs = NewConcurrentHash(context.Background(), 2, 10, WrapFnv64a)
	sum, err = cs.HashFile("./rand-file.txt")
	assert.NoError(t, err)
	assert.Equal(t, "3bdd3a9af0591c54", sum)

	cs = NewConcurrentHash(context.Background(), 2, 10, WrapMurmur32)
	sum, err = cs.HashFile("./rand-file.txt")
	assert.NoError(t, err)
	assert.Equal(t, "46846d69", sum)

	cs = NewConcurrentHash(context.Background(), 2, 10, WrapMurmur64)
	sum, err = cs.HashFile("./rand-file.txt")
	assert.NoError(t, err)
	assert.Equal(t, "3cdc2235f66f74bf", sum)
}
