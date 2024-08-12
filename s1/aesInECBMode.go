package s1

import (
	"crypto/aes"
	"encoding/base64"
	"io"
	"os"
)

func AESInECB() []byte {
	file, err := os.Open("./s1/7.txt")
	if err != nil {
		panic(err)
	}
	base64Encrypted, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	base64Decrypted, err := base64.StdEncoding.DecodeString(string(base64Encrypted))
	if err != nil {
		panic(err)
	}
	decrypted, err := DecryptECB(base64Decrypted, []byte("YELLOW SUBMARINE"))
	if err != nil {
		panic(err)
	}
	return decrypted
}

func DecryptECB(encoded, key []byte) ([]byte, error) {

	decoded := make([]byte, len(encoded))
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(encoded); i += 16 {
		cipher.Decrypt(decoded[i:i+16], encoded[i:i+16])
	}
	return decoded, nil
}
