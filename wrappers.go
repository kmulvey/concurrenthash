package concurrenthash

import (
	"crypto/sha512"
	"hash"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"

	"github.com/twmb/murmur3"
)

func WrapAdler32() hash.Hash {
	return adler32.New()
}

func WrapCrc32IEEE() hash.Hash {
	return crc32.NewIEEE()
}

func WrapCrc32Castagnoli() hash.Hash {
	return crc32.New(crc32.MakeTable(crc32.Castagnoli))
}

func WrapCrc32Koopman() hash.Hash {
	return crc32.New(crc32.MakeTable(crc32.Koopman))
}

/*
func WrapCrc32Custom(table uint32) func() hash.Hash {
	return func() hash.Hash {
		return  crc32.New(crc32.MakeTable(table))
	}
}
*/

func WrapCrc64ISO() hash.Hash {
	return crc64.New(crc64.MakeTable(crc64.ISO))
}

func WrapCrc64ECMA() hash.Hash {
	return crc64.New(crc64.MakeTable(crc64.ECMA))
}

func WrapFnv32() hash.Hash {
	return fnv.New32()
}

func WrapFnv32a() hash.Hash {
	return fnv.New32a()
}

func WrapFnv64() hash.Hash {
	return fnv.New64()
}

func WrapFnv64a() hash.Hash {
	return fnv.New64a()
}

func WrapMurmur32() hash.Hash {
	return murmur3.New32()
}

func WrapMurmur64() hash.Hash {
	return murmur3.New64()
}

func WrapSha512() hash.Hash {
	return sha512.New()
}

func WrapSha384() hash.Hash {
	return sha512.New384()
}

func WrapSha512224() hash.Hash {
	return sha512.New512_224()
}

func WrapSha512256() hash.Hash {
	return sha512.New512_256()
}
