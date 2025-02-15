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

// KthSmallestElement returns the k-th smallest element from a matrix
func KthSmallestElement(matrix [][]int, k int) int {
	// init the cell min heap to help getting the minimum cell for each iteration
	var cellMinHeap CellMinHeap
	heap.Init(&cellMinHeap)

	// push the value of matrix for each row with column 0
	for i := 0; i < len(matrix); i++ {
		heap.Push(&cellMinHeap, Cell{value: matrix[i][0], row: i, column: 0})
	}

	// init the smallest element and counter
	smallestElement := 0
	counter := 0
	for {
		// pop the smallest cell from the heap
		smallestCell := heap.Pop(&cellMinHeap).(Cell)
		rowIndex, columnIndex := smallestCell.row, smallestCell.column

		// add counter to mark the popped element position after sorted
		counter++

		// for the popped cell, push the value of matrix with row rowIndex and column nextColumn [columnIndex + 1]
		nextColumn := columnIndex + 1
		if nextColumn < len(matrix[0]) {
			heap.Push(&cellMinHeap, Cell{value: matrix[rowIndex][nextColumn], row: rowIndex, column: nextColumn})
		}

		// if counter equal k then set the smallest element and break the for loop
		if counter == k {
			smallestElement = smallestCell.value
			break
		}
	}

	return smallestElement
}

// KSmallestNumber returns the k-th smallest number from lists
func KSmallestNumber(lists [][]int, k int) int {
	// init the list min heap to help getting the minimum list element for each iteration
	var listMinHeap ListMinHeap
	heap.Init(&listMinHeap)

	// push the value of the first index of each list together with its list index and element index
	for i := 0; i < len(lists); i++ {
		if len(lists[i]) > 0 {
			heap.Push(&listMinHeap, ListElement{listIndex: i, elementIndex: 0, value: lists[i][0]})
		}
	}

	// init the smallest element and counter
	var smallestNumber, counter int

	// iterating as long as the list min heap not empty and counter less than k
	for !listMinHeap.Empty() && counter < k {
		// pop the smallest list element
		smallestListElement := heap.Pop(&listMinHeap).(ListElement)
		listIndex, elementIndex := smallestListElement.listIndex, smallestListElement.elementIndex

		// set the smallest number and add counter
		smallestNumber = smallestListElement.value
		counter += 1

		// if the smallest list element's list still has next element, push it to the list min heap
		if elementIndex+1 < len(lists[listIndex]) {
			heap.Push(&listMinHeap, ListElement{listIndex: listIndex, elementIndex: elementIndex + 1, value: lists[listIndex][elementIndex+1]})
		}
	}

	// when either list is empty or counter already equal to k then return the smallest number
	return smallestNumber
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

// Push pushes an element into the MinSumHeap
func (h *MinSumHeap) Push(x interface{}) {
	*h = append(*h, x.(Sum))
}

// Pop pops the element at the top of the MinSumHeap
func (h *MinSumHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// struct Cell initialization
type Cell struct {
	value  int
	row    int
	column int
}

// struct CellMinHeap initialization
type CellMinHeap []Cell

// Len returns the length of the heap
func (h CellMinHeap) Len() int {
	return len(h)
}

// Empty returns true if the heap is empty
func (h CellMinHeap) Empty() bool {
	return len(h) == 0
}

// Less returns true if the element with index i should sort before the element with index j
func (h CellMinHeap) Less(i, j int) bool {
	return h[i].value < h[j].value
}

// Swap swaps the elements with indexes i and j
func (h CellMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top show the fastest end time
func (h CellMinHeap) Top() interface{} {
	return h[0]
}

// Push pushes an element into the CellMinHeap
func (h *CellMinHeap) Push(x interface{}) {
	*h = append(*h, x.(Cell))
}

// Pop pops the element at the top of the CellMinHeap
func (h *CellMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// struct ListElement initialization
type ListElement struct {
	value        int
	listIndex    int
	elementIndex int
}

// struct ListMinHeap initialization
type ListMinHeap []ListElement

// Len returns the length of the heap
func (h ListMinHeap) Len() int {
	return len(h)
}

// Empty returns true if the heap is empty
func (h ListMinHeap) Empty() bool {
	return len(h) == 0
}

// Less returns true if the element with index i should sort before the element with index j
func (h ListMinHeap) Less(i, j int) bool {
	return h[i].value < h[j].value || (h[i].value == h[j].value && h[i].listIndex < h[j].listIndex)
}

// Swap swaps the elements with indexes i and j
func (h ListMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top show the fastest end time
func (h ListMinHeap) Top() interface{} {
	return h[0]
}

// Push pushes an element into the ListMinHeap
func (h *ListMinHeap) Push(x interface{}) {
	*h = append(*h, x.(ListElement))
}

// Pop pops the element at the top of the ListMinHeap
func (h *ListMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
