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

func EncryptAESWithECBMode(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, len(plaintext))
	blockSize := len(key)

	for start, end := 0, blockSize; start < len(plaintext); start, end = start+blockSize, end+blockSize {
		block.Encrypt(ciphertext[start:end], plaintext[start:end])
	}

	return ciphertext, nil
}

func DecryptAESWithECBMode(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, len(ciphertext))
	blockSize := len(key)

	for start, end := 0, blockSize; start < len(ciphertext); start, end = start+blockSize, end+blockSize {
		block.Decrypt(plaintext[start:end], ciphertext[start:end])
	}

	return plaintext, nil
}
