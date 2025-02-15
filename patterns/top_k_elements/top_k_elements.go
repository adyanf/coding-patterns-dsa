package top_k_elements

import (
	"container/heap"

	"github.com/adyanf/coding-patterns-dsa/structs"
)

// The top k elements pattern is an important technique in coding that helps us efficiently find a specific number of elements, known as k, from a set of data.
// This is particularly useful when we’re tasked with identifying the largest, smallest, or most/least frequent elements within an unsorted collection.
// Which data structure can we use to solve such problems? A heap is the best data structure to keep track of the smallest or largest k elements.
// With this pattern, we either use a max heap or a min heap to find the smallest or largest k elements, respectively, because they allow us to efficiently
// maintain a collection of elements ordered in a way that gives us quick access to the smallest (min heap) or largest (max heap) element.
// Use this pattern when these conditions are fulfilled:
// - Unsorted list analysis: We need to extract a specific subset of elements based on their size (largest or smallest), frequency (most or least frequent),
//   or other similar criteria from an unsorted list. This may be the requirement of the final solution, or it may be necessary as an intermediate step toward the final solution.
// - Identifying a specific subset: The goal is to identify a subset rather than just a single extreme value.
//   When phrases like top k, kth largest/smallest, k most frequent, k closest, or k highest/lowest describe our task,
//   it suggests the top k elements pattern is ideal for efficiently identifying a specific subset.
// Don't use this pattern if any of these conditions is fulfilled:
// - Presorted input: The input data is already sorted according to the criteria relevant to solving the problem.
// - Single extreme value: If only 1 extreme value (either the maximum or minimum) is required, that is, k=1,
//   as that problem can be solved in O(n) with a simple linear scan through the input data.

// TopKFrequent return top k element with the largest frequency
func TopKFrequent(nums []int, k int) []int {
	// calculate the frequency of each element
	frequencies := make(map[int]int)
	for _, num := range nums {
		frequencies[num] = frequencies[num] + 1
	}

	// init frequency min heap
	var topFrequencyHeap FrequencyMinHeap
	heap.Init(&topFrequencyHeap)

	// populate the heap with number frequencies
	for key, value := range frequencies {
		heap.Push(&topFrequencyHeap, Frequency{element: key, count: value})
		// if the heap has length more than k, then pop the smallest frequency (top of the heap)
		if topFrequencyHeap.Len() > k {
			heap.Pop(&topFrequencyHeap)
		}
	}

	// populate the result from the top frequency heap
	result := make([]int, 0, k)
	for !topFrequencyHeap.Empty() {
		popped := heap.Pop(&topFrequencyHeap).(Frequency)
		result = append(result, popped.element)
	}
	return result
}

// FindKthLargest returns the k-th largest number from unsorted nums
func FindKthLargest(nums []int, k int) int {
	// init min heap to store sorted numbers ascending
	var minHeap structs.MinHeap
	heap.Init(&minHeap)

	// populate the heap
	for _, num := range nums {
		heap.Push(&minHeap, num)
		// if heap has length more than k, then pop the smallest number (top of heap)
		if minHeap.Len() > k {
			heap.Pop(&minHeap)
		}
	}

	// the k-th largest element will be the top of the heap after the for loop
	result := heap.Pop(&minHeap).(int)
	return result
}

// struct Frequency initialization
type Frequency struct {
	element int
	count   int
}

// struct FrequencyMinHeap initialization
type FrequencyMinHeap []Frequency

// Len returns the length of the heap
func (h FrequencyMinHeap) Len() int {
	return len(h)
}

// Empty returns true if the heap is empty
func (h FrequencyMinHeap) Empty() bool {
	return len(h) == 0
}

// Less returns true if the element with index i should sort before the element with index j
func (h FrequencyMinHeap) Less(i, j int) bool {
	return h[i].count < h[j].count || (h[i].count == h[j].count && h[i].element < h[j].element)
}

// Swap swaps the elements with indexes i and j
func (h FrequencyMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top show the fastest end time
func (h FrequencyMinHeap) Top() interface{} {
	return h[0]
}

// Push pushes an element into the FrequencyMinHeap
func (h *FrequencyMinHeap) Push(x interface{}) {
	*h = append(*h, x.(Frequency))
}

// Pop pops the element at the top of the FrequencyMinHeap
func (h *FrequencyMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
