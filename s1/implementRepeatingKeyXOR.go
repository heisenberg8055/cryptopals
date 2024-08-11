package s1

import "encoding/hex"

func ImplementRepeatingKeyXOR(buffer1 string) string {
	key := "ICE"
	return hex.EncodeToString(repeatingXor(buffer1, key))
}
func repeatingXor(buffer, key string) []byte {
	n := len(key)
	arr := make([]byte, len(buffer))
	for i := range buffer {
		arr[i] = buffer[i] ^ key[i%n]
	}
	return arr
}
