package s1

import (
	"bytes"
	"testing"
)

func TestSingleByteXorCipher(t *testing.T) {
	hexEncoded := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	want := []byte("Cooking MC's like a pound of bacon")
	calc, _, err := SingleByteXorCipher(hexEncoded)
	if err != nil {
		t.Errorf("Error in function SingleByteXorCipher: %v", err)
	}
	if !bytes.Equal(want, calc) {
		t.Errorf("SingleByteXorCipher not calculated Correctly")
	}
}
