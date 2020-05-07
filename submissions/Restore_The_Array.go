/*
A program was supposed to print an array of integers. The program forgot to print whitespaces and the array is printed
as a string of digits and all we know is that all integers in the array were in the range [1, k] and there are no
leading zeros in the array.

Given the string s and the integer k. There can be multiple ways to restore the array.

Return the number of possible array that can be printed as a string s using the mentioned program.

The number of ways could be very large so return it modulo 10^9 + 7



Example 1:

Input: s = "1000", k = 10000
Output: 1
Explanation: The only possible array is [1000]
Example 2:

Input: s = "1000", k = 10
Output: 0
Explanation: There cannot be an array that was printed this way and has all integer >= 1 and <= 10.
Example 3:

Input: s = "1317", k = 2000
Output: 8
Explanation: Possible arrays are [1317],[131,7],[13,17],[1,317],[13,1,7],[1,31,7],[1,3,17],[1,3,1,7]
Example 4:

Input: s = "2020", k = 30
Output: 1
Explanation: The only possible array is [20,20]. [2020] is invalid because 2020 > 30. [2,020] is ivalid because 020 contains leading zeros.
Example 5:

Input: s = "1234567890", k = 90
Output: 34

*/


var m map[int] int
var M int

func ValidPrefix(s string, k int) []int {
    var l []int
    i := 0
    sum := 0
    if (len(s) == 0 || s[0] == '0') {return l}
    for i < len(s) {
        x := int(s[i] - '0')
        sum = sum * 10 + x
        if (sum > k) {
            break
        }
        if x > 0  || i < len(s)-1 && s[i+1] != 0 || i == len(s) -1 {
            l = append(l, i+1)
        }
        i = i + 1
    }
    return l
}
var C int

func NOA(s string, k int) int {
    //M := 10* 10 * 10 + 7
    if (len(s) == 0) { return 1}
    if (len(s) == 1 && s[0] != '0') {return 1}
    v,ok := m[len(s)]
    if (ok) { return v}
    l := ValidPrefix(s, k)
    n := 0
    for _,x:= range(l) {
        //fmt.Println(s[:x], x, n, s[x:], len(s), len(l))
        n = n + NOA(s[x:], k)
        C = (C + 1) //% M

    }
    //fmt.Println(s, n)
    m[len(s)] = n % M
    return n % M
}

func numberOfArrays(s string, k int) int {
    M = 1000 * 1000 * 1000 + 7

    m = make(map[int]int)
    return NOA(s,k)
}

func numberOfArrays(s string, k int) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	m := len(strconv.Itoa(k))
	MOD := 1000000007
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if i >= j && s[i-j] != '0' {
				thenum, _ := strconv.Atoi(s[i-j : i])
				if thenum <= k {
					dp[i] += dp[i-j]
					dp[i] %= MOD
				}
			}
		}
	}
	return dp[n]
}
