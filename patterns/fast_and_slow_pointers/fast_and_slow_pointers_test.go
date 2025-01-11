package fast_and_slow_pointers_test

import (
	"testing"

	"github.com/adyanf/coding-patterns-dsa/patterns/fast_and_slow_pointers"
	"github.com/adyanf/coding-patterns-dsa/structs"
)

func TestFindDuplicate(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name: "Case 1",
			nums: []int{3, 4, 4, 4, 2},
			// 3 3
			// 4 2
			// 2 2
			// ----
			// 3 2
			// 4 4
			// return 4
			expected: 4,
		},
		{
			name: "Case 2",
			nums: []int{1, 1},
			// 1 1
			// 1 1
			// -----
			// 1 1
			// return 1
			expected: 1,
		},
		{
			name: "Case 3",
			nums: []int{1, 3, 4, 2, 2},
			// 1 1
			// 3 2
			// 2 2
			// ----
			// 1 2
			// 3 4
			// 2 2
			// return 2
			expected: 2,
		},
		{
			name: "Case 4",
			nums: []int{1, 3, 6, 2, 7, 3, 5, 4},
			// 1 1
			// 3 2
			// 2 5
			// 6 2
			// 5 5
			// ------
			// 1 5
			// 3 3
			// return 3
			expected: 3,
		},
		{
			name: "Case 5",
			nums: []int{1, 2, 2},
			// 1 1
			// 2 2
			// ----
			// 1 2
			// 2 2
			// return 2
			expected: 2,
		},
	}

	for _, tc := range testCases {
		got := fast_and_slow_pointers.FindDuplicate(tc.nums)
		if got != tc.expected {
			t.Errorf("FindDuplicate(%v) = %v, expected %v", tc.nums, got, tc.expected)
		}
	}
}

func TestCircularArrayLoop(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Case 1",
			nums:     []int{1, 3, -2, -4, 1},
			expected: true,
		},
		{
			name:     "Case 2",
			nums:     []int{2, 1, -1, -2},
			expected: false,
		},
		{
			name:     "Case 3",
			nums:     []int{5, 4, -2, -1, 3},
			expected: false,
		},
		{
			name:     "Case 4",
			nums:     []int{1, 2, -3, 3, 4, 7, 1},
			expected: true,
		},
		{
			name:     "Case 5",
			nums:     []int{3, 3, 1, -1, 2},
			expected: true,
		},
	}

	for _, tc := range testCases {
		got := fast_and_slow_pointers.CircularArrayLoop(tc.nums)
		if got != tc.expected {
			t.Errorf("CircularArrayLoop(%v) = %v, expected %v", tc.nums, got, tc.expected)
		}
	}
}

func TestPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Case 1",
			nums:     []int{1, 2, 3, 2, 1},
			expected: true,
		},
		{
			name:     "Case 2",
			nums:     []int{4, 7, 9, 5, 4},
			expected: false,
		},
		{
			name:     "Case 3",
			nums:     []int{2, 3, 5, 5, 3, 2},
			expected: true,
		},
		{
			name:     "Case 4",
			nums:     []int{6, 1, 0, 5, 1, 6},
			expected: false,
		},
		{
			name:     "Case 5",
			nums:     []int{3, 6, 9, 8, 4, 8, 9, 6, 3},
			expected: true,
		},
	}

	for _, tc := range testCases {
		ll := &structs.LinkedList{}
		ll.CreateLinkedList(tc.nums)

		got := fast_and_slow_pointers.Palindrome(ll.Head)
		if got != tc.expected {
			t.Errorf("Palindrome(%v) = %v, expected %v", tc.nums, got, tc.expected)
		}
	}
}
