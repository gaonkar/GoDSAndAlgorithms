/*
Write a program to find the n-th ugly number.

Ugly numbers are positive integers which are divisible by a or b or c.



Example 1:

Input: n = 3, a = 2, b = 3, c = 5
Output: 4
Explanation: The ugly numbers are 2, 3, 4, 5, 6, 8, 9, 10... The 3rd is 4.
Example 2:

Input: n = 4, a = 2, b = 3, c = 4
Output: 6
Explanation: The ugly numbers are 2, 3, 4, 6, 8, 9, 10, 12... The 4th is 6.
Example 3:

Input: n = 5, a = 2, b = 11, c = 13
Output: 10
Explanation: The ugly numbers are 2, 4, 6, 8, 10, 11, 12, 13... The 5th is 10.
Example 4:

Input: n = 1000000000, a = 2, b = 217983653, c = 336916467
Output: 1999999984
*/

// TLE

func MinArr(a []int) int {
    ret := a[0]
    for _,x := range(a[1:]) {
        if x < ret {ret = x}
    }
    return ret
}

func IncrMin(a,i []int, v int) {
    for x, _ := range(a) {
        if v == a[x] {a[x]+=i[x]}
    }
}

func nthUglyNumber(n int, a int, b int, c int) int {
    ret := 0
    inc := []int{a,b,c}
    arr := []int{a,b,c}
    for n > 0 {
        ret = MinArr(arr)
        IncrMin(arr, inc, ret)
        n--
    }
    return ret
}

/* Binary search and Math

If some body asks find the Nth number that is multiple of a, the answer is simple --> F(N) = N/a
what if it is for a, b then then answer is F(N) = N/a + N/b - N/lcm(a,b)
so for a,b,c you have a + b + c - (ab+bc+ca) + abc 

*/

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func LCM(a,b int) int {
    return a * b / GCD(a,b)
}

func GenerateF(a,b,c int) []int {
    ret := []int{a,b,c}
    ret = append(ret, -LCM(a,b))
    ret = append(ret, -LCM(b,c))
    ret = append(ret, -LCM(c,a))
    ret = append(ret, LCM(c,-ret[3])) //LCM(a,b,c)
    return ret
}

func NthNum(i int, a[]int) int {
    ret := 0
    for _,x := range(a) {
        ret += i/x
    }
    return ret
}


func nthUglyNumber(n int, a int, b int, c int) int {
    lo,hi := 1, 2000000000
    lcm := GenerateF(a,b,c)
    for lo < hi {
        mid := lo + (hi - lo)/2
        num := NthNum(mid, lcm)
        if num < n {
            lo = mid + 1
            continue
        }
        hi = mid
    }
    return lo
}
