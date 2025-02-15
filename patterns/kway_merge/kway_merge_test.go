package kway_merge_test

import (
	"testing"

	"github.com/adyanf/coding-patterns-dsa/patterns/kway_merge"
)

func TestFindKSmallestPairs(t *testing.T) {
	tests := []struct {
		name     string
		list1    []int
		list2    []int
		k        int
		expected [][]int
	}{
		{
			name:     "Case 1",
			list1:    []int{1, 2, 300},
			list2:    []int{1, 11, 20, 35, 300},
			k:        30,
			expected: [][]int{{1, 1}, {2, 1}, {1, 11}, {2, 11}, {1, 20}, {2, 20}, {1, 35}, {2, 35}, {1, 300}, {300, 1}, {2, 300}, {300, 11}, {300, 20}, {300, 35}, {300, 300}},
		},
		{
			name:     "Case 2",
			list1:    []int{1, 1, 2},
			list2:    []int{1, 2, 3},
			k:        1,
			expected: [][]int{{1, 1}},
		},
		{
			name:     "Case 3",
			list1:    []int{4, 6},
			list2:    []int{2, 3},
			k:        2,
			expected: [][]int{{4, 2}, {4, 3}},
		},
		{
			name:     "Case 4",
			list1:    []int{4, 7, 9},
			list2:    []int{4, 7, 9},
			k:        5,
			expected: [][]int{{4, 4}, {4, 7}, {7, 4}, {4, 9}, {9, 4}},
		},
		{
			name:     "Case 5",
			list1:    []int{1, 1, 2},
			list2:    []int{1},
			k:        4,
			expected: [][]int{{1, 1}, {1, 1}, {2, 1}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := kway_merge.FindKSmallestPairs(test.list1, test.list2, test.k)

			diff := false
			if len(got) != len(test.expected) {
				diff = true
			} else {
				for i := 0; i < len(got); i++ {
					if got[i][0] != test.expected[i][0] || got[i][1] != test.expected[i][1] {
						diff = true
						break
					}
				}
			}

			if diff {
				t.Errorf("FindKSmallestPairs(%v, %v, %d) = %v, want %v", test.list1, test.list2, test.k, got, test.expected)
			}
		})
	}
}

func TestKthSmallestElement(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		k        int
		expected int
	}{
		{
			name:     "Case 1",
			matrix:   [][]int{{2, 6, 8}, {3, 7, 10}, {5, 8, 11}},
			k:        3,
			expected: 5,
		},
		{
			name:     "Case 2",
			matrix:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:        4,
			expected: 4,
		},
		{
			name:     "Case 3",
			matrix:   [][]int{{1, 4}, {2, 5}},
			k:        4,
			expected: 5,
		},
		{
			name:     "Case 4",
			matrix:   [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			k:        5,
			expected: 1,
		},
		{
			name:     "Case 5",
			matrix:   [][]int{{1, 3, 5, 7, 9}, {2, 4, 6, 8, 10}, {11, 13, 15, 17, 19}, {12, 14, 16, 18, 20}, {21, 22, 23, 24, 25}},
			k:        11,
			expected: 11,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := kway_merge.KthSmallestElement(test.matrix, test.k)
			if got != test.expected {
				t.Errorf("KthSmallestElement(%v, %d) = %d, want %d", test.matrix, test.k, got, test.expected)
			}
		})
	}
}

func TestKSmallestNumber(t *testing.T) {
	tests := []struct {
		name     string
		lists    [][]int
		k        int
		expected int
	}{
		{
			name:     "Case 1",
			lists:    [][]int{{2, 6, 8}, {3, 7, 10}, {5, 8, 11}},
			k:        5,
			expected: 7,
		},
		{
			name:     "Case 2",
			lists:    [][]int{{1, 2, 3}, {4, 5}, {6, 7, 8, 15}, {10, 11, 12, 13}, {5, 10}},
			k:        50,
			expected: 15,
		},
		{
			name:     "Case 3",
			lists:    [][]int{{1, 1, 1}, {1, 1, 1}},
			k:        4,
			expected: 1,
		},
		{
			name:     "Case 4",
			lists:    [][]int{{4, 6}, {2, 3}, {8, 9}},
			k:        10,
			expected: 9,
		},
		{
			name:     "Case 5",
			lists:    [][]int{{5, 8, 9, 17}, {1, 6, 6, 6}, {8, 17, 23, 24}},
			k:        6,
			expected: 8,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := kway_merge.KSmallestNumber(test.lists, test.k)
			if got != test.expected {
				t.Errorf("KSmallestNumber(%v, %d) = %d, want %d", test.lists, test.k, got, test.expected)
			}
		})
	}
}
