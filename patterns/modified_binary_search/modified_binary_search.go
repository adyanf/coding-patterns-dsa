package modified_binary_search

import "math"

// The modified binary search pattern is an extension of the traditional binary search algorithm and can be applied to a wide range of problems.
// Before we delve into the modified version, let’s first recap the classic binary search algorithm.
// Binary search is an efficient search algorithm for searching a target value in sorted arrays or sorted lists that support direct addressing
// (also known as random access). It follows a divide-and-conquer approach, significantly reducing the search space with each iteration.
// The algorithm uses three indexes—start, end, and middle. With middle = start + (end - start) / 2 or middle = (start + end) / 2.
// - If the target value is equal to the middle index element, we have found the target, and the search terminates.
// - If the target value is less than the middle index element, update the end index to middle−1 and repeat.
// - If the target value is greater than the middle index element, update the start index to middle+1 and repeat.
//
// The modified binary search pattern builds upon the basic binary search algorithm discussed above.
// It involves adapting the traditional binary search approach by applying certain conditions or transformations,
// allowing us to solve problems in which input data are modified in a certain way.
// - Binary search on a modified array: Sometimes, the array may be modified in a certain way, which affects the search process.
// - Binary search with multiple requirements: When searching for a target satisfying multiple requirements, a modified binary search can be used.
//   It involves adapting the comparison logic within the binary search to accommodate multiple specifications.
//
// Use this pattern when these conditions are fulfilled:
// - Target value in sorted data: The problem involves locating a specific target value—or identifying its first or last occurrence—within a sorted array or list.
//   This pattern applies to data structures that support direct addressing.
// - Partially sorted segments: We may use this pattern when segments of an input array are sorted.
//
// Don't use this pattern when these conditions are fulfilled:
// - Lack of direct addressing: The input data structure does not support direct addressing.
// - Unsorted or inappropriately sorted data: The data to search is not sorted according to criteria relevant to the search.
// - Non-value-based solutions: The problem does not require identifying a specific value or range of values.

// SingleNonDuplicate returns the element that appears only once in a sorted array where every other element appears twice.
func SingleNonDuplicate(nums []int) int {
	// set starting point as first and last index of array
	start, end := 0, len(nums)-1

	// continue iterate as long as start < end
	for start < end {
		// calculate the middle of start and end
		middle := start + (end-start)/2

		// if the middle index is odd, then subtract it by one
		// it is to make sure that middle always in the start of the duplicate elements
		if middle%2 == 1 {
			middle -= 1
		}

		// if the nums[middle] == nums[middle+1], it means that all elements up to this point were in pairs,
		// and the non duplicate element must appear after middle. Therefore, move the start pointer toward the end.
		// move the start as middle+2 (next occurrence of duplicate elements, because the middle+1 already equal middle).
		//
		// if the nums[middle] != nums[middle+1], it means that the non duplicate element must have appeared before middle.
		// therefore, move the end pointer toward the start.
		// move the end as middle (because the middle can be the non duplicate element).
		if nums[middle] == nums[middle+1] {
			start = middle + 2
		} else {
			end = middle
		}
	}

	// return the nums[start] because the start will point to non duplicate element
	return nums[start]
}

