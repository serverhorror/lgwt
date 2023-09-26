package arrays_and_slices

import "testing"

func TestSumAll(t *testing.T) {
	testCases := []struct {
		numbers  [][]int
		expected []int
	}{
		{[][]int{{1, 2}, {3, 4}}, []int{3, 7}},
		{[][]int{{1, 1, 1}, {0, 0, 0}}, []int{3, 0}},
		{[][]int{{}, {0, 0, 0}}, []int{0, 0}},
		{[][]int{{1, 2, 3}, {}}, []int{6, 0}},
		{[][]int{{}, {}}, []int{0, 0}},
	}
	for _, tc := range testCases {
		got := SumAll(tc.numbers...)
		if len(got) != len(tc.expected) {
			t.Errorf("got %v want %v given, %v", got, tc.expected, tc.numbers)
		}
		for i := 0; i < len(tc.expected); i++ {
			if got[i] != tc.expected[i] {
				t.Errorf("got %v want %v given, %v", got, tc.expected, tc.numbers)
			}
		}
	}
}

func TestSum(t *testing.T) {

	testCases := []struct {
		numbers  []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{1, 2, 3, 4, 5, 1}, 16},
		{[]int{0}, 0},
		{[]int{}, 0},
	}

	for _, tc := range testCases {
		got := Sum(tc.numbers)
		if got != tc.expected {
			t.Errorf("got %d want %d given, %v", got, tc.expected, tc.numbers)
		}
	}
}
