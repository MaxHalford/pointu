package pointu

import "testing"

func TestBisectLeftInt(t *testing.T) {
	var tests = []struct {
		value int
		ints  []int
		index int
	}{
		{
			value: -1,
			ints:  []int{0, 2, 4, 6, 8, 10},
			index: 0,
		},
		{
			value: 0,
			ints:  []int{0, 2, 4, 6, 8, 10},
			index: 0,
		},
		{
			value: 1,
			ints:  []int{0, 2, 4, 6, 8, 10},
			index: 1,
		},
		{
			value: 2,
			ints:  []int{0, 2, 4, 6, 8, 10},
			index: 1,
		},
		{
			value: 3,
			ints:  []int{0, 2, 4, 6, 8, 10},
			index: 2,
		},
		{
			value: 5,
			ints:  []int{0, 2, 4, 6, 8, 10},
			index: 3,
		},
		{
			value: 9,
			ints:  []int{0, 2, 4, 6, 8, 10},
			index: 5,
		},
		{
			value: 10,
			ints:  []int{0, 2, 4, 6, 8, 10},
			index: 5,
		},
		{
			value: 20,
			ints:  []int{0, 2, 4, 6, 8, 10},
			index: -1,
		},
	}
	for _, test := range tests {
		if bisectLeftInt(test.value, test.ints) != test.index {
			t.Error("bisectLeftInt didn't work as expected")
		}
	}
}
