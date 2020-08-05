/*
Given a string, find the length of the longest substring without repeating characters.
Status

Similar Questions
* [Longest Substring with At Most Two Distinct Characters (Hard)](https://leetcode.com/problems/longest-substring-with-at-most-two-distinct-characters/)
* [Longest Substring with At Most K Distinct Characters (Hard)](https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/)
* [Subarrays with K Different Integers (Hard)](https://leetcode.com/problems/subarrays-with-k-different-integers/)

*/

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func lengthOfLongestSubstring(s string) int {
	ln := 1
	best := 1
	if s == "" {
		return 0
	}
	for i := range s {

		for j := i - 1; j > i-ln; j-- {
			//fmt.Println(i, j, s[i], s[j], ln, best)
			if s[j] == s[i] {
				ln = i - j
			}
		}
		best = Max(ln, best)
		ln++
	}
	return best
}
