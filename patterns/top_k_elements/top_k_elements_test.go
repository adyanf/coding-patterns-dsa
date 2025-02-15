package top_k_elements_test

import (
	"testing"

	"github.com/adyanf/coding-patterns-dsa/patterns/top_k_elements"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Case 1",
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6},
			k:        10,
			expected: []int{10, 11, 4, 8, 3, 6, 7, 1, 2, 5},
		},
		{
			name:     "Case 2",
			nums:     []int{5, 1, -1, -8, -7, 8, -5, 0, 1, 10, 8, 0, -4, 3, -1, -1, 4, -5, 4, -3, 0, 2, 2, 2, 4, -2, -4, 8, -7, -7, 2, -8, 0, -8, 10, 8, -8, -2, -9, 4, -7, 6, 6, -1, 4, 2, 8, -3, 5, -9, -3, 6, -8, -5, 5, 10, 2, -5, -1, -5, 1, -3, 7, 0, 8, -2, -3, -1, -5, 4, 7, -9, 0, 2, 10, 4, 4, -4, -1, -1, 6, -8, -9, -1, 9, -9, 3, 5, 1, 6, -1, -2, 4, 2, 4, -6, 4, 4, 5, -5},
			k:        7,
			expected: []int{-8, 0, 8, -5, 2, -1, 4},
		},
		{
			name:     "Case 3",
			nums:     []int{1, 1, 1, 1, 1, 1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "Case 4",
			nums:     []int{6, 0, 1, 4, 9, 7, -3, 1, -4, -8, 4, -7, -3, 3, 2, -3, 9, 5, -4, 0},
			k:        6,
			expected: []int{-4, 0, 1, 4, 9, -3},
		},
		{
			name:     "Case 5",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := top_k_elements.TopKFrequent(test.nums, test.k)
			diff := false
			if len(got) != len(test.expected) {
				diff = true
			} else {
				for i := range got {
					if got[i] != test.expected[i] {
						diff = true
						break
					}
				}
			}
			if diff {
				t.Errorf("TopKFrequent(%v, %d) = %v, want %v", test.nums, test.k, got, test.expected)
			}
		})
	}
}
