package hash

import "fmt"

func SHA1(input []byte) [20]byte {
	var blocks [][64]byte
	// preparations
	if len(input) < 56 {
		var currentBlock [64]byte
		copy(currentBlock[:], input)
		currentBlock[len(input)] = 1 << 7 // add 1 bit to the end

		fillLastBitsWithInputLength(&currentBlock, len(input))

		blocks = append(blocks, currentBlock)
		// THIS STEP WAS TESTED WITH TEST CASE [http://book.itep.ru/6/sha1.htm]
		// UPD. TESTED WITH EMPTY INPUT
	} else {
		var currentBlock [64]byte
		iterator := 0
		for i := 0; i < len(input); i, iterator = i+1, iterator+1 {
			if iterator == 64 {
				blocks = append(blocks, currentBlock)
				currentBlock = [64]byte{}
				iterator = 0
			}
			currentBlock[iterator] = input[i]
		}
		fmt.Println(iterator)
		// https://ru.wikipedia.org/wiki/SHA-1 [Инициализация]
		if iterator < 55 {
			currentBlock[iterator] = 1 << 7
			fillLastBitsWithInputLength(&currentBlock, len(input))
			blocks = append(blocks, currentBlock)
		} else if iterator == 64 {
			blocks = append(blocks, currentBlock)
			currentBlock = [64]byte{}
			currentBlock[0] = 1 << 7
			fillLastBitsWithInputLength(&currentBlock, len(input))
			blocks = append(blocks, currentBlock)
		} else {
			currentBlock[iterator] = 1 << 7
			blocks = append(blocks, currentBlock)
			currentBlock = [64]byte{}
			fillLastBitsWithInputLength(&currentBlock, len(input))
			blocks = append(blocks, currentBlock)
		}
	}
	// initialize buffers
	var h0, h1, h2, h3, h4 uint32 = 0x67452301, 0xEFCDAB89, 0x98BADCFE, 0x10325476, 0xC3D2E1F0
	// processing blocks
	for _, chunk := range blocks {
		var W [80]uint32
		// fill first 16 words from chunk
		for i, j := 0, 0; i < 64; i, j = i+4, j+1 {
			W[j] = (uint32(chunk[i]) << 24) | (uint32(chunk[i+1]) << 16) | (uint32(chunk[i+2]) << 8) | uint32(chunk[i+3])
			// fmt.Printf("%x\n", W[j])
			// fmt.Println("----------------")
		}
		// THIS STEP WAS TESTED
		// extend from 16 to 80 words
		for i := 16; i < 80; i++ {
			W[i] = cyclicLeftShift((W[i-3] ^ W[i-8] ^ W[i-14] ^ W[i-16]), 1) // CIRLCE LEFT SHIFT! NOT JUST LEFT SHIFT !
		}
		// initialize working vars
		a, b, c, d, e := h0, h1, h2, h3, h4
		// MAIN LOOP
		for i := 0; i < 80; i++ {
			var K, F uint32
			switch {
			case i >= 0 && i <= 19:
				F = (b & c) | ((^b) & d)
				K = 0x5A827999
			case i >= 20 && i <= 39:
				F = b ^ c ^ d
				K = 0x6ED9EBA1
			case i >= 40 && i <= 59:
				F = (b & c) | (b & d) | (c & d)
				K = 0x8F1BBCDC
			case i >= 60 && i <= 79:
				F = b ^ c ^ d
				K = 0xCA62C1D6
			}

			temp := (cyclicLeftShift(a, 5)) + F + e + K + W[i] // CIRLCE LEFT SHIFT! NOT JUST LEFT SHIFT !
			e = d
			d = c
			c = cyclicLeftShift(b, 30) // CIRLCE LEFT SHIFT! NOT JUST LEFT SHIFT !
			b = a
			a = temp
		}
		// ADD CHUNK TO THE RESULT/NEXT-OPER.
		h0 = h0 + a
		h1 = h1 + b
		h2 = h2 + c
		h3 = h3 + d
		h4 = h4 + e
	}

	var result [20]byte
	for i := 0; i < 20; i += 4 {
		var h uint32
		if i == 0 {
			h = h0
		} else if i == 4 {
			h = h1
		} else if i == 8 {
			h = h2
		} else if i == 12 {
			h = h3
		} else if i == 16 {
			h = h4
		}

		result[i] = byte(h >> 24)
		result[i+1] = byte(h >> 16)
		result[i+2] = byte(h >> 8)
		result[i+3] = byte(h)
	}
	return result
}

/*
HELPERS
*/
func getBinaryLength(length int) [8]byte {
	// constant for amount of required shift
	right := [8]int{56, 48, 40, 32, 24, 16, 8, 0}
	var binary64bitlength [8]byte
	for i := 0; i < 8; i++ {
		binary64bitlength[i] = uint8(length >> right[i])
	}
	return binary64bitlength
}

func cyclicLeftShift(value uint32, shift int) uint32 {
	return (value << shift) | (value >> (32 - shift))
}

func fillLastBitsWithInputLength(currentBlock *[64]byte, input int) {
	binaryLength := getBinaryLength(input * 8) // get binary length of the input (64bit value in [8]byte array format)
	// add binary length to the 512 block
	for i, j := 56, 0; i < 64 && j < 8; i, j = i+1, j+1 {
		currentBlock[i] = binaryLength[j]
	}
}
