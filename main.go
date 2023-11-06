package main

import (
	"crypto/sha1"
	"gohash/hash"
)

func main() {
	h := sha1.New()
	h.Write([]byte("are you?"))

	//hashH := fmt.Sprintf("%x", h.Sum(nil))

	//fmt.Println([]byte("are you?"))
	//fmt.Println([]byte("are you?"))

	//hash.SHA1([]byte("are you?"))

	// from http://book.itep.ru/6/sha1.htm
	hash.SHA1([]byte{97, 98, 99, 100, 101})
}
