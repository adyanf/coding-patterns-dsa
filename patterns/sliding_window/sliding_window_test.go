package sliding_window

import "testing"

func TestFindLongestSubstring(t *testing.T) {
	testCases := []struct {
		name     string
		str      string
		expected int
	}{
		{
			name:     "Case 1",
			str:      "abcdbea",
			expected: 5,
		},
		{
			name:     "Case 2",
			str:      "aba",
			expected: 2,
		},
		{
			name:     "Case 3",
			str:      "abccabcabcc",
			expected: 3,
		},
		{
			name:     "Case 4",
			str:      "aaaabaaa",
			expected: 2,
		},
		{
			name:     "Case 5",
			str:      "bbbbb",
			expected: 1,
		},
	}

	for _, tc := range testCases {
		got := FindLongestSubstring(tc.str)
		if got != tc.expected {
			t.Errorf("FindLongestSubstring(%v) = %v, expected %v", tc.str, got, tc.expected)
		}
	}
}
