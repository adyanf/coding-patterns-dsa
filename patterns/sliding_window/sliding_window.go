package sliding_window

import (
	"math"

	"github.com/adyanf/coding-patterns-dsa/structs"
)

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

func LongestRepeatingCharacterReplacement(s string, k int) int {
	stringLength := len(s)
	start := 0
	lengthOfMaxSubstring := 0
	charFreq := make(map[byte]int)
	mostFreqChar := 0

	for end := 0; end < stringLength; end++ {
		// increment the frequency of the iterated character
		charFreq[s[end]]++

		// check the most frequent character
		mostFreqChar = maxInt(mostFreqChar, charFreq[s[end]])

		// if the window required more than k replacement to be valid, then move the start window forward
		// consequently, reduce the frequency of the character at the start of the window
		if end-start+1-mostFreqChar > k {
			charFreq[s[start]]--
			start++
		}

		// update the length of the longest substring
		lengthOfMaxSubstring = maxInt(lengthOfMaxSubstring, end-start+1)
	}

	return lengthOfMaxSubstring
}

func FindRepeatedSequences(dna string, k int) *structs.Set {
	output := structs.NewSet()
	sequenceFreq := make(map[string]int)
	start := 0

	for start <= len(dna)-k {
		dnaSequence := dna[start : start+k]
		sequenceFreq[dnaSequence]++
		if sequenceFreq[dnaSequence] > 1 && !output.Exists(dnaSequence) {
			output.Add(dnaSequence)
		}

		start++
	}
	return output
}

func FindRepeatedSequencesWithRabinKarpAlgorithm(dna string, k int) *structs.Set {
	stringLength := len(dna)

	nucleotidesMapping := map[rune]int{'A': 1, 'C': 2, 'G': 3, 'T': 4}
	baseValue := 4

	numbers := make([]int, stringLength)
	for i, ch := range dna {
		numbers[i] = nucleotidesMapping[ch]
	}

	hashValue := 0

	hashSet := structs.NewSet()
	output := structs.NewSet()

	for i := 0; i <= stringLength-k; i++ {
		if i == 0 {
			for j := 0; j < k; j++ {
				hashValue += numbers[j] * int(math.Pow(float64(baseValue), float64(k-j-1)))
			}
		} else {
			hashValue = (hashValue-(numbers[i-1]*int(math.Pow(float64(baseValue), float64(k-1)))))*baseValue + numbers[i+k-1]
		}

		if hashSet.Exists(hashValue) {
			output.Add(dna[i : i+k])
		}

		hashSet.Add(hashValue)
	}

	return output
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
