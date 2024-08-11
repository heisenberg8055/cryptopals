package main

import (
	"fmt"

	"github.com/heisenberg8055/cryptopals/converter"
)

func main() {
	expected := []byte("Cooking MC's like a pound of bacon")
	val, _ := converter.SingleByteXorCipher("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	fmt.Print(val, expected)
}
