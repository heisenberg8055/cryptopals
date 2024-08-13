package s1

import (
	"bufio"
	"log"
	"os"
)

func DetectSingleCharacterXOR() string {
	maxScore := 0.00
	var encrypted = make([]byte, 0)
	file, err := os.Open(".s1/4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hexEncoded := scanner.Text()
		decrypted, score, _ := SingleByteXorCipher(hexEncoded)
		if score > maxScore {
			maxScore = score
			encrypted = decrypted
		}
	}
	return string(encrypted)
}
