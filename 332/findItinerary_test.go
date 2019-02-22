package findItinerary

import "testing"

func match(l1, l2 []string) bool {
	if len(l1) != len(l2) {
		return false
	}

	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			return false
		}
	}

	return true
}

func TestCase1(t *testing.T) {
	input := [][]string{
		{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"},
	}
	output := []string{
		"JFK", "MUC", "LHR", "SFO", "SJC",
	}

	got := findItinerary(input)

	if !match(got, output) {
		t.Errorf("wanted: %v, but got :%v", output, got)
	}
}

func TestCase2(t *testing.T) {
	input := [][]string{
		{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"},
	}
	output := []string{
		"JFK", "ATL", "JFK", "SFO", "ATL", "SFO",
	}

	got := findItinerary(input)

	if !match(got, output) {
		t.Errorf("wanted: %v, but got :%v", output, got)
	}
}
