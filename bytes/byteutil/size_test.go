package byteutil

import (
	"testing"
)

func TestSize(t *testing.T) {
	var s Size = GB + 200*MB

	if s.String() != "1.195GB" {
		t.Fatalf("%s != 1.195GB", s.String())
	}
}
