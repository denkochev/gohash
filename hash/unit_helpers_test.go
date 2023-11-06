package hash

import "testing"

func Test_getBinaryLength(t *testing.T) {
	tests := []struct {
		input  []byte
		result [8]byte
	}{
		{[]byte{97, 98, 99, 100, 101}, [8]byte{0, 0, 0, 0, 0, 0, 0, 40}},
	}

	for _, test := range tests {
		testInput := test.input

		result := getBinaryLength(len(testInput) * 8)

		if result != test.result {
			t.Errorf("For: %d Expected:- 0x%X, but got - 0x%X", test.input, test.result, result)
		}
	}
}
