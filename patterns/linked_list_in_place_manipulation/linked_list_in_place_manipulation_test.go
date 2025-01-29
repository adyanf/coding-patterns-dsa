package linked_list_in_place_manipulation_test

import (
	"testing"

	"github.com/adyanf/coding-patterns-dsa/patterns/linked_list_in_place_manipulation"
	"github.com/adyanf/coding-patterns-dsa/structs"
)

func TestSwapPairs(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Case 1",
			nums:     []int{1, 2, 3, 4},
			expected: []int{2, 1, 4, 3},
		},
		{
			name:     "Case 2",
			nums:     []int{10, 1, 2, 3, 4, 5},
			expected: []int{1, 10, 3, 2, 5, 4},
		},
		{
			name:     "Case 3",
			nums:     []int{28, 21, 14, 7},
			expected: []int{21, 28, 7, 14},
		},
		{
			name:     "Case 4",
			nums:     []int{11, 12, 13, 14, 15},
			expected: []int{12, 11, 14, 13, 15},
		},
		{
			name:     "Case 5",
			nums:     []int{1, 2},
			expected: []int{2, 1},
		},
	}

	for _, tc := range testCases {
		ll := &structs.LinkedList{}
		ll.CreateLinkedList(tc.nums)

		got := linked_list_in_place_manipulation.SwapPairs(ll.Head)

		llResult := &structs.LinkedList{}
		llResult.InsertNodeAtHead(got)

		llExpected := &structs.LinkedList{}
		llExpected.CreateLinkedList(tc.expected)

		if llResult.String() != llExpected.String() {
			t.Errorf("SwapPairs(%v) = %v, expected %v", tc.nums, llResult.String(), llExpected.String())
		}
	}
}
