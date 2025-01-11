package fast_and_slow_pointers

// Fast and slow pointers pattern uses two pointers to traverse an iterable data structure,
// but the pointers move at different speeds, often to identify cycles or find a specific target.
// The speeds of the pointers can be adjusted according to the problem statement.
// Use this pattern when these conditions are fulfilled:
// - Linear data structure: The input data can be traversed in a linear fashion, such as an array, linked list, or string.
// In addition, if either of these conditions is fulfilled:
// - Cycle or intersection detection: The problem involves finding cycles or intersections in the data structure.
// - Find the starting element at the second quantile: The problem involves finding the starting element of the second quantile,
//   i.e., second half, second quartile, etc. For example, the problem asks to find the middle element of an array or a linked list.

func FindDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return fast
}

func CircularArrayLoop(nums []int) bool {
	arraySize := len(nums)
	for i := 0; i < arraySize; i++ {
		slow, fast := i, i
		direction := nums[i] > 0

		for {
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
			if slow == fast {
				return true
			}
		}
	}

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

func absInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
