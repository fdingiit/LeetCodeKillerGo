package toHex

import (
	"testing"
)

func Test(t *testing.T) {
	n := -1
	got := toHex(n)
	wanted := "ffffffff"

	if got != "ffffffff" {
		t.Errorf("wanted: %v, got: %v", wanted, got)
	}
}
