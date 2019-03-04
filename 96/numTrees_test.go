package numTrees

import "testing"

func TestCase0(t *testing.T) {
	res := map[int]int{
		1: 1,
		2: 2,
		3: 5,
		4: 14,
	}

	for k, v := range res {
		g := numTrees(k)
		if g != v {
			t.Errorf("%d shoule get %d, but get %d", k, v, g)
		}
	}
}
