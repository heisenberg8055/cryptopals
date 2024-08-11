package converter

import (
	"encoding/hex"
	"regexp"
	"strings"
)

var letterFrequencies = map[string]float64{
	"A": 8.167,
	"B": 1.492,
	"C": 2.782,
	"D": 4.253,
	"E": 12.702,
	"F": 2.228,
	"G": 2.015,
	"H": 6.094,
	"I": 6.966,
	"J": 0.153,
	"K": 0.772,
	"L": 4.025,
	"M": 2.406,
	"N": 6.749,
	"O": 7.507,
	"P": 1.929,
	"Q": 0.095,
	"R": 5.987,
	"S": 6.327,
	"T": 9.056,
	"U": 2.758,
	"V": 0.978,
	"W": 2.360,
	"X": 0.150,
	"Y": 1.974,
	"Z": 0.074,
}

func singleXOR(buffer []byte, key byte) []byte {
	ans := make([]byte, len(buffer))
	for i := range buffer {
		ans[i] = buffer[i] ^ key

	}
	return ans
}

func SingleByteXorCipher(hexCode string) ([]byte, error, float64) {
	hexDecoded, err := hex.DecodeString(hexCode)
	if err != nil {
		return nil, err, 0.00
	}
	bestKey := byte(0)
	bestScore := 0.00
	for k := range 256 {
		key := byte(k)
		decoded := singleXOR(hexDecoded, key)
		score := calculateScore(decoded)
		if score > bestScore {
			bestKey = key
			bestScore = score
		}
	}
	return singleXOR(hexDecoded, bestKey), nil, bestScore
}

var reff = regexp.MustCompile("^[a-zA-Z ]$")

func calculateScore(decoded []byte) float64 {
	score := 0.00
	for _, val := range decoded {
		str := string(val)
		if reff.MatchString(str) {
			score += letterFrequencies[strings.ToUpper(str)]
		} else {
			score -= 10.00
		}
	}
	return score
}
