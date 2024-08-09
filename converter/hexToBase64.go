package converter

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func HexToBase64(v string) string {
	decoder, err := hex.DecodeString(v)
	if err != nil {
		log.Fatal(err)
	}
	ans := base64.StdEncoding.EncodeToString(decoder)
	return ans
}
