package sliding_window_test

import (
	"testing"

	"github.com/adyanf/coding-patterns-dsa/patterns/sliding_window"
)

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
		got := sliding_window.FindLongestSubstring(tc.str)
		if got != tc.expected {
			t.Errorf("FindLongestSubstring(%v) = %v, expected %v", tc.str, got, tc.expected)
		}
	}
}

func TestLongestRepeatingCharacterReplacement(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{
			name:     "Case 1",
			s:        "aaacbbbaabab",
			k:        2,
			expected: 6,
		},
		{
			name:     "Case 2",
			s:        "aaacbbbaabab",
			k:        1,
			expected: 4,
		},
		{
			name:     "Case 3",
			s:        "dippitydip",
			k:        4,
			expected: 6,
		},
		{
			name:     "Case 4",
			s:        "coollooc",
			k:        2,
			expected: 6,
		},
		{
			name:     "Case 5",
			s:        "aaaaaaaaaa",
			k:        2,
			expected: 10,
		},
		{
			name:     "Case 6",
			s:        "mbjhzdidgyzfmegqmabvdqjdlkzhqejjnwwarshmziqokbnalmtqjxzcpofgfjfembxdaubqmfjedchojpveyzlcyhjbyvlflmdizempazgrmsvxjyzrslamvszzukvrudzghrcmohoittwrjjdpyrfpexciuczivimdbgvddyrvhxtkrlpixifovlvgawpslhyiuqypdckfvyincjkliskzsofckfjqitirvmzevxmtgkpkylucrwqqtkltvtzuuyzgpyiudfftuhcpkykrlmhywwwoqfsxkjupbikymlzosythoboyomkebergpmajnwqxuarhssgweaziuyeppubxmnbqjsopfxvlzwaqdjxgledtppepakcqewlniiwkitoemvkxktcwrilnotrtwjiszvhfetnenxcvnczohlllwdeirjkkljjukzrgjnauupwkwijqxzaosryjrcojmxqyfrmokuuyywyotgywbujgugvtdbqdkuxtgoobximfixpgrktbcwdyyznlmibkdfbqbyrfwaegxceedbxoevioclgpmwclnxvnrwlftmfrzkdthrrdudqaiuxrclvukhonhwbxuvfrquvbylkkztyjbwihiztcowvzpcsvhowttljzgwmjynlmxhreepvmmgsofqpbzqmrhebztogfqvncmtrorvujcknvlyueixqwvvpiogecwqmfkqddazcwmyxdpaheupibmmhqhwtzvtxkumzwretgfidzfsttdsafvqfojvdemhaqovaczjwshaivysrmsinndvwstvfbjxcvqiwkaqfvzuxkrkguymuuazxopfotdphzowpngnmrmgvxmdvdycyniaunlviwpuvdhvhnngrfzfiqnjhsmeqemhzbtfaynupqcxggftgzvfwgdetzlxmraeytijttudiywbctrwikcjwcjqnaxmucqanjfffmbbuubhrgqnrsvvfqenbynbpiiptlwram",
			k:        903,
			expected: 952,
		},
	}

	for _, tc := range testCases {
		got := sliding_window.LongestRepeatingCharacterReplacement(tc.s, tc.k)
		if got != tc.expected {
			t.Errorf("FindLongestSubstringWithCharacterReplacement(%v, %v) = %v, expected %v", tc.s, tc.k, got, tc.expected)
		}
	}
}

