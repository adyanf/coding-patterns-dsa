package modified_binary_search

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
