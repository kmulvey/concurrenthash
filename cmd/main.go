package main

import (
	"fmt"

	"github.com/kmulvey/concurrenthash"
	"github.com/twmb/murmur3"
)

func main() {
	var ch = concurrenthash.NewConcurrentHash(2, 2, murmur3.New64())
	fmt.Println(ch.HashFile("1g.img"))
}
