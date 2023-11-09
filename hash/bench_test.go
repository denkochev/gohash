package hash

import (
	"crypto/sha1"
	"fmt"
	"os"
	"testing"
)

var result string

func BenchmarkMySHA1(b *testing.B) {
	hugefile, _ := os.ReadFile("./hash/test_cases/big.txt")
	var s string
	for i := 0; i < b.N; i++ {
		h := Hash{}
		// BUILD IN METHOD FOR FILE READING BUT YOU ALSO CAN USE DEFAULT .NEW() METHOD WITH []byte ARG FORMAT
		h.New(hugefile)
		s = fmt.Sprintf("%x", h.Get())
	}
	result = s
}

func BenchmarkLibSHA1(b *testing.B) {
	hugefile, _ := os.ReadFile("./hash/test_cases/big.txt")
	var s string
	for i := 0; i < b.N; i++ {
		hLib := sha1.New()
		hLib.Write(hugefile)
		s = fmt.Sprintf("%x", hLib.Sum(nil))
	}
	result = s
}
