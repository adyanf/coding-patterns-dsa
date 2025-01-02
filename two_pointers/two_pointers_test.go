package two_pointers_test

import (
	"testing"

	"github.com/adyanf/coding-patterns-dsa/two_pointers"
)

func TestFindSumOfThree(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected bool
	}{
		{
			name:     "Case 1",
			nums:     []int{1, -1, 0},
			target:   -1,
			expected: false,
		},
		{
			name:     "Case 2",
			nums:     []int{3, 7, 1, 2, 8, 4, 5},
			target:   10,
			expected: true,
		},
		{
			name:     "Case 3",
			nums:     []int{3, 7, 1, 2, 8, 4, 5},
			target:   21,
			expected: false,
		},
		{
			name:     "Case 4",
			nums:     []int{-1, 2, 1, -4, 5, -3},
			target:   -8,
			expected: true,
		},
		{
			name:     "Case 5",
			nums:     []int{-1, 2, 1, -4, 5, -3},
			target:   0,
			expected: true,
		},
	}

	for _, tc := range testCases {
		got := two_pointers.FindSumOfThree(tc.nums, tc.target)
		if got != tc.expected {
			t.Errorf("FindSumOfThree(%v, %v) = %v, expected %v", tc.nums, tc.target, got, tc.expected)
		}
	}
}
