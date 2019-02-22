package monotoneIncreasingDigits

import (
	"testing"
)

var check = map[int]int{
	0:     0,
	3:     3,
	10:    9,
	100:   99,
	332:   299,
	1231:  1229,
	1234:  1234,
	10100: 9999,
	11924: 11899,
	28951: 28899,
	58478: 57999,
}

func Test(t *testing.T) {
	for k, v := range check {
		got := monotoneIncreasingDigits(k)

		if got != v {
			t.Errorf("wanted: %v, but got: %v", v, got)
		}
	}
}
