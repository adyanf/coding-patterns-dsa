package merge_intervals

import (
	"sort"
)

// The merge intervals pattern deals with problems involving overlapping intervals.
// Each interval is represented by a start and an end time.
// This pattern involves tasks such as merging intersecting intervals, inserting new intervals
// into existing sets, or determining the minimum number of intervals needed to cover a given range.
// The most common problems solved using this pattern are event scheduling, resource allocation, and time slot consolidation.
// The key to understanding this pattern and exploiting its power lies in understanding how any two intervals may overlap.
// Use this pattern when these conditions are fulfilled:
// - Array of intervals: The input data is an array of intervals.
// - Overlapping intervals: The problem requires dealing with overlapping intervals, either to find their union,
// 	 their intersection, or the gaps between them.

// LeastTime calculate the least amount of time unit needed to process tasks
// with cooling down period n between identical task.
func LeastTime(tasks []byte, n int) int {
	// first find the frequency of each task
	taskFreqs := make(map[byte]int)
	for _, task := range tasks {
		taskFreqs[task]++
	}

	// create a sorted task frequency slice with ascending sort
	type taskFreq struct {
		task byte
		freq int
	}
	sortedTaskFreq := make([]taskFreq, 0, len(taskFreqs))
	for task, freq := range taskFreqs {
		sortedTaskFreq = append(sortedTaskFreq, taskFreq{task: task, freq: freq})
	}
	sort.Slice(sortedTaskFreq, func(i, j int) bool {
		return sortedTaskFreq[i].freq < sortedTaskFreq[j].freq
	})

	// calculate the maxFreq based on the frequency of last index
	maxFreq := sortedTaskFreq[len(sortedTaskFreq)-1].freq
	sortedTaskFreq = sortedTaskFreq[:len(sortedTaskFreq)-1]

	// get the initial idle time based on cooling down period multiplied by max freq minus one
	// because the first one don't need cooling period
	idleTime := (maxFreq - 1) * n

	// calculate the idle time used by the less frequent tasks
	for len(sortedTaskFreq) > 0 && idleTime > 0 {
		idleTime -= min(maxFreq-1, sortedTaskFreq[len(sortedTaskFreq)-1].freq)
		sortedTaskFreq = sortedTaskFreq[:len(sortedTaskFreq)-1]
	}
	// idle time shouldn't be negative
	if idleTime < 0 {
		idleTime = 0
	}

	// the processing time is the number of tasks (one task need one time unit for processing) + idle time
	return len(tasks) + idleTime
}

// MergeIntervals merge the overlapping intervals so the result will consist of non-overlapping intervals.
func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	var result [][]int
	result = append(result, intervals[0])

	for i := 1; i < len(intervals); i++ {
		lastAddedInterval := result[len(result)-1]
		currentStart := intervals[i][0]
		currentEnd := intervals[i][1]
		prevEnd := lastAddedInterval[1]

		// two intervals intersect
		if currentStart <= prevEnd {
			// if overlapping - update the end time of the last interval
			if currentEnd > prevEnd {
				lastAddedInterval[1] = currentEnd
			} else {
				// completely inside, do nothing
			}
		} else {
			// two intervals not intersect at all
			result = append(result, intervals[i])
		}
	}

	return result
}

// InsertInterval inserts a new interval into an existing intervals
// The result will be sorted and non-overlapping
func InsertInterval(existingIntervals [][]int, newInterval []int) [][]int {
	newStart, newEnd := newInterval[0], newInterval[1]
	i, n := 0, len(existingIntervals)
	var result [][]int

	// insert existing intervals which have start time less than or equal to the new interval start time
	for i < n && existingIntervals[i][0] <= newStart {
		result = append(result, existingIntervals[i])
		i++
	}

	// insert the new interval
	// if the result is still empty or the new interval start time is greater than the last interval end time (not overlap), then append the new interval
	// otherwise, if the new interval end time is greater than the last interval end time (not completely inside), then update the last interval end time
	if len(result) == 0 || newStart > result[len(result)-1][1] {
		result = append(result, []int{newStart, newEnd})
	} else {
		if newEnd > result[len(result)-1][1] {
			result[len(result)-1][1] = newEnd
		}
	}

	for i < n {
		lastIndex := len(result) - 1
		// two intervals intersect
		if existingIntervals[i][0] <= result[lastIndex][1] {
			// if two intervals overlapping - update the end time of the last interval
			if existingIntervals[i][1] > result[lastIndex][1] {
				result[lastIndex][1] = existingIntervals[i][1]
			}
		} else {
			// if the new interval is not inserted, then append it to the end of the result
			result = append(result, existingIntervals[i])
		}

		i++
	}

	return result
}
