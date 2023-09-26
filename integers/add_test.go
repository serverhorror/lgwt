package integers

import "testing"

func TestAdder(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{2, 2, 4},
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
		{-1, -1, -2},
	}
	for _, tc := range testCases {
		sum := Add(tc.a, tc.b)
		if sum != tc.expected {
			t.Errorf("expected '%d' but got '%d'", tc.expected, sum)
		}
	}
}
