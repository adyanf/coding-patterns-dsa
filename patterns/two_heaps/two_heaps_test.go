package twoheaps_test

import (
	"testing"

	twoheaps "github.com/adyanf/coding-patterns-dsa/patterns/two_heaps"
)

func TestMostBooked(t *testing.T) {
	tests := []struct {
		name     string
		meetings [][]int
		rooms    int
		expected int
	}{
		{
			name:     "Case 1",
			meetings: [][]int{{0, 5}, {1, 6}, {6, 7}, {7, 8}, {8, 9}},
			rooms:    2,
			expected: 0,
		},
		{
			name:     "Case 2",
			meetings: [][]int{{0, 10}, {1, 11}, {2, 12}, {3, 13}, {4, 14}, {5, 15}},
			rooms:    3,
			expected: 0,
		},
		{
			name:     "Case 3",
			meetings: [][]int{{0, 9}, {1, 2}, {2, 3}, {3, 4}, {5, 6}, {6, 7}, {7, 8}, {8, 9}},
			rooms:    3,
			expected: 1,
		},
		{
			name:     "Case 4",
			meetings: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}},
			rooms:    1,
			expected: 0,
		},
		{
			name:     "Case 5",
			meetings: [][]int{{0, 4}, {1, 3}, {2, 4}, {3, 5}, {4, 6}, {5, 7}},
			rooms:    4,
			expected: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := twoheaps.MostBooked(test.meetings, test.rooms); got != test.expected {
				t.Errorf("MostBooked(%v, %v) = %v, want %v", test.meetings, test.rooms, got, test.expected)
			}
		})
	}
}

func TestMinimumMachines(t *testing.T) {
	tests := []struct {
		name     string
		tasks    [][]int
		expected int
	}{
		{
			name:     "Case 1",
			tasks:    [][]int{{1, 1}, {5, 5}, {8, 8}, {4, 4}, {6, 6}, {10, 10}, {7, 7}},
			expected: 1,
		},
		{
			name:     "Case 2",
			tasks:    [][]int{{1, 7}, {1, 7}, {1, 7}, {1, 7}, {1, 7}, {1, 7}},
			expected: 6,
		},
		{
			name:     "Case 3",
			tasks:    [][]int{{1, 7}, {8, 13}, {5, 6}, {10, 14}, {6, 7}},
			expected: 2,
		},
		{
			name:     "Case 4",
			tasks:    [][]int{{1, 3}, {3, 5}, {5, 9}, {9, 12}, {12, 13}, {13, 16}, {16, 17}},
			expected: 1,
		},
		{
			name:     "Case 5",
			tasks:    [][]int{{12, 13}, {13, 15}, {17, 20}, {13, 14}, {19, 21}, {18, 20}},
			expected: 3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := twoheaps.MinimumMachines(test.tasks); got != test.expected {
				t.Errorf("MinimumMachines(%v) = %v, want %v", test.tasks, got, test.expected)
			}
		})
	}
}
