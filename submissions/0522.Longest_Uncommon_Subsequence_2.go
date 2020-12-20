/*
Given a list of strings, you need to find the longest uncommon subsequence among them.
The longest uncommon subsequence is defined as the longest subsequence of one of these strings and this
 subsequence should not be any subsequence of the other strings.

A subsequence is a sequence that can be derived from one sequence by deleting some characters without changing the order
of the remaining elements. Trivially, any string is a subsequence of itself and an empty string is a subsequence of any string.

The input will be a list of strings, and the output needs to be the length of the longest uncommon subsequence.
If the longest uncommon subsequence doesn't exist, return -1.

Example 1:
Input: "aba", "cdc", "eae"
Output: 3
Note:

All the given strings' lengths will not exceed 10.
The length of the given list will be in the range of [2, 50].


Notes:
1. If there is only 1 string with maxlength, that is the length of uncommon SubSequence.
2. Remove duplicate strings

*/import (
	"sort"
)

// checking if sub is a SubSequence of s
func SubSequence(sub, s string) bool {
	s1 := 0
	if len(sub) > len(s) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] == sub[s1] {
			s1++
		}
		if s1 == len(sub) {
			return true
		}
	}
	return s1 == len(s)
}

func findLUSlength(strs []string) int {

	// sort the string on decreasing order of length
	sort.SliceStable(strs, func(i, j int) bool { return len(strs[i]) > len(strs[j]) })
	// eliminate duplicate strings
	M := map[string]int{}
	for _, x := range strs {
		M[x]++
	}
	for i := 0; i < len(strs); i++ {
		issubseq := false
		if M[strs[i]] != 1 {
			continue
		}
		for j := 0; j < i && !issubseq; j++ {
			issubseq = SubSequence(strs[i], strs[j])
		}
		if !issubseq {
			return len(strs[i])
		}
	}
	return -1
}
