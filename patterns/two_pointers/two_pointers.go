package two_pointers

import (
	"sort"

	"github.com/adyanf/coding-patterns-dsa/structs"
)

// Two pointers is a versatile technique used in problem-solving to efficiently traverse or manipulate
// sequential data structures, such as arrays or linked lists.
// It involves maintaining two pointers that move through the data in a coordinated manner.
// Usually, these pointers start from different positions in the data structure or move in different directions / steps.
// These pointers dynamically adjust based on specific conditions or criteria,
// allowing for the efficient exploration of the data and enabling solutions with optimal time and space complexity.
// Use this pattern when these conditions are fulfilled:
// 1. Linear data structure: The input data can be traversed in a linear fashion, such as an array, linked list, or string.
// 2. Process pairs: Process data elements at two different positions simultaneously.
// 3. Dynamic pointer movement: Both pointers move independently of each other according to certain conditions or criteria.
//    In addition, both pointers might move along the same or two different data structures.

func FindSumOfThree(nums []int, target int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		start := i + 1
		end := len(nums) - 1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if sum == target {
				return true
			} else if sum < target {
				start++
			} else if sum > target {
				end--
			}
		}
	}
	return false
}

func RemoveNthLastNode(head *structs.LinkedListNode, n int) *structs.LinkedListNode {
	left := head
	right := head

	for i := 0; i < n; i++ {
		right = right.Next
	}

	if right == nil {
		return head.Next
	}

	for right.Next != nil {
		right = right.Next
		left = left.Next
	}

	left.Next = left.Next.Next

	return head
}

func SortColors(colors []int) []int {
	left := 0
	right := len(colors) - 1
	i := 0

	for i <= right {
		if colors[i] == 0 {
			colors[i], colors[left] = colors[left], colors[i]
			left++
			i++
		} else if colors[i] == 1 {
			i++
		} else if colors[i] == 2 {
			colors[i], colors[right] = colors[right], colors[i]
			right--
		}
	}

	return colors
}
