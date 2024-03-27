package main

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"math"
	"os"
	"testing"
	"time"

	"github.com/kmulvey/concurrenthash"
	"github.com/stretchr/testify/assert"
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

// do not run this with -race
func TestBench(t *testing.T) {

	var filename = createRandFile(t)
	defer removeFile(t, filename)

	var ctx = context.Background()
	for name, f := range argToHashFuncMap {
		for blockSize := int64(10000); blockSize <= 1e8; blockSize *= 10 {
			var start = time.Now()
			var ch = concurrenthash.NewConcurrentHash(4, blockSize, f)
			var _, err = ch.HashFile(ctx, filename)
			if err != nil {
				fmt.Printf("Encountered an error: %s \n", err.Error())
				return
			}
			fmt.Printf("name: %s, block size: %d, milliseconds: %d \n", name, blockSize, time.Since(start).Milliseconds())
		}
	}
}

func createRandFile(t *testing.T) string {

	var filename = "./rand.txt"
	removeFile(t, filename)

	file, err := os.Create(filename)
	assert.NoError(t, err)
	defer file.Close()

	token := make([]byte, 100)
	var bytesWritten int

	for bytesWritten <= int(math.Pow(1024, 2))*50 {
		rand.Read(token)
		n, err := file.Write(token)
		assert.NoError(t, err)
		bytesWritten += n
	}

	return filename
}

func removeFile(t *testing.T, file string) {
	if _, err := os.Stat(file); err == nil {
		assert.NoError(t, os.RemoveAll(file))
	}
}
