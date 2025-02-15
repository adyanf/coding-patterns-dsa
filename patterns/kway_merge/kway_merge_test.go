package kway_merge_test

import (
	"testing"

	"github.com/adyanf/coding-patterns-dsa/patterns/kway_merge"
)

func TestFindKSmallestPairs(t *testing.T) {
	tests := []struct {
		name  string
		list1 []int
		list2 []int
		k     int
		want  [][]int
	}{
		{
			name:  "Case 1",
			list1: []int{1, 2, 300},
			list2: []int{1, 11, 20, 35, 300},
			k:     30,
			want:  [][]int{{1, 1}, {2, 1}, {1, 11}, {2, 11}, {1, 20}, {2, 20}, {1, 35}, {2, 35}, {1, 300}, {300, 1}, {2, 300}, {300, 11}, {300, 20}, {300, 35}, {300, 300}},
		},
		{
			name:  "Case 2",
			list1: []int{1, 1, 2},
			list2: []int{1, 2, 3},
			k:     1,
			want:  [][]int{{1, 1}},
		},
		{
			name:  "Case 3",
			list1: []int{4, 6},
			list2: []int{2, 3},
			k:     2,
			want:  [][]int{{4, 2}, {4, 3}},
		},
		{
			name:  "Case 4",
			list1: []int{4, 7, 9},
			list2: []int{4, 7, 9},
			k:     5,
			want:  [][]int{{4, 4}, {4, 7}, {7, 4}, {4, 9}, {9, 4}},
		},
		{
			name:  "Case 5",
			list1: []int{1, 1, 2},
			list2: []int{1},
			k:     4,
			want:  [][]int{{1, 1}, {1, 1}, {2, 1}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := kway_merge.FindKSmallestPairs(test.list1, test.list2, test.k)

			diff := false
			if len(got) != len(test.want) {
				diff = true
			} else {
				for i := 0; i < len(got); i++ {
					if got[i][0] != test.want[i][0] || got[i][1] != test.want[i][1] {
						diff = true
						break
					}
				}
			}

			if diff {
				t.Errorf("FindKSmallestPairs(%v, %v, %d) = %v, want %v", test.list1, test.list2, test.k, got, test.want)
			}
		})
	}
}
