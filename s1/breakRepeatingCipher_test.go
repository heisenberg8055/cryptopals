package s1

import (
	"testing"
)

func TestHammingDistance(t *testing.T) {
	want := 37
	calc := HammingDistance([]byte("this is a test"), []byte("wokka wokka!!!"))
	if want != calc {
		t.Error("Calculation not worked")
	}
}
