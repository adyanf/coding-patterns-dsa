package twoheaps

import (
	"container/heap"
	"sort"

	"github.com/adyanf/coding-patterns-dsa/structs"
)

// Heaps are a special data structure that helps you efficiently manage priorities.
// A heap is a specific data structure with a fixed ordering (min or max),
// while a priority queue is an abstract data type that handles custom priority requirements for elements.
// Heaps are typically implemented using arrays to efficiently access the parent and child nodes. The major operations performed on heaps are:
// - Add: This inserts a new element into the heap, which takes O(log n) time.
// - Delete: This removes the root element and rebalances the heap, taking O(log n) time.
// - Peek: This retrieves the smallest or largest element in O(1).
// Use this pattern when these conditions are fulfilled:
// - Linear data: If the input data is linear, it can be sorted or unsorted.
//   - A heap efficiently finds the maximum or minimum elements if the data is unsorted.
//     Operations like insertion and deletion take O(log n) time, ensuring fast access to the top elements.
//   - If the data is sorted, a heap can still be useful when frequent insertions and deletions are required,
//     as it allows for efficient updates and retrieval of the highest or lowest elements, with both insertion and deletion operations also taking O(logn) time.
// - Stream of data: The input data continuously arrives in real time, often in an unpredictable order, requiring efficient handling and processing as it flows in.
//   Heaps automatically enforce priority ordering (e.g., largest weight, smallest cost, highest frequency). This saves you from manually resorting to or scanning each time your data changes.
// - Calculation of maxima and minima: The input data can be categorized into two parts, and we need to repeatedly calculate two maxima, two minima, or one maximum and one minimum from each set.
//   O(logn) insertion/removal and O(1) retrieval.
// - Custom priority-based selection: The problem involves selecting the next element based on specific priority at each step, such as processing the largest task or earliest event.

// MostBooked returns the most booked room in the given meetings.
func MostBooked(meetings [][]int, rooms int) int {
	// Initialize count array with 0s, it represents the number of meetings in each room
	count := make([]int, rooms)

	// Create a min heaps for available rooms and used rooms
	available := &structs.MinHeap{}
	heap.Init(available)

	usedRooms := &UsageHeap{}
	heap.Init(usedRooms)

	// Initialize the available rooms heap
	for i := 0; i < rooms; i++ {
		heap.Push(available, i)
	}

	// Sort the meetings by their start time
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	for i := 0; i < len(meetings); i++ {
		startTime := meetings[i][0]
		endTime := meetings[i][1]

		// free up the rooms that have finished their meetings by the start time
		for !usedRooms.Empty() && (*usedRooms)[0].endTime <= startTime {
			room := heap.Pop(usedRooms).(Usage).id
			heap.Push(available, room)
		}

		// if no rooms are available, delay the meeting until a room becomes free
		if available.Empty() {
			pair := heap.Pop(usedRooms).(Usage)
			endTime = pair.endTime + (endTime - startTime)
			heap.Push(available, pair.id)
		}

		// Allocate the meeting into the available room with the lowest number
		room := heap.Pop(available).(int)
		heap.Push(usedRooms, Usage{endTime: endTime, id: room})
		count[room]++
	}

	// Find room which holds the most meetings
	maxMettingsRoom := 0
	for i := range count {
		if count[i] > count[maxMettingsRoom] {
			maxMettingsRoom = i
		}
	}
	return maxMettingsRoom
}

// MinimumMachines returns the minimum machines needed to execute all tasks without delay.
func MinimumMachines(tasks [][]int) int {
	// Sort tasks at hand based on their start time
	sort.Slice(tasks, func(i int, j int) bool {
		return tasks[i][0] < tasks[j][0]
	})

	// Create machine usage heap
	machines := &UsageHeap{}
	heap.Init(machines)

	for _, task := range tasks {
		startTime, endTime := task[0], task[1]

		// if machine usage heap is not empty and the top of machine end time is less than or equal new task start time
		// then pop the machine usage, because the machine will be reused by new task
		if !machines.Empty() && machines.Top().(Usage).endTime <= startTime {
			heap.Pop(machines)
		}

		// push machine usage with end time of new task
		heap.Push(machines, Usage{endTime: endTime})
	}

	return machines.Len()
}

// MedianOfStream is a data structure to search a median from a stream of numbers
type MedianOfStream struct {
	// We are using two heaps to solve this problem
	maximumList structs.MaxHeap
	minimumList structs.MinHeap
}

// Init will initializes underlying data to handle median of stream of numbers
func (this *MedianOfStream) Init() {
	// Init each heap by using container/heap library
	heap.Init(&this.maximumList)
	heap.Init(&this.minimumList)
}

// InsertNum inserts new number to the underlying data store
func (this *MedianOfStream) InsertNum(num int) float64 {
	// If maximum list is empty, or if the number is smaller or equal to the top of maximum list
	// Then push the number to maximum list, otherwise push the number to minimum list
	if this.maximumList.Empty() || this.maximumList.Top() >= num {
		heap.Push(&this.maximumList, num)
	} else {
		heap.Push(&this.minimumList, num)
	}

	// Rebalancing the heaps, so that the number of elements in each heap has max diff of 1
	if this.maximumList.Len() > this.minimumList.Len()+1 {
		heap.Push(&this.minimumList, heap.Pop(&this.maximumList).(int))
	} else if this.maximumList.Len() < this.minimumList.Len() {
		heap.Push(&this.maximumList, heap.Pop(&this.minimumList).(int))
	}

	return this.FindMedian()
}

// FindMedian finds the median of stream of numbers
func (this *MedianOfStream) FindMedian() float64 {
	// If the number of elements is equal (even), calculate the top of each heap and divide it by 2
	if this.maximumList.Len() == this.minimumList.Len() {
		return (float64(this.maximumList.Top()) + float64(this.minimumList.Top())) / 2.0
	}
	// Otherwise return the top of maximum list
	return float64(this.maximumList.Top())
}

// struct Usage initialization
type Usage struct {
	id      int
	endTime int
}

// struct UsageHeap initialization
type UsageHeap []Usage

// Len returns the length of the heap
func (h UsageHeap) Len() int {
	return len(h)
}

// Empty returns true if the heap is empty
func (h UsageHeap) Empty() bool {
	return len(h) == 0
}

// Less returns true if the element with index i should sort before the element with index j
func (h UsageHeap) Less(i, j int) bool {
	return h[i].endTime < h[j].endTime || (h[i].endTime == h[j].endTime && h[i].id < h[j].id)
}

// Swap swaps the elements with indexes i and j
func (h UsageHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top show the fastest end time
func (h UsageHeap) Top() interface{} {
	return h[0]
}

// Push pushes an element into the UsageHeap
func (h *UsageHeap) Push(x interface{}) {
	*h = append(*h, x.(Usage))
}

// Pop pops the element at the top of the UsageHeap
func (h *UsageHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
