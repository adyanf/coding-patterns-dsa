package merge_intervals_test

import (
	"fmt"
	"testing"

	"github.com/adyanf/coding-patterns-dsa/patterns/merge_intervals"
)

func TestLeastTime(t *testing.T) {
	testCases := []struct {
		name     string
		tasks    []byte
		n        int
		expected int
	}{
		{
			name:     "Case 1",
			tasks:    []byte{'A', 'A', 'B', 'B'},
			n:        2,
			expected: 5,
		},
		{
			name:     "Case 2",
			tasks:    []byte{'A', 'A', 'A', 'B', 'B', 'C', 'C'},
			n:        1,
			expected: 7,
		},
		{
			name:     "Case 3",
			tasks:    []byte{'S', 'I', 'V', 'U', 'W', 'D', 'U', 'X'},
			n:        0,
			expected: 8,
		},
		{
			name:     "Case 4",
			tasks:    []byte{'A', 'K', 'X', 'M', 'W', 'D', 'X', 'B', 'D', 'C', 'O', 'Z', 'D', 'E', 'Q'},
			n:        3,
			expected: 15,
		},
		{
			name:     "Case 5",
			tasks:    []byte{'A', 'B', 'C', 'O', 'Q', 'C', 'Z', 'O', 'X', 'C', 'W', 'Q', 'Z', 'B', 'M', 'N', 'R', 'L', 'C', 'J'},
			n:        10,
			expected: 34,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := merge_intervals.LeastTime(test.tasks, test.n)
			if result != test.expected {
				t.Errorf("LeastTime(%s, %d) = %v, want %v", test.tasks, test.n, result, test.expected)
			}
		})
	}
}

func TestMergeIntervals(t *testing.T) {
	testCases := []struct {
		name      string
		intervals [][]int
		expected  [][]int
	}{
		{
			name:      "Case 1",
			intervals: [][]int{{1, 5}, {3, 7}, {4, 6}},
			expected:  [][]int{{1, 7}},
		},
		{
			name:      "Case 2",
			intervals: [][]int{{1, 5}, {4, 6}, {6, 8}, {11, 15}},
			expected:  [][]int{{1, 8}, {11, 15}},
		},
		{
			name:      "Case 3",
			intervals: [][]int{{1, 5}},
			expected:  [][]int{{1, 5}},
		},
		{
			name:      "Case 4",
			intervals: [][]int{{1, 9}, {3, 8}, {4, 4}},
			expected:  [][]int{{1, 9}},
		},
		{
			name:      "Case 5",
			intervals: [][]int{{1, 2}, {3, 4}, {8, 8}},
			expected:  [][]int{{1, 2}, {3, 4}, {8, 8}},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := merge_intervals.MergeIntervals(test.intervals)
			if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", test.expected) {
				t.Errorf("MergeIntervals(%v) = %v, want %v", test.intervals, result, test.expected)
			}
		})
	}
}

func TestInsertInterval(t *testing.T) {
	testCases := []struct {
		name     string
		existing [][]int
		new      []int
		expected [][]int
	}{
		{
			name:     "Case 1",
			existing: [][]int{{1, 2}, {3, 4}, {5, 8}, {9, 15}},
			new:      []int{2, 5},
			expected: [][]int{{1, 8}, {9, 15}},
		},
		{
			name:     "Case 2",
			existing: [][]int{{1, 6}, {8, 9}, {10, 15}, {16, 18}},
			new:      []int{9, 10},
			expected: [][]int{{1, 6}, {8, 15}, {16, 18}},
		},
		{
			name:     "Case 3",
			existing: [][]int{{1, 2}, {3, 4}, {5, 8}, {9, 15}},
			new:      []int{16, 17},
			expected: [][]int{{1, 2}, {3, 4}, {5, 8}, {9, 15}, {16, 17}},
		},
		{
			name:     "Case 4",
			existing: [][]int{{1, 4}, {5, 6}, {7, 8}, {9, 10}},
			new:      []int{1, 5},
			expected: [][]int{{1, 6}, {7, 8}, {9, 10}},
		},
		{
			name:     "Case 5",
			existing: [][]int{{1, 3}, {4, 6}, {7, 8}, {9, 10}},
			new:      []int{1, 10},
			expected: [][]int{{1, 10}},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := merge_intervals.InsertInterval(test.existing, test.new)
			if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", test.expected) {
				t.Errorf("InsertInterval(%v, %v) = %v, want %v", test.existing, test.new, result, test.expected)
			}
		})
	}
}
