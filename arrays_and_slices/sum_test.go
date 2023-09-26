package arrays_and_slices

import "testing"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
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
