package concurrenthash

import (
	"hash"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"

	"github.com/twmb/murmur3"
)

func WrapAdler32() hash.Hash {
	var h = adler32.New()
	return h
}

func WrapCrc32IEEE() hash.Hash {
	var h = crc32.NewIEEE()
	return h
}

func WrapCrc32Castagnoli() hash.Hash {
	var h = crc32.New(crc32.MakeTable(crc32.Castagnoli))
	return h
}

func WrapCrc32Koopman() hash.Hash {
	var h = crc32.New(crc32.MakeTable(crc32.Koopman))
	return h
}

/*
func WrapCrc32Custom(table uint32) func() hash.Hash {
	return func() hash.Hash {
		var h = crc32.New(crc32.MakeTable(table))
		return h
	}
}
*/

func WrapCrc64ISO() hash.Hash {
	var h = crc64.New(crc64.MakeTable(crc64.ISO))
	return h
}

func WrapCrc64ECMA() hash.Hash {
	var h = crc64.New(crc64.MakeTable(crc64.ECMA))
	return h
}

func WrapFnv32() hash.Hash {
	var h = fnv.New32()
	return h
}

func WrapFnv32a() hash.Hash {
	var h = fnv.New32a()
	return h
}

func WrapFnv64() hash.Hash {
	var h = fnv.New64()
	return h
}

func WrapFnv64a() hash.Hash {
	var h = fnv.New64a()
	return h
}

func WrapMurmur32() hash.Hash {
	var h = murmur3.New32()
	return h
}

func WrapMurmur64() hash.Hash {
	var h = murmur3.New64()
	return h
}
