package main

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"math"
	"os"
	"runtime"
	"sort"
	"strings"

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
var MB = int64(math.Pow(1024, 2))

func main() {
	var file string
	var hashFunc string
	var threads int
	var blockSize int64
	var algos bool
	var ctx = context.Background()

	flag.BoolVar(&algos, "algos", false, "file to hash (abs path)")
	flag.StringVar(&file, "file", "", "file to hash (abs path)")
	flag.StringVar(&hashFunc, "hash-func", "sha256", "hash algorithm code, run: `concurrenthash -algos` for list ")
	flag.IntVar(&threads, "threads", 1, "number of threads to use, >1 only useful when rebuilding the cache")
	flag.Int64Var(&blockSize, "block-size", 1, "size of chunk to hash in MB")
	flag.Parse()

	if algos {
		var i int
		var names = make([]string, len(argToHashFuncMap))
		for name := range argToHashFuncMap {
			names[i] = name
			i++
		}
		sort.Strings(names)
		fmt.Println("Supported hashing algorithms: " + strings.Join(names, ", "))
		os.Exit(0)
	}

	blockSize *= MB

	if threads <= 0 || threads > runtime.GOMAXPROCS(0) {
		threads = 1
	}

	if _, exists := argToHashFuncMap[hashFunc]; !exists {
		fmt.Println("Hash function", hashFunc, "is not supported")
		os.Exit(1)
	}

	var ch = concurrenthash.NewConcurrentHash(threads, blockSize, argToHashFuncMap[hashFunc])
	var hash, err = ch.HashFile(ctx, file)
	if err != nil {
		fmt.Printf("Encountered an error: %s", err.Error())
		return
	}

	fmt.Printf("%s: %s\n", file, hash)
}
