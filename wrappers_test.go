package concurrenthash

import (
	"context"
	"hash"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testPair struct {
	HashFunc func() hash.Hash
	Expected string
}

var testMatrix = []testPair{
	{HashFunc: WrapAdler32, Expected: "6815116b"},
	{HashFunc: WrapCrc32IEEE, Expected: "6f81847d"},
	{HashFunc: WrapCrc32Castagnoli, Expected: "1e094105"},
	{HashFunc: WrapCrc32Koopman, Expected: "028f318a"},
	{HashFunc: WrapCrc64ISO, Expected: "e7d62d1e76bd73da"},
	{HashFunc: WrapCrc64ECMA, Expected: "621d482973c9cc9c"},
	{HashFunc: WrapFnv32, Expected: "8229dab3"},
	{HashFunc: WrapFnv32a, Expected: "6bfaa839"},
	{HashFunc: WrapFnv64, Expected: "faf332cf3012c68c"},
	{HashFunc: WrapFnv64a, Expected: "3bdd3a9af0591c54"},
	{HashFunc: WrapMurmur32, Expected: "46846d69"},
	{HashFunc: WrapMurmur64, Expected: "3cdc2235f66f74bf"},
	{HashFunc: WrapSha512, Expected: "570679214223a4a0f7cd82cdf6f433fd897b7e8e1776a9c3e6842f8d2f8f7211b1aaabf06703e78d95e44d1068b88a15fb6bffc70e870580b070548dd63e9d64"},
	{HashFunc: WrapSha384, Expected: "fe7ce20c68d58076774402eb2bb27c57978a69f3ae7778387ddd52bc9e4d48f7a42496a97dd97bb560be23ecccecabd5"},
	{HashFunc: WrapSha512224, Expected: "8fa68c2902b434728a005ffdc77338b84b7784f58e7f6832b09d1d2b"},
	{HashFunc: WrapSha512256, Expected: "800b50f0c527cd902b7d5478ea9bbc277b2aaa6e15d5e9ff024cbf31ed582661"},
}

func TestWrappers(t *testing.T) {
	t.Parallel()

	for _, pair := range testMatrix {
		var cs = NewConcurrentHash(context.Background(), 2, 10, pair.HashFunc)
		var sum, err = cs.HashFile("./rand-file.txt")
		assert.NoError(t, err)
		assert.Equal(t, pair.Expected, sum)
	}
}
