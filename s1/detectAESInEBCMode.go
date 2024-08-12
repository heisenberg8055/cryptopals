package s1

import (
	"bufio"
	"encoding/hex"
	"os"
)

func DetectAESInECBMode(ciphertext []byte, keySize int) bool {
	blocks := make(map[string]bool)

	for i := 0; i < len(ciphertext); i += keySize {
		block := string(ciphertext[i : i+keySize])
		if blocks[block] {
			return true
		}
		blocks[block] = true
	}

	return false
}

func DetectAESInECBModeFile(filename string, keySize int) ([]byte, error) {
	lines, err := readFileLines(filename)
	if err != nil {
		return nil, err
	}
	for _, line := range lines {
		decoded, err := hex.DecodeString(line)
		if err != nil {
			return nil, err
		}

		if DetectAESInECBMode(decoded, keySize) {
			return decoded, nil

		}

	}

	return nil, nil
}

func readFileLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
