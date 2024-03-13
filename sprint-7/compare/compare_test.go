package compare

import (
	"testing"
)

func TestMaxInt(t *testing.T) {
	a, b := 2, 7

	max := MaxInt(a, b)

	if max != b {
		t.Errorf("expected %d, got %d", b, max)
	}
}
