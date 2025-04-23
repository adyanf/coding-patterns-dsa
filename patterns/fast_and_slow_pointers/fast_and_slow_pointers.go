package fast_and_slow_pointers

import "github.com/adyanf/coding-patterns-dsa/structs"

// Fast and slow pointers pattern uses two pointers to traverse an iterable data structure,
// but the pointers move at different speeds, often to identify cycles or find a specific target.
// The speeds of the pointers can be adjusted according to the problem statement.
// Use this pattern when these conditions are fulfilled:
// - Linear data structure: The input data can be traversed in a linear fashion, such as an array, linked list, or string.
// In addition, if either of these conditions is fulfilled:
// - Cycle or intersection detection: The problem involves finding cycles or intersections in the data structure.
// - Find the starting element at the second quantile: The problem involves finding the starting element of the second quantile,
//   i.e., second half, second quartile, etc. For example, the problem asks to find the middle element of an array or a linked list.

// FindDuplicate finds the single duplicate number in an array where integers range from 1 to n and length n+1.
// Uses a fast and slow pointers pattern to find the duplicate number.
// This solution has time complexity of O(n) and space complexity of O(1).
func FindDuplicate(nums []int) int {
	// detects the cycle with fast and slow pointers
	// use the array element as the index of the next element
	// break when the slow pointer equals the fast pointer (reach intersection point)
	slow, fast := nums[0], nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	// move the slow pointer to the starting point and keep the fast pointer at the intersection point
	// move both one steps at a time
	// break when the slow pointer equals the fast pointer (reach entry point)
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	// the entry point should be the duplicated number
	return fast
}

// CircularArrayLoop checks if there exists a cycle in the given integer array following specific conditions.
// A cycle exists if we can continuously move forward or backward in the array, following indices, without changing in direction.
// It returns true if the cycle exists, otherwise false.
// Forward or backward direction is determined by the sign of the array elements.
// Uses fast and slow pointer technique to detect cycles.
// This solution has time complexity of O(n^2) and space complexity of O(1).
func CircularArrayLoop(nums []int) bool {
	arraySize := len(nums)
	for i := 0; i < arraySize; i++ {
		slow, fast := i, i
		direction := nums[i] > 0

		for {
			// check for every step if there is a changing in direction or cycling in itself
			// if yes, then break the current iteration
			slow = nextIndex(slow, nums[slow], arraySize)
			if isNotCycle(nums, direction, slow) {
				break
			}
			fast = nextIndex(fast, nums[fast], arraySize)
			if isNotCycle(nums, direction, fast) {
				break
			}
			fast = nextIndex(fast, nums[fast], arraySize)
			if isNotCycle(nums, direction, fast) {
				break
			}

			// if slow and fast pointers are equal, then the circular array is cyclic
			if slow == fast {
				return true
			}
		}
	}

	return false
}

// Palindrome checks whether a given linked list is a palindrome.
// Uses a fast and slow pointers pattern to find the middle element and reverse the second part of the linked list.
// This solution has time complexity of O(n) and space complexity of O(1).
func Palindrome(head *structs.LinkedListNode) bool {
	// find the middle element of the linked list
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse the second half of the linked list
	reverseHalfList := structs.ReverseLinkedList(slow)
	// move the slow pointer to head
	// move the fast pointer to the reversed second-half linked list
	slow = head
	fast = reverseHalfList

	// start the palindrome check
	for fast != nil {
		// if the element is not equal, then the linked list is not palindrome
		if slow.Data != fast.Data {
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}
	// reverse back the second half of the linked list to return the original linked list
	structs.ReverseLinkedList(reverseHalfList)

	return true
}

// IsHappy checks if a number is a "happy number" by iteratively summing the squares of its digits until it equals 1 or loops.
// Uses a fast and slow pointers pattern to check if the squared sum of its digits is cyclic or not (equals 1).
// This solution has time complexity of O(log n) and space complexity of O(1).
func IsHappy(num int) bool {
	// iterate through the squared sum of its digit until the fast pointer equals 1 or its slow pointer equals its fast pointer
	slow, fast := num, calculateSumOfSquaredDigits(num)
	for fast != 1 && slow != fast {
		slow = calculateSumOfSquaredDigits(slow)
		fast = calculateSumOfSquaredDigits(calculateSumOfSquaredDigits(fast))
	}

	// if its fast pointer equals 1, it means that the number is a happy number
	if fast == 1 {
		return true
	}
	// if cyclic (its slow pointer equals its fast pointer), then it is not a happy number
	return false
}

func nextIndex(currentIndex int, indexValue int, arraySize int) int {
	nextIdx := currentIndex + indexValue
	if nextIdx < 0 {
		nextIdx = nextIdx + arraySize
	}
	if nextIdx >= arraySize {
		nextIdx = nextIdx - arraySize
	}
	return nextIdx
}

func isNotCycle(nums []int, prevDirection bool, index int) bool {
	nextDirection := nums[index] > 0

	// check if the next direction is the same as the previous direction
	// and check if the cycle is of an element pointing to itself
	// if so, it is not a cycle
	if (prevDirection != nextDirection) || (absInt(nums[index])%len(nums) == 0) {
		return true
	} else {
		return false
	}
}

func calculateSumOfSquaredDigits(num int) int {
	sum := 0
	for num > 0 {
		digit := num % 10
		num = num / 10
		sum += digit * digit
	}

	return sum
}

func absInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
