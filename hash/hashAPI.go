package hash

import (
	"fmt"
	"os"
)

/*
API FOR LIBRARY
*/
type Hash struct {
	hash [20]byte
	hex  string
}

func (h *Hash) New(input []byte) {
	hash := SHA1(input)
	h.hash = hash
	h.hex = fmt.Sprintf("%x", hash)
}

func (h *Hash) NewFile(path string) {
	// ./hash/test_cases/go1.21.1.windows-amd64.msi
	file, _ := os.ReadFile(path)
	hash := SHA1(file)
	h.hash = hash
	h.hex = fmt.Sprintf("%x", hash)
}

func (h *Hash) Get() [20]byte {
	return h.hash
}

func (h *Hash) GetHex() string {
	return h.hex
}
