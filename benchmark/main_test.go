package main

import (
	"crypto/rand"

	"math"
	"os"
	"testing"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kmulvey/concurrenthash"
	"github.com/stretchr/testify/assert"
)

func BenchmarkHashes(b *testing.B) {

	var filename = createRandFile(b)
	defer removeFile(b, filename)

	var ctx = b.Context()

	var grid = table.NewWriter()
	grid.SetOutputMirror(os.Stdout)
	grid.AppendHeader(table.Row{"Name", "Block Size", "Milliseconds"})

	for name, f := range concurrenthash.HashNamesToHashFuncs {
		for blockSize := int64(10000); blockSize <= 1e8; blockSize *= 10 {
			var start = time.Now()
			var ch = concurrenthash.NewConcurrentHash(4, blockSize, f)
			var _, err = ch.HashFile(ctx, filename)
			if err != nil {
				assert.NoError(b, err)
			}
			grid.AppendRow([]any{name, blockSize, time.Since(start).Milliseconds()})
		}
	}
	grid.Render()
}

func createRandFile(b *testing.B) string {
	b.Helper()

	var filename = "./rand.txt"
	removeFile(b, filename)

	file, err := os.Create(filename)
	assert.NoError(b, err)
	defer file.Close()

	token := make([]byte, 100)
	var bytesWritten int

	for bytesWritten <= int(math.Pow(1024, 2))*250 {
		var _, err = rand.Read(token)
		assert.NoError(b, err)
		n, err := file.Write(token)
		assert.NoError(b, err)
		bytesWritten += n
	}

	return filename
}

func removeFile(b *testing.B, file string) {
	b.Helper()
	if _, err := os.Stat(file); err == nil {
		assert.NoError(b, os.RemoveAll(file))
	}
}
