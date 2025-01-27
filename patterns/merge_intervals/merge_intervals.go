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
			// completely inside
			if currentEnd <= prevEnd {
			} else {
				// overlapping
				lastAddedInterval[1] = currentEnd
			}
		} else {
			// two intervals not intersect at all
			result = append(result, intervals[i])
		}
	}

	return result
}
