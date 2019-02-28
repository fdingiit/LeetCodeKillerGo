package canFinish

import "testing"

func TestCase0(t *testing.T) {
	if !canFinish(2, [][]int{{1, 0}}) {
		t.Errorf("should pass")
	}
}

func TestCase1(t *testing.T) {
	if canFinish(2, [][]int{{1, 0}, {0, 1}}) {
		t.Errorf("should not pass")
	}
}

func TestCase2(t *testing.T) {
	courses := [][]int{
		{5, 2},
		{2, 3},
		{3, 1},
		{4, 1},
		{5, 0},
		{4, 0},
	}

	if !canFinish(6, courses) {
		t.Errorf("should pass")
	}
}

func TestCase3(t *testing.T) {
	courses := [][]int{
		{5, 2},
		{2, 3},
		{3, 1},
		{4, 1},
		{5, 0},
		{4, 0},
		{1, 2},
	}

	if canFinish(6, courses) {
		t.Errorf("should not pass")
	}
}

func TestCase4(t *testing.T) {
	courses := [][]int{
		{0, 1},
		{1, 3},
		{2, 3},
		{0, 4},
		{4, 3},
		{3, 5},
	}

	if !canFinish(6, courses) {
		t.Errorf("should pass")
	}
}
