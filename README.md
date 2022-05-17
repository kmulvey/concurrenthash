# concurrenthash
[![concurrenthash](https://github.com/kmulvey/concurrenthash/actions/workflows/release_build.yml/badge.svg)](https://github.com/kmulvey/concurrenthash/actions/workflows/release_build.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/kmulvey/concurrenthash)](https://goreportcard.com/report/github.com/kmulvey/concurrenthash) [![Go Reference](https://pkg.go.dev/badge/github.com/kmulvey/concurrenthash.svg)](https://pkg.go.dev/github.com/kmulvey/concurrenthash)


[Motivation](https://stackoverflow.com/questions/7015544/calculating-a-hash-code-for-a-large-file-in-parallel)

## Benchmarks
Time to hash a 10GB file of /dev/urandom data
| algo              | time (s) |
|-------------------|----------
| adler32           | 8.02  
| crc32Castagnoli   | 6.60
| crc32IEEE         | 4.99
| crc32Koopman      | 6.18
| crc64ECMA         | 5.54
| crc64ISO          | 4.20
| fnv32             | 4.15
| fnv32a            | 4.14
| fnv64             | 4.58
| fnv64a            | 4.33
| md5               | 19.01
| murmur32          | 5.05
| murmur64          | 5.36
| sha1              | 13.07
| sha256            | 18.14
| sha512            | 18.10
