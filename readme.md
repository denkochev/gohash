# SHA-1 implementation in Go/Golang
Example of implementation SHA-1 hash algorythm in Go language. 

## Usage
You have several options to use SHA-1 hash function by this package. There are class/struct way to use it and raw function.  

Here is example:
```
// Example of struct usage
inputToHash := "You are in Denys git repo!"

hasher := hash.Hash{}
hasher.New([]byte(inputToHash))
// this func returns 160-bit digest in [20]byte format
hashSlice := hasher.Get()
fmt.Println(hashSlice)
// you have 2 options get Hex digest
hexDigest := hasher.GetHex()
hexDigestRaw := fmt.Sprintf("%x", hashSlice)
fmt.Println(hexDigest)
fmt.Println(hexDigestRaw)
// [248 235 36 172 209 248 111 178 133 141 197 76 17 124 178 49 179 150 87 75]
// f8eb24acd1f86fb2858dc54c117cb231b396574b
// f8eb24acd1f86fb2858dc54c117cb231b396574b
```
For lower memory usage you can use SHA-1 func straight.
```
	directHash := hash.SHA1([]byte(inputToHash))
	fmt.Printf("%x\n", directHash)
```
Also you can hash any file on your PC. Package contains build in decoder. Example:
``` 
// let's hash Golang compiler itself
	
// path to desire file
hasher.NewFile("./hash/test_cases/go1.21.1.windows-amd64.msi")
hashForCompiler := hasher.Get()
fmt.Println(hashForCompiler)
fmt.Printf("%x\n", hashForCompiler)
// [36 125 64 190 9 175 100 155 166 165 202 206 51 199 145 54 114 22 193 49]
// 247d40be09af649ba6a5cace33c791367216c131
```

## Testing

This code was tested on rich amount of different data. This package includes test that compare lib hash with original SHA-1 digest. Core tests in file ./hash/compare_test.go

## Performance

There is benchmark test for this implementation and implementation from crypto/sha1 lib. Test is in ./hash/bench_test.go file. You can run these test by call `go test -bench . -benchmem`

Results:
```
goos: windows
goarch: amd64
pkg: gohash/hash
cpu: Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz
BenchmarkMySHA1-8        1291868               904.8 ns/op           256 B/op          7 allocs/op
BenchmarkLibSHA1-8       3653446               296.1 ns/op            96 B/op          3 allocs/op
PASS
ok      gohash/hash     4.045s

```