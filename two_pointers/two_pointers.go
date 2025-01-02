package two_pointers

import "sort"

// Two pointers is a versatile technique used in problem-solving to efficiently traverse or manipulate
// sequential data structures, such as arrays or linked lists.
// It involves maintaining two pointers that move through the data in a coordinated manner.
// Usually, these pointers start from different positions in the data structure or move in different directions / steps.
// These pointers dynamically adjust based on specific conditions or criteria,
// allowing for the efficient exploration of the data and enabling solutions with optimal time and space complexity.
// Use this pattern when this conditions are fulfilled:
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