// FindClosestElement returns k elements closest to target, including target if found.
func FindClosestElement(nums []int, k int, target int) []int {
	// if the length of nums is equal to k return nums, because the target is irrelevant
	if len(nums) == k {
		return nums
	}

	// if the target is less than or equal to the first element of nums
	// then return the first k elements of nums
	if target <= nums[0] {
		return nums[:k]
	}

	// if the target is greater than or equal to the last element of nums
	// then return the last k elements of nums
	if target >= nums[len(nums)-1] {
		return nums[len(nums)-k:]
	}

	// init the start and end pointers based on nums
	start, end := 0, len(nums)-1

	// find the index of the target or the closest element to target
	var closest int
	// iterate as long as the start pointer < end pointer
	for start <= end {
		closest = start + (end-start)/2

		// break if the target index already found
		if nums[closest] == target {
			break
		}

		// otherwise, move the start pointer or end pointer to the closest
		if nums[closest] < target {
			start = closest + 1
		} else {
			end = closest - 1
		}
	}
	// if start > end, then closest is the index where the target should be inserted
	if start > end {
		closest = start
	}

	// set the window left based on closest index
	windowLeft := closest - 1
	// set the window right based on window left + 1
	windowRight := windowLeft + 1
	// iterate as long as the windows size is less than k
	for (windowRight - windowLeft - 1) < k {
		// if window left already less than 0, then increment window right
		if windowLeft < 0 {
			windowRight++
			continue
		}

		// if window right already length of nums then decrement window Left
		// if |nums[windowLeft]-target| <= |nums[windowRight]-target|, then decrement window left
		// otherwise, increment window right
		if windowRight == len(nums) || math.Abs(float64(nums[windowLeft]-target)) <= math.Abs(float64(nums[windowRight]-target)) {
			windowLeft--
		} else {
			windowRight++
		}
	}

	// return the window slice of nums
	return nums[windowLeft+1 : windowRight]
}

// RotatedBinarySearchIterative search a target in array nums, which might be rotated, with iterative style
func RotatedBinarySearchIterative(nums []int, target int) int {
	// init search parameter
	start, end := 0, len(nums)-1
	// keep iterating as long as start index is less than or equal end index
	for start <= end {
		// calculate the mid index based on the search parameter
		mid := start + (end-start)/2
		// if the mid index contain the target return the index immediately
		if nums[mid] == target {
			return mid
		}

		// if the array is sorted from start index to mid index
		// check if the target is less than mid index and larger or equal to start index
		// 	- if yes, then search the first half of the search parameter by end = mid - 1
		// 	- if no, then search the latter half of the search parameter by start = mid + 1
		if nums[start] <= nums[mid] {
			if nums[start] <= target && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
			continue
		}

		// here, array is sorted from mid index to end index
		// check if the target is larger than mid index and less or equal to end index
		// 	- if yes, then search the latter half of the search parameter by start = mid + 1
		// 	- if no, then search the first half of the search parameter by end = mid - 1
		if nums[mid] < target && target <= nums[end] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	// if the for loop break, then we didn't found the target
	return -1
}

// RotatedBinarySearchRecursive search a target in array nums, which might be rotated, with recursive style
func RotatedBinarySearchRecursive(nums []int, target int) int {
	return binarySearchParametered(nums, 0, len(nums)-1, target)
}

// binarySearchParametered search a target in an array but limited by the start and the end indexes
// in this example the array might be rotated
func binarySearchParametered(nums []int, start int, end int, target int) int {
	// base case for recursive function, if start > end then we didn't found the target
	if start > end {
		return -1
	}

	// calculate the mid index based on the search parameter
	mid := start + (end-start)/2

	// if the mid index contain the target return the index immediately
	if nums[mid] == target {
		return mid
	}

	// if the array is sorted from start index to mid index
	// check if the target is less than mid index and larger or equal to start index
	// 	- if yes, then search the first half of the search parameter
	// 	- if no, then search the latter half of the search parameter
	if nums[start] <= nums[mid] {
		if target < nums[mid] && target >= nums[start] {
			return binarySearchParametered(nums, start, mid-1, target)
		}
		return binarySearchParametered(nums, mid+1, end, target)
	}

	// here, the array is rotated (not sorted) in between start index to mid index
	// check if the target is larger than mid index and less or equal to end index
	// 	- if yes, then search the latter half of the search parameter
	// 	- if no, then search the first half of the search parameter
	if nums[mid] < target && target <= nums[end] {
		return binarySearchParametered(nums, mid+1, end, target)
	}
	return binarySearchParametered(nums, start, mid-1, target)
}
