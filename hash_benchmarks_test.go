package concurrenthash

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

// nolint: gochecknoglobals
var (
	text        = []byte("this is a password")
	benchMatrix = map[string]testPair{
		"Adler32":         {HashFunc: WrapAdler32, Expected: "23a01065", BenchExpected: "3ebb06c9"},
		"Crc32IEEE":       {HashFunc: WrapCrc32IEEE, Expected: "adb01c6e", BenchExpected: "be4ba7f3"},
		"Crc32Castagnoli": {HashFunc: WrapCrc32Castagnoli, Expected: "99b59cc7", BenchExpected: "7c017eb2"},
		"Crc32Koopman":    {HashFunc: WrapCrc32Koopman, Expected: "e9babd88", BenchExpected: "9f72811d"},
		"Crc64ISO":        {HashFunc: WrapCrc64ISO, Expected: "a9d108f9820cbe08", BenchExpected: "1a6bdcf4a1c9ba03"},
		"Crc64ECMA":       {HashFunc: WrapCrc64ECMA, Expected: "7b0ffccb93d003c8", BenchExpected: "28487e2cbef9ad8b"},
		"Fnv32":           {HashFunc: WrapFnv32, Expected: "fc61f007", BenchExpected: "4123a0f9"},
		"Fnv32a":          {HashFunc: WrapFnv32a, Expected: "84eaae15", BenchExpected: "44dedf25"},
		"Fnv64":           {HashFunc: WrapFnv64, Expected: "17c6956e1d89b318", BenchExpected: "4d237a2a6f6d06b9"},
		"Fnv64a":          {HashFunc: WrapFnv64a, Expected: "0d2bc6f6ebf6d810", BenchExpected: "cc1156362b0151a5"},
		"Murmur32":        {HashFunc: WrapMurmur32, Expected: "824cebde", BenchExpected: "4ff4843f"},
		"Murmur64":        {HashFunc: WrapMurmur64, Expected: "b75d2949c06573bd", BenchExpected: "ea2e3dc47ab75da4"},
		"Sha224":          {HashFunc: WrapSha224, Expected: "68808b99d5d240e3dc11db9904cd1891a204b7795639cb2e15eea029", BenchExpected: "ba69b65fb22a4187f46f452853503687bfb025b71de3bb463ceb66bc"},
		"Sha256":          {HashFunc: WrapSha256, Expected: "4870bc3d9a751543ee66685fde1a81a78265bed8e532878fb45da0cb08aa5f3c", BenchExpected: "289ca48885442b5480dd76df484e1f90867a2961493b7c60e542e84addce5d1e"},
		"Sha384":          {HashFunc: WrapSha384, Expected: "f38e9bc1b649513b3c4eb6bb4b11c86cd23a55b42d78087eddc0f649c810c542c16beb1c35de339d001884cb79b8c4c4", BenchExpected: "e830d319ba7ac2a65dcc2db9edeae29610861b14fb2e1549019db1684a2eb892e97513646f82e8f58b6d38e0971e15e8"},
		"Sha512":          {HashFunc: WrapSha512, Expected: "abb5de305b09ed982ead4fd13855ea1b6e50f462e01002e9d174309e82ead36c159f743b8e7208c10aca8c3ac116b2398afab4611b2f9efc0652a84e126d515a", BenchExpected: "e2c89ac20dfa255b9e7aaa99f9f356d9fdf4f3f0f1e10d46c26e1799fd8433608349a52bc3b341393bdff49a2862dbf9282e3204586a09d2bc3795eb7fb835cb"},
		"Sha512224":       {HashFunc: WrapSha512224, Expected: "1b851853155f85ea87e0b96c2ba2e6f47f166acbd764641a790f6bac", BenchExpected: "3d9655754e4fae0d6f1ba9db5836917d44ed7a185301ceed58a9d72c"},
		"Sha512256":       {HashFunc: WrapSha512256, Expected: "c7e8e3d5eec72a5375ba6a51f03ee6027237cc6edbb9213957c6bd5fc72f62f5", BenchExpected: "aea31c2217a2decdbf057dd219d793dc7068570f8ecc9211737505af51dc028b"},
	}
)

