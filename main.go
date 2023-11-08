package main

import (
	"fmt"
	"gohash/hash"
)

func main() {
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

	directHash := hash.SHA1([]byte(inputToHash))
	fmt.Printf("%x\n", directHash)

	// hash file
	// path to desire file
	hasher.NewFile("./hash/test_cases/go1.21.1.windows-amd64.msi")
	hashForCompiler := hasher.Get()
	fmt.Println(hashForCompiler)
	fmt.Printf("%x\n", hashForCompiler)
	// [36 125 64 190 9 175 100 155 166 165 202 206 51 199 145 54 114 22 193 49]
	// 247d40be09af649ba6a5cace33c791367216c131
}
