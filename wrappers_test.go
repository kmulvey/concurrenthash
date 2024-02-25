package concurrenthash

import (
	"context"
	"hash"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testPair struct {
	HashFunc      func() hash.Hash
	Expected      string
	BenchExpected string
}

var testMatrix = []testPair{
	{HashFunc: WrapAdler32, Expected: "23a01065"},
	{HashFunc: WrapCrc32IEEE, Expected: "adb01c6e"},
	{HashFunc: WrapCrc32Castagnoli, Expected: "99b59cc7"},
	{HashFunc: WrapCrc32Koopman, Expected: "e9babd88"},
	{HashFunc: WrapCrc64ISO, Expected: "a9d108f9820cbe08"},
	{HashFunc: WrapCrc64ECMA, Expected: "7b0ffccb93d003c8"},
	{HashFunc: WrapFnv32, Expected: "fc61f007"},
	{HashFunc: WrapFnv32a, Expected: "84eaae15"},
	{HashFunc: WrapFnv64, Expected: "17c6956e1d89b318"},
	{HashFunc: WrapFnv64a, Expected: "0d2bc6f6ebf6d810"},
	{HashFunc: WrapMurmur32, Expected: "824cebde"},
	{HashFunc: WrapMurmur64, Expected: "b75d2949c06573bd"},
	{HashFunc: WrapSha512, Expected: "abb5de305b09ed982ead4fd13855ea1b6e50f462e01002e9d174309e82ead36c159f743b8e7208c10aca8c3ac116b2398afab4611b2f9efc0652a84e126d515a"},
	{HashFunc: WrapSha384, Expected: "f38e9bc1b649513b3c4eb6bb4b11c86cd23a55b42d78087eddc0f649c810c542c16beb1c35de339d001884cb79b8c4c4"},
	{HashFunc: WrapSha512224, Expected: "1b851853155f85ea87e0b96c2ba2e6f47f166acbd764641a790f6bac"},
	{HashFunc: WrapSha512256, Expected: "c7e8e3d5eec72a5375ba6a51f03ee6027237cc6edbb9213957c6bd5fc72f62f5"},
	{HashFunc: WrapSha256, Expected: "4870bc3d9a751543ee66685fde1a81a78265bed8e532878fb45da0cb08aa5f3c"},
	{HashFunc: WrapSha224, Expected: "68808b99d5d240e3dc11db9904cd1891a204b7795639cb2e15eea029"},
}

func TestWrappers(t *testing.T) {
	t.Parallel()

	for _, pair := range testMatrix {
		var cs = NewConcurrentHash(2, 10, pair.HashFunc)
		var sum, err = cs.HashFile(context.Background(), "./rand-file.txt")
		assert.NoError(t, err)
		assert.Equal(t, pair.Expected, sum)
	}
}
