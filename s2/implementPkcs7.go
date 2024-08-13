package s2

func PKCS7(block []byte, size int) []byte {
	var n int = len(block)
	if n%size == 0 {
		return block
	}

	var blockSize int = (n / size) + 1
	paddingSize := (blockSize * size) - n

	totalLength := blockSize * size

	padded := make([]byte, totalLength)

	for i := 0; i < n; i++ {
		padded[i] = block[i]
	}
	for i := totalLength - 1; i >= totalLength-paddingSize; i-- {
		padded[i] = byte(paddingSize)
	}
	return padded
}
