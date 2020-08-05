/*
Given an array, strs, with strings consisting of only 0s and 1s. Also two integers m and n.

Now your task is to find the maximum number of strings that you can form with given m 0s and n 1s.
Each 0 and 1 can be used at most once.




Example 1:

Input: strs = ["10","0001","111001","1","0"], m = 5, n = 3
Output: 4
Explanation: This are totally 4 strings can be formed by the using of 5 0s and 3 1s, which are "10","0001","1","0".
Example 2:

Input: strs = ["10","0","1"], m = 1, n = 1
Output: 2
Explanation: You could form "10", but then you'd have nothing left. Better form "0" and "1".

3D Knapsack problem
You can use 2D array to simplify memory requirement because we never refer to it beyond i-1 access

say you are m = 1 n = 1
i=0
000
000
000
,now we encounter "01"
000
011
011
,now we encounter ANOTHER "01"
000
011
012
,now we encounter "01", it wont update because we have no more slots left
000
011
012



*/
func CountZero(s string) (r int) {
	for _, x := range s {
		if x == '0' {
			r++
		}
	}
	return r
}

func Make2D(r, c, v int) [][]int {
	ret := make([][]int, r)
	for i := 0; i < len(ret); i++ {
		ret[i] = make([]int, c)
		if v > 0 {
			for j := 0; j < c; j++ {
				ret[i][j] = v
			}
		}
	}
	return ret
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func findMaxForm(strs []string, m int, n int) int {
	DP := make([][][]int, len(strs)+1)
	z, o := 0, 0
	for i := 0; i <= len(strs); i++ {
		if i > 0 {
			z = CountZero(strs[i-1])
			o = len(strs[i-1]) - z
		}
		DP[i] = Make2D(m+1, n+1, 0)
		if i == 0 {
			continue
		} // null string uses none of the slots
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				prev := DP[i-1][j][k]
				if j >= z && k >= o {
					//
					DP[i][j][k] = Max(prev, DP[i-1][j-z][k-o]+1)
				} else {
					DP[i][j][k] = prev
				}
			}
		}
	}
	return DP[len(strs)][m][n]
}
