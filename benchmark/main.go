package main

import (
	"context"
	//nolint:gosec
	"crypto/md5"
	//nolint:gosec
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"time"

	"github.com/kmulvey/concurrenthash"
)

var argToHashFuncMap = map[string]func() hash.Hash{
	"adler32":         concurrenthash.WrapAdler32,
	"crc32IEEE":       concurrenthash.WrapCrc32IEEE,
	"crc32Castagnoli": concurrenthash.WrapCrc32Castagnoli,
	"crc32Koopman":    concurrenthash.WrapCrc32Koopman,
	"crc64ISO":        concurrenthash.WrapCrc64ISO,
	"crc64ECMA":       concurrenthash.WrapCrc64ECMA,
	"fnv32":           concurrenthash.WrapFnv32,
	"fnv32a":          concurrenthash.WrapFnv32a,
	"fnv64":           concurrenthash.WrapFnv64,
	"fnv64a":          concurrenthash.WrapFnv64a,
	"sha256":          sha256.New,
	"md5":             md5.New,
	"sha1":            sha1.New,
	"sha512":          sha512.New,
	"murmur32":        concurrenthash.WrapMurmur32,
	"murmur64":        concurrenthash.WrapMurmur64,
}

func main() {
	var ctx = context.Background()
	for name, f := range argToHashFuncMap {
		for blockSize := int64(10000); blockSize <= 1e8; blockSize *= 10 {
			var start = time.Now()
			var ch = concurrenthash.NewConcurrentHash(4, blockSize, f)
			var _, err = ch.HashFile(ctx, "../rand-file.txt")
			if err != nil {
				fmt.Printf("Encountered an error: %s \n", err.Error())
				return
			}
			fmt.Printf("%s, %d, %v\n", name, blockSize, time.Since(start).Milliseconds())
		}
	}
}
