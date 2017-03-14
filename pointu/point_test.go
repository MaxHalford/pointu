package pointu

import (
	"testing"
)

func TestSortByX(t *testing.T) {
	var tests = []struct {
		initial  Points
		expected Points
	}{
		{
			initial:  Points{{5, 0}, {1, 0}, {3, 0}, {2, 0}, {4, 0}},
			expected: Points{{1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}},
		},
	}
	for _, test := range tests {
		test.initial.sortByX()
		for i := range test.initial {
			if test.initial[i] != test.expected[i] {
				t.Error("sortByX didn't work as expected")
			}
		}
	}
}

func TestSortByY(t *testing.T) {
	var tests = []struct {
		initial  Points
		expected Points
	}{
		{
			initial:  Points{{0, 5}, {0, 1}, {0, 3}, {0, 2}, {0, 4}},
			expected: Points{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}},
		},
	}
	for _, test := range tests {
		test.initial.sortByY()
		for i := range test.initial {
			if test.initial[i] != test.expected[i] {
				t.Error("sortByY didn't work as expected")
			}
		}
	}
}
