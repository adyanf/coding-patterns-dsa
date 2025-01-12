package sliding_window

// Sliding window pattern is used to process sequential data, arrays, and strings, for example, to efficiently solve
// subarray or substring problems. It involves maintaining a dynamic window that slides through the array or string,
// adjusting its boundaries as needed to track relevant elements or characters. The window is used to slide over the data
// in chunks corresponding to the window size, and this can be set according to the problem’str requirements.
// It may be viewed as a variation of the two pointers pattern, with the pointers being used to set the window bounds.
// Use this pattern when these conditions are fulfilled:
// - Contiguous data: The input data is stored in a contiguous manner, such as an array or string.
// - Processing subsets of elements: The problem requires repeated computations on a contiguous subset of data elements
//   (a subarray or a substring), such that the window moves across the input array from one end to the other.
//   The size of the window may be fixed or variable, depending on the requirements of the problem.
// - Efficient computation time complexity: The computations performed every time the window moves take constant or very small time.

func FindLongestSubstring(str string) int {
	lastSeenAt := make(map[byte]int)
	start, maxLen, currLen, i := 0, 0, 0, 0

	for i = 0; i < len(str); i++ {
		lastIndex, ok := lastSeenAt[str[i]]
		if ok && lastIndex >= start {
			currLen = i - start
			if currLen > maxLen {
				maxLen = currLen
			}
			start = lastIndex + 1
		}
		lastSeenAt[str[i]] = i
	}

	currLen = i - start
	if currLen > maxLen {
		maxLen = currLen
	}
	return maxLen
}
