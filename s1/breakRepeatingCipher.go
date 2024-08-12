package s1

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

const minXORKeyLength = 2
const maxXORKeyLength = 40
const numTestChunks = 4
const space = ' '

func editDistance(s1, s2 string) int {
	distance := 0
	for i, c1 := range s1 {
		r1 := fmt.Sprintf("%8b", c1)
		r2 := fmt.Sprintf("%8b", s2[i])
		for j := 0; j < 8; j++ {
			if r1[j] != r2[j] {
				distance++
			}
		}
	}

	return distance
}

func normalizedEditDistance(s1, s2 string) float64 {
	return float64(editDistance(s1, s2)) / float64(len(s1))
}

// Given two characters c1 and c2 which are iid random variables,
// and two characters x and y, it seems intuitive that the expected
// edit distance between (c1 ^ x) and (c2 ^ y) is minimized when
// x = y.  Thus, if we let i vary over the possible repeating XOR
// key lengths, take 4 (consecutive) test chunks of the ciphertext
// of length i, and compute the sum of their normalized pairwise
// edit distances, we expect this sum to be minimized when i is
// any multiple of the true key length.
//
// This is the intuition behind the following heuristic algorithm.
func probableKeyLength(ciphertext string) int {
	keyLength := 0
	bestScore := math.MaxFloat64

	for i := minXORKeyLength; i <= maxXORKeyLength && i*numTestChunks <= len(ciphertext); i++ {
		score := 0.0
		for j := 0; j < numTestChunks-1; j++ {
			for k := j + 1; k < numTestChunks; k++ {
				score += normalizedEditDistance(ciphertext[j*i:j*i+i], ciphertext[k*i:k*i+i])
			}
		}

		if score < bestScore {
			bestScore = score
			keyLength = i
		}
	}

	return keyLength
}

// The most common plaintext character is ' ', thus XOR-ing
// the most common character in some ciphertext with ' ' yields
// the most likely key, assuming the ciphertext was obtained
// from the plaintext via a single character repeating XOR.
func detectSingleCharXORKey(ciphertext string) byte {
	counts := map[byte]int{}
	for _, c := range []byte(ciphertext) {
		if count, ok := counts[c]; ok {
			counts[c] = count + 1
		} else {
			counts[c] = 1
		}
	}

	maxCount := 0
	var mostCommonChar byte
	for c, count := range counts {
		if count > maxCount {
			mostCommonChar = c
			maxCount = count
		}
	}

	return mostCommonChar ^ space
}

func encryptSingleCharXOR(text string, key byte) string {
	textAsBytes := []byte(text)
	encryptedChars := make([]byte, len(textAsBytes))
	for i, c := range textAsBytes {
		encryptedChars[i] = c ^ key
	}
	return string(encryptedChars)
}

func decryptSingleCharXOR(ciphertext string) string {
	key := detectSingleCharXORKey(ciphertext)
	return encryptSingleCharXOR(ciphertext, key)
}

func Decrypt() string {

	file, err := os.Open("./s1/6.txt")
	if err != nil {
		panic(err)
	}
	rawEncodedCiphertext, err := ioutil.ReadAll(file)
	file.Close()
	if err != nil {
		panic(err)
	}
	encodedCiphertext := strings.Replace(string(rawEncodedCiphertext), "\n", "", -1)

	rawCiphertext, err := base64.StdEncoding.DecodeString(encodedCiphertext)
	if err != nil {
		panic(err)
	}
	ciphertext := string(rawCiphertext)

	keyLength := probableKeyLength(ciphertext)

	subCiphertexts := make([]string, keyLength)
	for i, c := range ciphertext {
		subCiphertexts[i%keyLength] += string(c)
	}

	subPlaintexts := make([]string, keyLength)
	for i, sc := range subCiphertexts {
		subPlaintexts[i] = decryptSingleCharXOR(sc)
	}

	plaintext := ""
	for i := 0; i < len(ciphertext); i++ {
		plaintext += string(subPlaintexts[i%keyLength][i/keyLength])
	}

	return plaintext
}
