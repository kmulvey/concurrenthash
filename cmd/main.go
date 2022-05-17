package main

import (
	"flag"
	"fmt"
	"math"
	"runtime"

	"github.com/kmulvey/concurrenthash"
	"github.com/twmb/murmur3"
)

var MB = int64(math.Pow(1024, 2))

func main() {
	var file string
	var threads int
	var blockSize int64

	flag.StringVar(&file, "file", "", "file to hash (abs path)")
	flag.IntVar(&threads, "threads", 1, "number of threads to use, >1 only useful when rebuilding the cache")
	flag.Int64Var(&blockSize, "block-size", 1, "size of chunk to hash in MB")
	flag.Parse()

	blockSize *= MB

	if threads <= 0 || threads > runtime.GOMAXPROCS(0) {
		threads = 1
	}

	var ch = concurrenthash.NewConcurrentHash(threads, blockSize, murmur3.New64())
	var hash, err = ch.HashFile(file)
	if err != nil {
		fmt.Printf("Encountered an error: %s", err.Error())
		return
	}

	fmt.Printf("%s: %s\n", file, hash)
}
