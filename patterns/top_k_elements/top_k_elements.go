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
	var frequencyMinHeap FrequencyMinHeap
	heap.Init(&frequencyMinHeap)

	// populate the heap with number frequencies
	for key, value := range frequencies {
		heap.Push(&frequencyMinHeap, Frequency{element: key, count: value})
		// if the heap has length more than k, then pop the smallest frequency (top of the heap)
		if frequencyMinHeap.Len() > k {
			heap.Pop(&frequencyMinHeap)
		}
	}

	// populate the result from the frequency min heap
	result := make([]int, 0, k)
	for !frequencyMinHeap.Empty() {
		popped := heap.Pop(&frequencyMinHeap).(Frequency)
		result = append(result, popped.element)
	}
	return result
}

// FindKthLargest returns the k-th largest number from unsorted nums
func FindKthLargest(nums []int, k int) int {
	// init min heap to store sorted numbers ascending
	var minHeap structs.MinHeap
	heap.Init(&minHeap)

	// populate the heap with k-first elements
	for i := 0; i < k; i++ {
		heap.Push(&minHeap, nums[i])
	}

	// at this point since the heap already has k element
	for i := k; i < len(nums); i++ {
		// every element that got pushed to the heap must have value larger than the top of the heap
		if nums[i] > minHeap.Top() {
			// we must popped the top, before push the current element to maintain heap with size of k
			heap.Pop(&minHeap)
			heap.Push(&minHeap, nums[i])
		}
	}

	// the k-th largest element will be the top of the heap after the for loop
	return minHeap.Top()
}

// ReorganizeString returns a string that has no identical adjacent characters if possible or return empty string
func ReorganizeString(str string) string {
	// calculate the frequency of each element
	frequencies := make(map[rune]int)
	for _, ch := range str {
		frequencies[ch] = frequencies[ch] + 1
	}

	// init frequency max heap
	var frequencyMaxHeap FrequencyMaxHeap
	heap.Init(&frequencyMaxHeap)

	// populate the max heap with number frequencies sorted with the largest frequency in the root
	for key, value := range frequencies {
		heap.Push(&frequencyMaxHeap, Frequency{element: int(key), count: value})
	}

	// init result and the previous character
	result := ""
	previous := Frequency{}

	// keep iterate as long as the max heap is not empty or we still has previous frequency that hasn't been re-push to the max heap
	for !frequencyMaxHeap.Empty() || (previous.count != 0 && rune(previous.element) != 0) {
		// if we have previous frequency but the max heap already empty, then it's not possible to reorganize string without identical adjacent characters
		if (previous.count != 0 && rune(previous.element) != 0) && frequencyMaxHeap.Empty() {
			return ""
		}

		// pop element from the top of max heap, add the element char to result, and reduce the count
		poppedElement := heap.Pop(&frequencyMaxHeap).(Frequency)
		count, char := poppedElement.count, rune(poppedElement.element)
		result += string(char)
		count -= 1

		// if previous element has count > 0, then push it back to max heap, and reset previous frequency
		if previous.count != 0 && rune(previous.element) != 0 {
			heap.Push(&frequencyMaxHeap, previous)
			previous = Frequency{}
		}

		// if current element has count > 0, then set it as previous frequency
		if count != 0 {
			previous.count = count
			previous.element = int(char)
		}
	}

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

// struct FrequencyMinHeap initialization
type FrequencyMaxHeap []Frequency

// Len returns the length of the heap
func (h FrequencyMaxHeap) Len() int {
	return len(h)
}

// Empty returns true if the heap is empty
func (h FrequencyMaxHeap) Empty() bool {
	return len(h) == 0
}

// Less returns true if the element with index i should sort before the element with index j
func (h FrequencyMaxHeap) Less(i, j int) bool {
	return h[i].count > h[j].count || (h[i].count == h[j].count && h[i].element < h[j].element)
}

// Swap swaps the elements with indexes i and j
func (h FrequencyMaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top show the fastest end time
func (h FrequencyMaxHeap) Top() interface{} {
	return h[0]
}

// Push pushes an element into the FrequencyMaxHeap
func (h *FrequencyMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Frequency))
}

// Pop pops the element at the top of the FrequencyMaxHeap
func (h *FrequencyMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
