package main

import (
	"context"

	"flag"
	"math"
	"os"
	"runtime"
	"sort"
	"strings"

	"github.com/kmulvey/concurrenthash"
	log "github.com/sirupsen/logrus"
)

// nolint:gochecknoglobals
// MB is a megabyte.
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
		var names = make([]string, len(concurrenthash.HashNamesToHashFuncs))
		for name := range concurrenthash.HashNamesToHashFuncs {
			names[i] = name
			i++
		}
		sort.Strings(names)
		log.Info("Supported hashing algorithms: " + strings.Join(names, ", "))
		os.Exit(0)
	}

	blockSize *= MB

	if threads <= 0 || threads > runtime.GOMAXPROCS(0) {
		threads = 1
	}

	if _, exists := concurrenthash.HashNamesToHashFuncs[hashFunc]; !exists {
		log.Error("Hash function", hashFunc, "is not supported")
		os.Exit(1)
	}

	var ch = concurrenthash.NewConcurrentHash(threads, blockSize, concurrenthash.HashNamesToHashFuncs[hashFunc])
	var hash, err = ch.HashFile(ctx, file)
	if err != nil {
		log.Errorf("Encountered an error: %s", err.Error())
		os.Exit(1)
	}

	log.Infof("%s: %s\n", file, hash)
}
