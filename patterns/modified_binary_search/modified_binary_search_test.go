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

func TestFindClosestElements(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		target   int
		expected []int
	}{
		{
			name:     "Case 1",
			nums:     []int{1, 2, 3, 4, 5},
			k:        4,
			target:   3,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Case 2",
			nums:     []int{1, 2, 3, 4, 5},
			k:        4,
			target:   -1,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Case 3",
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        5,
			target:   7,
			expected: []int{3, 4, 5, 6, 7},
		},
		{
			name:     "Case 4",
			nums:     []int{-29, -11, -3, 0, 5, 10, 50, 63, 198},
			k:        6,
			target:   8,
			expected: []int{-29, -11, -3, 0, 5, 10},
		},
		{
			name:     "Case 5",
			nums:     []int{-10, -6, -4, -3},
			k:        2,
			target:   5,
			expected: []int{-4, -3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := modified_binary_search.FindClosestElement(test.nums, test.k, test.target)
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
				t.Errorf("FindClosestElement(%v,%d,%d) = %v, want %v", test.nums, test.k, test.target, got, test.expected)
			}
		})
	}
}

func TestRotatedBinarySearchIterative(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Case 1",
			nums:     []int{6, 7, 1, 2, 3, 4, 5},
			target:   3,
			expected: 4,
		},
		{
			name:     "Case 2",
			nums:     []int{6, 7, 1, 2, 3, 4, 5},
			target:   6,
			expected: 0,
		},
		{
			name:     "Case 3",
			nums:     []int{4, 5, 6, 1, 2, 3},
			target:   3,
			expected: 5,
		},
		{
			name:     "Case 4",
			nums:     []int{4, 5, 6, 1, 2, 3},
			target:   6,
			expected: 2,
		},
		{
			name:     "Case 5",
			nums:     []int{4},
			target:   1,
			expected: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := modified_binary_search.RotatedBinarySearchIterative(test.nums, test.target)
			if got != test.expected {
				t.Errorf("RotatedBinarySearchIterative(%v,%d) = %d, want %d", test.nums, test.target, got, test.expected)
			}
		})
	}
}

func TestRotatedBinarySearchRecursive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Case 1",
			nums:     []int{6, 7, 1, 2, 3, 4, 5},
			target:   3,
			expected: 4,
		},
		{
			name:     "Case 2",
			nums:     []int{6, 7, 1, 2, 3, 4, 5},
			target:   6,
			expected: 0,
		},
		{
			name:     "Case 3",
			nums:     []int{4, 5, 6, 1, 2, 3},
			target:   3,
			expected: 5,
		},
		{
			name:     "Case 4",
			nums:     []int{4, 5, 6, 1, 2, 3},
			target:   6,
			expected: 2,
		},
		{
			name:     "Case 5",
			nums:     []int{4},
			target:   1,
			expected: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := modified_binary_search.RotatedBinarySearchRecursive(test.nums, test.target)
			if got != test.expected {
				t.Errorf("RotatedBinarySearchRecursive(%v,%d) = %d, want %d", test.nums, test.target, got, test.expected)
			}
		})
	}
}
