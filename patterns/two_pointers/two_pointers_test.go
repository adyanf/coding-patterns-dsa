package two_pointers_test

import (
	"testing"

	"github.com/adyanf/coding-patterns-dsa/patterns/two_pointers"
	"github.com/adyanf/coding-patterns-dsa/structs"
	"github.com/stretchr/testify/assert"
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

func TestRemoveNthLastNode(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		delete   int
		expected []int
	}{
		{
			name:     "Case 1",
			nums:     []int{23, 28, 10, 5, 67, 39, 70, 28},
			delete:   2,
			expected: []int{23, 28, 10, 5, 67, 39, 28},
		},
		{
			name:     "Case 2",
			nums:     []int{34, 53, 6, 95, 38, 28, 17, 63, 16, 76},
			delete:   3,
			expected: []int{34, 53, 6, 95, 38, 28, 17, 16, 76},
		},
		{
			name:     "Case 3",
			nums:     []int{288, 224, 275, 390, 4, 383, 330, 60, 193},
			delete:   4,
			expected: []int{288, 224, 275, 390, 4, 330, 60, 193},
		},
		{
			name:     "Case 4",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			delete:   1,
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:     "Case 5",
			nums:     []int{69, 8, 49, 106, 116, 112},
			delete:   6,
			expected: []int{8, 49, 106, 116, 112},
		},
	}

	for _, tc := range testCases {
		ll := &structs.LinkedList{}
		ll.CreateLinkedList(tc.nums)

		got := two_pointers.RemoveNthLastNode(ll.Head, tc.delete)

		llResult := &structs.LinkedList{}
		llResult.InsertNodeAtHead(got)

		llExpected := &structs.LinkedList{}
		llExpected.CreateLinkedList(tc.expected)

		if llResult.String() != llExpected.String() {
			t.Errorf("RemoveNthLastNode(%v, %v) = %v, expected %v", tc.nums, tc.delete, llResult.String(), llExpected.String())
		}
	}
}

func TestSortColors(t *testing.T) {
	testCases := []struct {
		name     string
		colors   []int
		expected []int
	}{
		{
			name:     "Case 1",
			colors:   []int{0, 1, 0},
			expected: []int{0, 0, 1},
		},
		{
			name:     "Case 2",
			colors:   []int{1},
			expected: []int{1},
		},
		{
			name:     "Case 3",
			colors:   []int{2, 2},
			expected: []int{2, 2},
		},
		{
			name:     "Case 4",
			colors:   []int{1, 1, 0, 2},
			expected: []int{0, 1, 1, 2},
		},
		{
			name:     "Case 5",
			colors:   []int{2, 1, 1, 0, 0},
			expected: []int{0, 0, 1, 1, 2},
		},
	}

	for _, tc := range testCases {
		got := two_pointers.SortColors(tc.colors)
		assert.Equal(t, tc.expected, got)
	}
}
