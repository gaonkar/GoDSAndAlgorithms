/*
Given 2 integers n and start. Your task is return any permutation p of (0,1,2.....,2^n -1) such that :

p[0] = start
p[i] and p[i+1] differ by only one bit in their binary representation.
p[0] and p[2^n -1] must also differ by only one bit in their binary representation.


Example 1:

Input: n = 2, start = 3
Output: [3,2,0,1]
Explanation: The binary representation of the permutation is (11,10,00,01).
All the adjacent element differ by one bit. Another valid permutation is [3,1,0,2]
Example 2:

Input: n = 3, start = 2
Output: [2,6,7,5,4,0,1,3]
Explanation: The binary representation of the permutation is (010,110,111,101,100,000,001,011).


Constraints:

1 <= n <= 16
0 <= start < 2 ^ n
*/
func Power(n int) int {
	if n < 1 {
		return 1
	}
	return 2 * Power(n-1)
}
func grayCode(n int) []int {
	N := Power(n)
	r := make([]int, N)
	if n == 0 {
		return r
	}
	r[1] = 1
	n = 4
	for n <= N {
		for i := n / 2; i < n; i++ {
			r[i] = (n / 2) | r[n-i-1]
		}
		n *= 2
	}
	return r
}

func Reverse(nums []int) {
	n := len(nums) - 1
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[n] = nums[n], nums[i]
		n--
	}
}

// cool reverse
func rotateleft(nums []int, k int) {
	n := len(nums) - k
	Reverse(nums)
	Reverse(nums[:n])
	Reverse(nums[n:])
}

func circularPermutation(n int, start int) []int {
	r := grayCode(n)
	i := 0
	for i < len(r) {
		if start == r[i] {
			break
		}
		i++
	}
	if i == len(r) {
		return r
	}
	rotateleft(r, i)
	return r
}