func BenchmarkSha512256(b *testing.B) {
	for b.Loop() {
		var hasher = benchMatrix["Sha512256"].HashFunc()
		hasher.Write(text)
		assert.Equal(b, benchMatrix["Sha512256"].BenchExpected, hex.EncodeToString(hasher.Sum(nil)))
	}
}
func BenchmarkSha512224(b *testing.B) {
	for b.Loop() {
		var hasher = benchMatrix["Sha512224"].HashFunc()
		hasher.Write(text)
		assert.Equal(b, benchMatrix["Sha512224"].BenchExpected, hex.EncodeToString(hasher.Sum(nil)))
	}
}
func BenchmarkSha512(b *testing.B) {
	for b.Loop() {
		var hasher = benchMatrix["Sha512"].HashFunc()
		hasher.Write(text)
		assert.Equal(b, benchMatrix["Sha512"].BenchExpected, hex.EncodeToString(hasher.Sum(nil)))
	}
}
func BenchmarkSha384(b *testing.B) {
	for b.Loop() {
		var hasher = benchMatrix["Sha384"].HashFunc()
		hasher.Write(text)
		assert.Equal(b, benchMatrix["Sha384"].BenchExpected, hex.EncodeToString(hasher.Sum(nil)))
	}
}
func BenchmarkSha256(b *testing.B) {
	for b.Loop() {
		var hasher = benchMatrix["Sha256"].HashFunc()
		hasher.Write(text)
		assert.Equal(b, benchMatrix["Sha256"].BenchExpected, hex.EncodeToString(hasher.Sum(nil)))
	}
}
func BenchmarkSha224(b *testing.B) {
	for b.Loop() {
		var hasher = benchMatrix["Sha224"].HashFunc()
		hasher.Write(text)
		assert.Equal(b, benchMatrix["Sha224"].BenchExpected, hex.EncodeToString(hasher.Sum(nil)))
	}
}
func BenchmarkMurmur64(b *testing.B) {
	for b.Loop() {
		var hasher = benchMatrix["Murmur64"].HashFunc()
		hasher.Write(text)
		assert.Equal(b, benchMatrix["Murmur64"].BenchExpected, hex.EncodeToString(hasher.Sum(nil)))
	}
}
func BenchmarkMurmur32(b *testing.B) {
	for b.Loop() {
		var hasher = benchMatrix["Murmur32"].HashFunc()
		hasher.Write(text)
		assert.Equal(b, benchMatrix["Murmur32"].BenchExpected, hex.EncodeToString(hasher.Sum(nil)))
	}
}
func BenchmarkFnv64a(b *testing.B) {
	for b.Loop() {
		var hasher = benchMatrix["Fnv64a"].HashFunc()
		hasher.Write(text)
		assert.Equal(b, benchMatrix["Fnv64a"].BenchExpected, hex.EncodeToString(hasher.Sum(nil)))
	}
}
func BenchmarkFnv64(b *testing.B) {
	for b.Loop() {
		var hasher = benchMatrix["Fnv64"].HashFunc()
		hasher.Write(text)
		assert.Equal(b, benchMatrix["Fnv64"].BenchExpected, hex.EncodeToString(hasher.Sum(nil)))
	}
}
func BenchmarkFnv32a(b *testing.B) {
	for b.Loop() {
		var hasher = benchMatrix["Fnv32a"].HashFunc()
		hasher.Write(text)
		assert.Equal(b, benchMatrix["Fnv32a"].BenchExpected, hex.EncodeToString(hasher.Sum(nil)))
	}
}
func BenchmarkFnv32(b *testing.B) {
	for b.Loop() {
		var hasher = benchMatrix["Fnv32"].HashFunc()
		hasher.Write(text)
		assert.Equal(b, benchMatrix["Fnv32"].BenchExpected, hex.EncodeToString(hasher.Sum(nil)))
	}
}
func BenchmarkCrc64ECMA(b *testing.B) {
	for b.Loop() {
		var hasher = benchMatrix["Crc64ECMA"].HashFunc()
		hasher.Write(text)
		assert.Equal(b, benchMatrix["Crc64ECMA"].BenchExpected, hex.EncodeToString(hasher.Sum(nil)))
	}
}
func BenchmarkCrc64ISO(b *testing.B) {
	for b.Loop() {
		var hasher = benchMatrix["Crc64ISO"].HashFunc()
		hasher.Write(text)
		assert.Equal(b, benchMatrix["Crc64ISO"].BenchExpected, hex.EncodeToString(hasher.Sum(nil)))
	}
}
func BenchmarkCrc32Koopman(b *testing.B) {
	for b.Loop() {
		var hasher = benchMatrix["Crc32Koopman"].HashFunc()
		hasher.Write(text)
		assert.Equal(b, benchMatrix["Crc32Koopman"].BenchExpected, hex.EncodeToString(hasher.Sum(nil)))
	}
}
func BenchmarkCrc32Castagnoli(b *testing.B) {
	for b.Loop() {
		var hasher = benchMatrix["Crc32Castagnoli"].HashFunc()
		hasher.Write(text)
		assert.Equal(b, benchMatrix["Crc32Castagnoli"].BenchExpected, hex.EncodeToString(hasher.Sum(nil)))
	}
}
func BenchmarkCrc32IEEE(b *testing.B) {
	for b.Loop() {
		var hasher = benchMatrix["Crc32IEEE"].HashFunc()
		hasher.Write(text)
		assert.Equal(b, benchMatrix["Crc32IEEE"].BenchExpected, hex.EncodeToString(hasher.Sum(nil)))
	}
}
func BenchmarkAdler32(b *testing.B) {
	for b.Loop() {
		var hasher = benchMatrix["Adler32"].HashFunc()
		hasher.Write(text)
		assert.Equal(b, benchMatrix["Adler32"].BenchExpected, hex.EncodeToString(hasher.Sum(nil)))
	}
}
