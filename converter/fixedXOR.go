package converter

import (
	"encoding/hex"
	"fmt"
)

func xor(b1, b2 []byte) ([]byte, error) {
	if len(b1) != len(b2) {
		return nil, fmt.Errorf("length not equal")
	}
	arr := make([]byte, len(b1))
	for i := 0; i < len(b1); i++ {
		arr[i] = b1[i] ^ b2[i]
	}
	return arr, nil
}
func FixedXor(v1, v2 string) string {
	s1, err1 := hex.DecodeString(v1)
	if err1 != nil {
		return ""
	}
	s2, err2 := hex.DecodeString(v2)
	if err2 != nil {
		return ""
	}
	ans, err := xor(s1, s2)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(ans)

}
