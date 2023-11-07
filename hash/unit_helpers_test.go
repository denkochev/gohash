package hash

import (
	"os"
	"testing"
)

func Test_getBinaryLength(t *testing.T) {
	bigtext, _ := os.ReadFile("./test_cases/big.txt")

	tests := []struct {
		input  []byte
		result [8]byte
	}{
		{[]byte{97, 98, 99, 100, 101}, [8]byte{0, 0, 0, 0, 0, 0, 0, 40}},
		{[]byte{97}, [8]byte{0, 0, 0, 0, 0, 0, 0, 8}},
		{[]byte{1, 2}, [8]byte{0, 0, 0, 0, 0, 0, 0, 16}},
		{[]byte{97, 98, 99, 100, 101, 97, 98, 99, 100, 101, 97, 98, 99, 100, 101, 97, 98, 99, 100, 101, 97, 98, 99, 100, 101}, [8]byte{0, 0, 0, 0, 0, 0, 0, 200}},
		{[]byte{0, 0, 0, 0, 1}, [8]byte{0, 0, 0, 0, 0, 0, 0, 40}},
		{bigtext, [8]byte{0, 0, 0, 0, 3, 24, 18, 208}},
		{[]byte("asdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasdddddddddddasdddddddddddddddddasddddddddddd"), [8]byte{0, 0, 0, 0, 0, 0, 99, 0}},
	}

	for _, test := range tests {
		testInput := test.input

		result := getBinaryLength(len(testInput) * 8)

		if result != test.result {
			t.Errorf("For: %d Expected:- 0x%X, but got - 0x%X", test.input, test.result, result)
		}
	}
}
