package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/kmulvey/concurrenthash"
	"github.com/twmb/murmur3"
)

func main() {
	var file string
	var threads int
	var blockSize int

	flag.StringVar(&file, "file", "", "file to hash (abs path)")
	flag.IntVar(&threads, "threads", 1, "number of threads to use, >1 only useful when rebuilding the cache")
	flag.IntVar(&blockSize, "block-size", 1e6, "size of chunk to hash")
	flag.Parse()

	if threads <= 0 || threads > runtime.GOMAXPROCS(0) {
		threads = 1
	}

	var ch = concurrenthash.NewConcurrentHash(2, 2, murmur3.New64())
	var hash, err = ch.HashFile("1g.img")
	if err != nil {
		fmt.Printf("Encountered an error: %s", err.Error())
		return
	}

	fmt.Printf("%s: %s\n", file, hash)
}
