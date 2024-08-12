package s1

import (
	"bufio"
	"encoding/base64"
	"log"
	"os"
)

func BreakRepeatingCipher() {
	file, err := os.Open("/home/heisenberg/Documents/gooo/cryptopals/s1/6.txt")
	if err != nil {
		log.Fatalf("Error during file import: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		base64Encoded := scanner.Text()
		_, err := base64.StdEncoding.DecodeString(base64Encoded)
		if err != nil {
			log.Fatalf("Error during Base64 conversion: %v", err)
		}
	}
}

func HammingDistance(buffer1, buffer2 []byte) int {
	ans := 0
	maxLen := max(len(buffer1), len(buffer2))
	buffer3 := make([]byte, maxLen)
	for i := range maxLen {
		val1 := byte(0)
		val2 := byte(0)
		if i < len(buffer1) {
			val1 = buffer1[i]
		}
		if i < len(buffer2) {
			val2 = buffer2[i]
		}
		buffer3[i] = val1 ^ val2
	}

	oneSetBitByteSlice := []byte{
		0b00000001,
		0b00000010,
		0b00000100,
		0b00001000,
		0b00010000,
		0b00100000,
		0b01000000,
		0b10000000,
	}

	for _, b := range buffer3 {
		for _, oneSetBitByte := range oneSetBitByteSlice {
			if (b & oneSetBitByte) > 0 {
				ans += 1
			}
		}
	}
	return ans
}
