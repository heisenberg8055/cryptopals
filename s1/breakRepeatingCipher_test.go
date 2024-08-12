package s1

import (
	"testing"
)

func TestHammingDistance(t *testing.T) {
	want := 37
	calc := editDistance("this is a test", "wokka wokka!!!")
	if want != calc {
		t.Error("Calculation not worked")
	}
}
