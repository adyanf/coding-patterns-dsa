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
	for i := 0; i < len(nums); i++ {
		slow, fast := i, i

		for {
			slow = correctIndex(nums, slow+nums[slow])
			fast = correctIndex(nums, fast+nums[fast]+nums[correctIndex(nums, fast+nums[fast])])
			if slow == fast {
				break
			}
		}

		count, direction, diffDirection := 0, 1, false
		if nums[slow] < 0 {
			direction = -1
		}

		slow = correctIndex(nums, slow+nums[slow])
		if nums[slow]*direction < 0 {
			continue
		}
		count++

		for slow != fast {
			slow = correctIndex(nums, slow+nums[slow])
			if nums[slow]*direction < 0 {
				diffDirection = true
				break
			}
			count++
		}

		if diffDirection || count < 2 {
			continue
		}

		return true
	}

	return false
}

func correctIndex(nums []int, index int) int {
	trueIndex := index
	for trueIndex < 0 {
		trueIndex = trueIndex + len(nums)
	}
	for trueIndex >= len(nums) {
		trueIndex = trueIndex - len(nums)
	}
	return trueIndex
}
