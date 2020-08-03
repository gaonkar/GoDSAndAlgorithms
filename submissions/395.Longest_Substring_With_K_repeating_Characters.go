/*
Find the length of the longest substring T of a given string (consists of lowercase letters only)
such that every character in T appears no less than k times.

Example 1:

Input:
s = "aaabb", k = 3

Output:
3

The longest substring is "aaa", as 'a' is repeated 3 times.
Example 2:

Input:
s = "ababbc", k = 2

Output:
5

The longest substring is "ababb", as 'a' is repeated 2 times and 'b' is repeated 3 times.

consider string  "x.......y". For this letter x and y in string will matter if the inner component is valid.

*/import "fmt"

func idx(x byte) int {
	return int(x - 'a')
}

func longestSubstring(s string, k int) int {
	C := make([]int, 26)
	for i := 0; i < len(s); i++ {
		C[idx(s[i])]++
	}
	fmt.Println(C)
	i, j := 0, len(s)-1
	l, r := i, j

	for i <= j {
		if C[idx(s[i])] < k {
			for l <= i {
				C[idx(s[l])]--
				l++
			}
			i++
			j = r
			continue
		}
		if C[idx(s[j])] < k {
			for r >= j {
				C[idx(s[r])]--
				r--
			}
			j--
			i = l
			continue
		}
		i++
		j--
	}

	return r - l + 1
}
