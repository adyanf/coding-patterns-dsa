package kway_merge

import (
	"container/heap"
)

// The K-way merge pattern is an essential algorithmic strategy for merging K sorted data structures, such as arrays and linked lists, into a single sorted data structure.
// This technique is an expansion of the standard merge sort algorithm, which traditionally merges two sorted data structures into one.
// The K-way merge algorithm works by repeatedly selecting the smallest (or largest, if we’re sorting in descending order) element from among
// the first elements of the K input lists and adding this element to a new output list (with the same data type as the inputs).
// This process is repeated until all elements from all input lists have been merged into the output list, maintaining the sorted order.
//
// We can use the K-way merge in two ways:
// - Using a min heap
//   1. Insert the first element of each list into a min heap. This sets up our starting point, with the heap helping us efficiently track the smallest current element among the lists.
//   2. Remove the smallest element from the heap (which is always at the top) and add it to the output list. This ensures that our output list is being assembled in sorted order.
//   3. Keep track of which list each element in the heap came from. This is for knowing where to find the next element to add to the heap.
//   4. After removing the smallest element from the heap and adding it to the output list, replace it with the next element from the same list the removed element belonged to.
//   5. Repeat steps 2–4 until all elements from all input lists have been merged into the output list.
// - Making groups of two and repeatedly merging them
//   1. Start by dividing the K sorted lists into pairs, making groups of two. This organizes our lists into manageable units for merging.
//   2. For each pair of lists, perform a standard two-way merge operation. This is similar to the merge step in merge sort, where two sorted lists are combined into a single sorted list. This step results in k/2 merged lists.
//   3. If there are an odd number of lists in a group at any point, simply leave one list unmerged in that round. This ensures that no list is left out of the merging process.
//   4. Repeat the process of pairing up the resulting lists from the previous merge and merging them again until only one sorted list remains, which is the final result.
//
// Use this pattern when these conditions are fulfilled:
// - Involves merging sorted arrays or a matrix: The problem involves a collection of sorted arrays or a matrix with rows or columns sorted in a specific order that needs to be merged. This could be the core of the problem or a step toward the solution.
// - Seeking the k-th smallest/largest across sorted collections: The problem involves identifying the k-th smallest or largest element across multiple sorted arrays or linked lists.

// FindKSmallestPairs finds the k smallest pairs from given array list1 and list2
func FindKSmallestPairs(list1 []int, list2 []int, k int) [][]int {
	// init the min sum heap to help getting the minimum sum for each iteration
	var minSumHeap MinSumHeap
	heap.Init(&minSumHeap)

	// push the sum of every element of list1 with the first element of list2
	for i := 0; i < len(list1); i++ {
		heap.Push(&minSumHeap, Sum{sum: list1[i] + list2[0], left: i, right: 0})
	}

	// init result and counter
	var result [][]int
	var counter int

	// keep the for loop as long as the heap is not empty and the counter still smaller than k
	for !minSumHeap.Empty() && counter < k {
		// pop the smallest sum from minSumHeap
		smallest := heap.Pop(&minSumHeap).(Sum)
		left, right := smallest.left, smallest.right

		// add the smallest pair to result and add the counter
		result = append(result, []int{list1[left], list2[right]})
		counter++

		// for the popped pair, push the sum of left elementh of list1 and right+1 elementh of list2, if right+1 still less than length list2
		nextRight := right + 1
		if nextRight < len(list2) {
			heap.Push(&minSumHeap, Sum{sum: list1[left] + list2[nextRight], left: left, right: nextRight})
		}
	}

	return result
}

// struct Sum initialization
type Sum struct {
	sum   int
	left  int
	right int
}

// struct MinSumHeap initialization
type MinSumHeap []Sum

// Len returns the length of the heap
func (h MinSumHeap) Len() int {
	return len(h)
}

// Empty returns true if the heap is empty
func (h MinSumHeap) Empty() bool {
	return len(h) == 0
}

// Less returns true if the element with index i should sort before the element with index j
func (h MinSumHeap) Less(i, j int) bool {
	return h[i].sum < h[j].sum || (h[i].sum == h[j].sum && h[i].left < h[j].left)
}

// Swap swaps the elements with indexes i and j
func (h MinSumHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top show the fastest end time
func (h MinSumHeap) Top() interface{} {
	return h[0]
}

// Push pushes an element into the UsageHeap
func (h *MinSumHeap) Push(x interface{}) {
	*h = append(*h, x.(Sum))
}

// Pop pops the element at the top of the UsageHeap
func (h *MinSumHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
