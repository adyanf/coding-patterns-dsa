package structs

// struct MinHeap initialization
type MinHeap []int

// Len returns the length of the heap
func (h MinHeap) Len() int {
	return len(h)
}

// Empty returns true if the heap is empty
func (h MinHeap) Empty() bool {
	return len(h) == 0
}

// Less returns true if the element with index i should sort before the element with index j
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap swaps the elements with indexes i and j
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top returns the top element of the MinHeap
func (h MinHeap) Top() int {
	return h[0]
}

// Push pushes an element into the MinHeap
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop pops the element at the top of the MinHeap
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// struct MaxHeap initialization
type MaxHeap []int

// Len returns the length of the heap
func (h MaxHeap) Len() int {
	return len(h)
}

// Empty returns true if the heap is empty
func (h MaxHeap) Empty() bool {
	return len(h) == 0
}

// Less returns true if the element with index i should sort before the element with index j
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

// Swap swaps the elements with indexes i and j
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top returns the top element of the MaxHeap
func (h MaxHeap) Top() int {
	return h[0]
}

// Push pushes an element into the MaxHeap
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop pops the element at the top of the MaxHeap
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
