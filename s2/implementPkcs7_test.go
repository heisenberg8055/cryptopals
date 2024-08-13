package s2

import (
	"testing"
)

func TestPKCS7(t *testing.T) {
	t.Run("Padding Needed", func(t *testing.T) {
		want := "YELLOW SUBMARINE\x04\x04\x04\x04"
		calc := PKCS7([]byte("YELLOW SUBMARINE"), 20)

		if string(want) != string(calc) {
			t.Error("PKCS calculation Failed")
		}
	})
	t.Run("No Padding Needed", func(t *testing.T) {
		want := "YELLOW SUBMARINE"
		calc := PKCS7([]byte("YELLOW SUBMARINE"), 16)

		if string(want) != string(calc) {
			t.Error("PKCS calculation Failed")
		}
	})
}
