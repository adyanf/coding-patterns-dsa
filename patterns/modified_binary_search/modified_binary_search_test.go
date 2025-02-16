package modified_binary_search_test

import (
	"testing"

	"github.com/adyanf/coding-patterns-dsa/patterns/modified_binary_search"
)

func TestSingleNonDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Case 1",
			nums:     []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 8, 8},
			expected: 5,
		},
		{
			name:     "Case 2",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "Case 3",
			nums:     []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6},
			expected: 6,
		},
		{
			name:     "Case 4",
			nums:     []int{1, 1, 4, 4, 7, 7, 10, 10, 13, 13, 16, 16, 19, 19, 22, 22, 25},
			expected: 25,
		},
		{
			name:     "Case 5",
			nums:     []int{0, 0, 1, 1, 2, 2, 4, 8, 8, 16, 16, 32, 32},
			expected: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := modified_binary_search.SingleNonDuplicate(test.nums)
			if got != test.expected {
				t.Errorf("SingleNonDuplicate(%v) = %d, want %d", test.nums, got, test.expected)
			}
		})
	}
}
