package hash

const (
	h0 = 0x67452301
	h1 = 0xEFCDAB89
	h2 = 0x98BADCFE
	h3 = 0x10325476
	h4 = 0xC3D2E1F0
)

func SHA1(input []byte) []byte {
	var blocks [][64]byte

	// preparations
	if len(input) < 56 {
		var currentBlock [64]byte
		copy(currentBlock[:], input)
		currentBlock[len(input)] = 1 << 7               // add 1 bit to the end
		binaryLength := getBinaryLength(len(input) * 8) // get binary length of the input (in 64 bit format)
		// add binary length to the 512 block
		for i, j := 56, 0; i < 64 && j < 8; i, j = i+1, j+1 {
			currentBlock[i] = binaryLength[j]
		}
		blocks = append(blocks, currentBlock)
		// THIS CODE WAS TESTED WITH TEST CASE [http://book.itep.ru/6/sha1.htm]
	}

	return input
}

/*
HELPERS
*/
func getBinaryLength(length int) [8]byte {
	// constant for amount of right shift
	right := [8]int{56, 48, 40, 32, 24, 16, 8, 0}
	var binary64bitlength [8]byte
	for i := 0; i < 8; i++ {
		binary64bitlength[i] = uint8(length >> right[i])
	}
	return binary64bitlength
}