func TestFindRepeatedSequences(t *testing.T) {
	testCases := []struct {
		name     string
		dna      string
		k        int
		expected map[interface{}]bool
	}{
		{
			name: "Case 1",
			dna:  "AAAAACCCCCAAAAACCCCCC",
			k:    8,
			expected: map[interface{}]bool{
				"AAAAACCC": true,
				"AAAACCCC": true,
				"AAACCCCC": true,
			},
		},
		{
			name: "Case 2",
			dna:  "GGGGGGGGGGGGGGGGGGGGGGGGG",
			k:    9,
			expected: map[interface{}]bool{
				"GGGGGGGGG": true,
			},
		},
		{
			name: "Case 3",
			dna:  "TTTTTCCCCCCCTTTTTTCCCCCCCTTTTTTT",
			k:    10,
			expected: map[interface{}]bool{
				"CCCCCCCTTT": true,
				"CCCCCCTTTT": true,
				"CCCCCTTTTT": true,
				"CCCCTTTTTT": true,
				"TCCCCCCCTT": true,
				"TTCCCCCCCT": true,
				"TTTCCCCCCC": true,
				"TTTTCCCCCC": true,
				"TTTTTCCCCC": true,
			},
		},
		{
			name: "Case 4",
			dna:  "AAAAAACCCCCCCAAAAAAAACCCCCCCTG",
			k:    10,
			expected: map[interface{}]bool{
				"AAAAAACCCC": true,
				"AAAAACCCCC": true,
				"AAAACCCCCC": true,
				"AAACCCCCCC": true,
			},
		},
		{
			name: "Case 5",
			dna:  "ATATATATATATATAT",
			k:    6,
			expected: map[interface{}]bool{
				"ATATAT": true,
				"TATATA": true,
			},
		},
	}

	for _, tc := range testCases {
		got := sliding_window.FindRepeatedSequences(tc.dna, tc.k)
		for key := range tc.expected {
			if !got.Exists(key) {
				t.Errorf("FindRepeatedSequences(%v, %v) = %v, expected %v", tc.dna, tc.k, got, tc.expected)
				break
			}
		}
	}
}

func TestFindRepeatedSequencesWithRabinKarpAlgorithm(t *testing.T) {
	testCases := []struct {
		name     string
		dna      string
		k        int
		expected map[interface{}]bool
	}{
		{
			name: "Case 1",
			dna:  "AAAAACCCCCAAAAACCCCCC",
			k:    8,
			expected: map[interface{}]bool{
				"AAAAACCC": true,
				"AAAACCCC": true,
				"AAACCCCC": true,
			},
		},
		{
			name: "Case 2",
			dna:  "GGGGGGGGGGGGGGGGGGGGGGGGG",
			k:    9,
			expected: map[interface{}]bool{
				"GGGGGGGGG": true,
			},
		},
		{
			name: "Case 3",
			dna:  "TTTTTCCCCCCCTTTTTTCCCCCCCTTTTTTT",
			k:    10,
			expected: map[interface{}]bool{
				"CCCCCCCTTT": true,
				"CCCCCCTTTT": true,
				"CCCCCTTTTT": true,
				"CCCCTTTTTT": true,
				"TCCCCCCCTT": true,
				"TTCCCCCCCT": true,
				"TTTCCCCCCC": true,
				"TTTTCCCCCC": true,
				"TTTTTCCCCC": true,
			},
		},
		{
			name: "Case 4",
			dna:  "AAAAAACCCCCCCAAAAAAAACCCCCCCTG",
			k:    10,
			expected: map[interface{}]bool{
				"AAAAAACCCC": true,
				"AAAAACCCCC": true,
				"AAAACCCCCC": true,
				"AAACCCCCCC": true,
			},
		},
		{
			name: "Case 5",
			dna:  "ATATATATATATATAT",
			k:    6,
			expected: map[interface{}]bool{
				"ATATAT": true,
				"TATATA": true,
			},
		},
	}

	for _, tc := range testCases {
		got := sliding_window.FindRepeatedSequencesWithRabinKarpAlgorithm(tc.dna, tc.k)
		for key := range tc.expected {
			if !got.Exists(key) {
				t.Errorf("FindRepeatedSequencesWithRabinKarpAlgorithm(%v, %v) = %v, expected %v", tc.dna, tc.k, got, tc.expected)
				break
			}
		}
	}
}
