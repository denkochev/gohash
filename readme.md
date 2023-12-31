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

Results for 1,13 GB (1 222 355 456 bytes) file:
```
goos: windows
goarch: amd64
pkg: gohash/hash
cpu: Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz
BenchmarkMySHA1-8              1        8904308200 ns/op        8391299920 B/op      168 allocs/op
BenchmarkLibSHA1-8             1        1460139800 ns/op        1222390032 B/op       20 allocs/op
PASS
ok      gohash/hash     11.268s

```

CPU:
```
      flat  flat%   sum%        cum   cum%
     3.55s 44.88% 44.88%      5.22s 65.99%  gohash/hash.SHA1
     1.17s 14.79% 59.67%      1.17s 14.79%  runtime.memmove
     0.71s  8.98% 68.65%      0.71s  8.98%  crypto/sha1.blockAVX2
     0.53s  6.70% 75.35%      0.53s  6.70%  runtime.procyield
     0.51s  6.45% 81.80%      0.51s  6.45%  runtime.cgocall
     0.49s  6.19% 87.99%      0.49s  6.19%  runtime.stdcall3
     0.31s  3.92% 91.91%      0.31s  3.92%  runtime.memclrNoHeapPointers
     0.20s  2.53% 94.44%      0.20s  2.53%  gohash/hash.cyclicLeftShift (inline)
     0.12s  1.52% 95.95%      0.12s  1.52%  runtime.stdcall1
     0.05s  0.63% 96.59%      0.05s  0.63%  runtime.writeHeapBits.flush

```

Memory:

![memory](./docs/imgs/memory.png)

```
Showing nodes accounting for 9325.77MB, 99.43% of 9379.58MB total
Dropped 21 nodes (cum <= 46.90MB)
      flat  flat%   sum%        cum   cum%
 6919.06MB 73.77% 73.77%  6919.06MB 73.77%  gohash/hash.SHA1
 2337.66MB 24.92% 98.69%  2337.66MB 24.92%  os.ReadFile
   67.04MB  0.71% 99.40%   200.80MB  2.14%  gohash/hash.Test_CompareWithCryptoSHA1Lib
       2MB 0.021% 99.43%  6924.06MB 73.82%  gohash/hash.(*Hash).New
         0     0% 99.43%  1165.73MB 12.43%  gohash/hash.BenchmarkLibSHA1
         0     0% 99.43%  8004.25MB 85.34%  gohash/hash.BenchmarkMySHA1
         0     0% 99.43%  9169.98MB 97.77%  testing.(*B).run1.func1
         0     0% 99.43%  9169.98MB 97.77%  testing.(*B).runN
         0     0% 99.43%   206.99MB  2.21%  testing.tRunner
```