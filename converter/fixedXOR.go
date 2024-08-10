package converter

import (
	"encoding/hex"
	"fmt"
	"log"
)

func FixedXor(value string) {
	decoded, err := hex.DecodeString(value)
	if err != nil {
		log.Fatal("Fixed XOR Failed")
	}
	fmt.Printf("FixedXOR: %v", string(decoded))
}
