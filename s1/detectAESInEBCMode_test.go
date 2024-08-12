package s1

import (
	"testing"
)

func TestDetectAESInECBModeFile(t *testing.T) {
	keySize := 16

	candidate, _ := DetectAESInECBModeFile("8.txt", keySize)

	if candidate == nil {
		t.Errorf("Expected to find a candidate")
	}
}
