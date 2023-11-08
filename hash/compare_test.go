package hash

import (
	"crypto/sha1"
	"encoding/csv"
	"fmt"
	"os"
	"testing"
)

type DataToHash struct {
	bytesSlice string
}

func Test_CompareWithCryptoSHA1Lib(t *testing.T) {
	// TEST ON 30 MB CSV OF DIFFERENT STRINGS
	tests, _ := readCSV("./test_cases/large_strings.csv")

	for _, test := range tests {
		testCase := test.bytesSlice
		// mine
		h := Hash{}
		h.New([]byte(testCase))
		myHash := fmt.Sprintf("%x", h.Get())

		// library
		hLib := sha1.New()
		hLib.Write([]byte(testCase))
		libHash := fmt.Sprintf("%x", hLib.Sum(nil))
		if myHash != libHash {
			t.Errorf("For: %s Expected:-%s, but got - %s", testCase, libHash, myHash)
		}
	}

	// TEST ON SAME CSV BUT AS ONE FILE INSTEAD OF EACH STRINGS
	for i := 0; i <= 1; i++ {
		// mine
		h := Hash{}
		// BUILD IN METHOD FOR FILE READING BUT YOU ALSO CAN USE DEFAULT .NEW() METHOD WITH []byte ARG FORMAT
		h.NewFile("./hash/test_cases/large_strings.csv")
		myHash := fmt.Sprintf("%x", h.Get())

		// library
		csv, _ := os.ReadFile("./hash/test_cases/large_strings.csv")
		hLib := sha1.New()
		hLib.Write([]byte(csv))
		libHash := fmt.Sprintf("%x", hLib.Sum(nil))
		if myHash != libHash {
			t.Errorf("For: ./hash/test_cases/large_strings.csv Expected:-%s, but got - %s", libHash, myHash)
		}
	}

	// TEST ON 60MB FILE [for test i've used Go-64 windows compiler]
	for i := 0; i <= 1; i++ {
		// mine
		h := Hash{}
		// BUILD IN METHOD FOR FILE READING BUT YOU ALSO CAN USE DEFAULT .NEW() METHOD WITH []byte ARG FORMAT
		h.NewFile("./hash/test_cases/go1.21.1.windows-amd64.msi")
		myHash := fmt.Sprintf("%x", h.Get())

		// library
		hugefile, _ := os.ReadFile("./hash/test_cases/go1.21.1.windows-amd64.msi")
		hLib := sha1.New()
		hLib.Write([]byte(hugefile))
		libHash := fmt.Sprintf("%x", hLib.Sum(nil))
		if myHash != libHash {
			t.Errorf("For: ./hash/test_cases/go1.21.1.windows-amd64.msi Expected:-%s, but got - %s", libHash, myHash)
		}
	}

	// TEST ON .txt PLAIN TEXT FILE
	for i := 0; i <= 1; i++ {
		// mine
		h := Hash{}
		// BUILD IN METHOD FOR FILE READING BUT YOU ALSO CAN USE DEFAULT .NEW() METHOD WITH []byte ARG FORMAT
		h.NewFile("./hash/test_cases/big.txt")
		myHash := fmt.Sprintf("%x", h.Get())

		// library
		hugefile, _ := os.ReadFile("./hash/test_cases/big.txt")
		hLib := sha1.New()
		hLib.Write([]byte(hugefile))
		libHash := fmt.Sprintf("%x", hLib.Sum(nil))
		if myHash != libHash {
			t.Errorf("For: ./hash/test_cases/big.txt Expected:-%s, but got - %s", libHash, myHash)
		}
	}

	// TEST ON RAW BINARY FILE
	for i := 0; i <= 1; i++ {
		// mine
		h := Hash{}
		// BUILD IN METHOD FOR FILE READING BUT YOU ALSO CAN USE DEFAULT .NEW() METHOD WITH []byte ARG FORMAT
		h.NewFile("./hash/test_cases/binary_file.bin")
		myHash := fmt.Sprintf("%x", h.Get())

		// library
		hugefile, _ := os.ReadFile("./hash/test_cases/binary_file.bin")
		hLib := sha1.New()
		hLib.Write([]byte(hugefile))
		libHash := fmt.Sprintf("%x", hLib.Sum(nil))
		if myHash != libHash {
			t.Errorf("For: ./hash/test_cases/binary_file.bin Expected:-%s, but got - %s", libHash, myHash)
		}
	}

	// TEST ON EXECUTEBLE FILE
	for i := 0; i <= 1; i++ {
		// mine
		h := Hash{}
		// BUILD IN METHOD FOR FILE READING BUT YOU ALSO CAN USE DEFAULT .NEW() METHOD WITH []byte ARG FORMAT
		h.NewFile("./hash/test_cases/gohash.exe")
		myHash := fmt.Sprintf("%x", h.Get())

		// library
		hugefile, _ := os.ReadFile("./hash/test_cases/gohash.exe")
		hLib := sha1.New()
		hLib.Write([]byte(hugefile))
		libHash := fmt.Sprintf("%x", hLib.Sum(nil))
		if myHash != libHash {
			t.Errorf("For: gohash.exe Expected:-%s, but got - %s", libHash, myHash)
		}
	}

	// ZERO BIT TEST
	for i := 0; i <= 1; i++ {
		// mine
		h := Hash{}
		// BUILD IN METHOD FOR FILE READING BUT YOU ALSO CAN USE DEFAULT .NEW() METHOD WITH []byte ARG FORMAT
		h.New([]byte{})
		myHash := fmt.Sprintf("%x", h.Get())

		// library
		hLib := sha1.New()
		hLib.Write([]byte{})
		libHash := fmt.Sprintf("%x", hLib.Sum(nil))
		if myHash != libHash {
			t.Errorf("For: %v Expected:-%s, but got - %s", []byte{}, libHash, myHash)
		}
	}
}

func readCSV(filename string) ([]DataToHash, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var records []DataToHash
	for _, line := range lines[1:] { // Skip header
		records = append(records, DataToHash{
			bytesSlice: line[0],
		})
		records = append(records, DataToHash{
			bytesSlice: line[1],
		})
		records = append(records, DataToHash{
			bytesSlice: line[2],
		})
		records = append(records, DataToHash{
			bytesSlice: line[3],
		})
		records = append(records, DataToHash{
			bytesSlice: line[4],
		})
	}
	return records, nil
}
