package main

import (
	"fmt"
	"github.com/heisenberg8055/cryptopals/converter"
)

func main() {
	fmt.Printf("Hex To Base64: %v", converter.HexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))
}
