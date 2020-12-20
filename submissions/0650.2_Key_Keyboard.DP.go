/*
Initially on a notepad only one character 'A' is present. You can perform two operations on this notepad for each step:

Copy All: You can copy all the characters present on the notepad (partial copy is not allowed).
Paste: You can paste the characters which are copied last time.


Given a number n. You have to get exactly n 'A' on the notepad by performing the minimum number of steps permitted. Output the minimum number of steps to get n 'A'.

Example 1:

Input: 3
Output: 3
Explanation:
Intitally, we have one character 'A'.
In step 1, we use Copy All operation.
In step 2, we use Paste operation to get 'AA'.
In step 3, we use Paste operation to get 'AAA'.

memoization works better as it only has to deal with factors of N
*/

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func minSteps(n int) int {
	M := make(map[int]int)
	var Helper func(n int) int
	Helper = func(n int) int {
		if n < 2 {
			return 0
		}
		if n < 4 {
			return n
		}
		val, ok := M[n]
		if ok {
			return val
		}
		ret := Helper(1) + n
		for i := 2; i <= n/2; i++ {
			if n%i == 0 {
				ret = Min(ret, Helper(n/i)+i)
			}
		}
		M[n] = ret
		return ret
	}
	return Helper(n)
}
