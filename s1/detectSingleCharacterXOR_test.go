package s1

import (
	"testing"
)

func TestDetectSingleCharacterXOR(t *testing.T) {
	want := "Now that the party is jumping\n"
	calc := DetectSingleCharacterXOR()
	if calc != want {
		t.Errorf("DetectSingleCharacterXOR Calculation Not Worked")
	}
}
