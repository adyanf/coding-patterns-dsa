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

func TestReverseBetween(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		left     int
		right    int
		expected []int
	}{
		{
			name:     "Case 1",
			nums:     []int{1, 2, 3, 4, 5, 4, 3, 2, 1},
			left:     1,
			right:    9,
			expected: []int{1, 2, 3, 4, 5, 4, 3, 2, 1},
		},
		{
			name:     "Case 2",
			nums:     []int{1, 2, 3, 4, 5},
			left:     2,
			right:    4,
			expected: []int{1, 4, 3, 2, 5},
		},
		{
			name:     "Case 3",
			nums:     []int{103, 7, 10, -9, 105, 67, 31, 63},
			left:     1,
			right:    8,
			expected: []int{63, 31, 67, 105, -9, 10, 7, 103},
		},
		{
			name:     "Case 4",
			nums:     []int{-499, 399, -299, 199, -99, 9},
			left:     3,
			right:    5,
			expected: []int{-499, 399, -99, 199, -299, 9},
		},
		{
			name:     "Case 5",
			nums:     []int{7, -9, 2, -10, 1, -8, 6},
			left:     2,
			right:    5,
			expected: []int{7, 1, -10, 2, -9, -8, 6},
		},
	}

	for _, tc := range testCases {
		ll := &structs.LinkedList{}
		ll.CreateLinkedList(tc.nums)

		got := linked_list_in_place_manipulation.ReverseBetween(ll.Head, tc.left, tc.right)

		llResult := &structs.LinkedList{}
		llResult.InsertNodeAtHead(got)

		llExpected := &structs.LinkedList{}
		llExpected.CreateLinkedList(tc.expected)

		if llResult.String() != llExpected.String() {
			t.Errorf("ReverseBetween(%v, %d, %d) = %v, expected %v", tc.nums, tc.left, tc.right, llResult.String(), llExpected.String())
		}
	}
}

func TestReverseBetweenV2(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		left     int
		right    int
		expected []int
	}{
		{
			name:     "Case 1",
			nums:     []int{1, 2, 3, 4, 5, 4, 3, 2, 1},
			left:     1,
			right:    9,
			expected: []int{1, 2, 3, 4, 5, 4, 3, 2, 1},
		},
		{
			name:     "Case 2",
			nums:     []int{1, 2, 3, 4, 5},
			left:     2,
			right:    4,
			expected: []int{1, 4, 3, 2, 5},
		},
		{
			name:     "Case 3",
			nums:     []int{103, 7, 10, -9, 105, 67, 31, 63},
			left:     1,
			right:    8,
			expected: []int{63, 31, 67, 105, -9, 10, 7, 103},
		},
		{
			name:     "Case 4",
			nums:     []int{-499, 399, -299, 199, -99, 9},
			left:     3,
			right:    5,
			expected: []int{-499, 399, -99, 199, -299, 9},
		},
		{
			name:     "Case 5",
			nums:     []int{7, -9, 2, -10, 1, -8, 6},
			left:     2,
			right:    5,
			expected: []int{7, 1, -10, 2, -9, -8, 6},
		},
	}

	for _, tc := range testCases {
		ll := &structs.LinkedList{}
		ll.CreateLinkedList(tc.nums)

		got := linked_list_in_place_manipulation.ReverseBetweenV2(ll.Head, tc.left, tc.right)

		llResult := &structs.LinkedList{}
		llResult.InsertNodeAtHead(got)

		llExpected := &structs.LinkedList{}
		llExpected.CreateLinkedList(tc.expected)

		if llResult.String() != llExpected.String() {
			t.Errorf("ReverseBetween(%v, %d, %d) = %v, expected %v", tc.nums, tc.left, tc.right, llResult.String(), llExpected.String())
		}
	}
}
