/*
Given a positive integer n, find the smallest integer which has exactly the same digits existing in the integer n and is
greater in value than n. If no such positive integer exists, return -1.
Note that the returned integer should fit in 32-bit integer, if there is a valid answer but it does not fit in 32-bit
integer, return -1.

Example 1:

Input: n = 12
Output: 21
Example 2:

Input: n = 21
Output: -1

 Constraints:

1 <= n <= 231 - 1
*/
// from stack exchange
const (
	MinUint uint32 = 0 // binary: all zeroes

	// Perform a bitwise NOT to change every bit from 0 to 1
	MaxUint = ^MinUint // binary: all ones

	// Shift the binary number to the right (i.e. divide by two)
	// to change the high bit to 0
	MaxInt = int(MaxUint >> 1) // binary: all ones except high bit

	// Perform another bitwise NOT to change the high bit to 1 and
	// all other bits to 0
	MinInt = ^MaxInt // binary: all zeroes except high bit
)


func Reverse(r []int, i,n int) {
    //fmt.Println(r,i,n)
    for i < n {
        r[i], r[n] = r[n], r[i]
        n--
        i++
    }
}
//https://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order
func nextPermutation(nums []int)  {
    i:=0
    L:=len(nums)
    k:=-1
    // find the largest index k such that a[k] < a[k+1]
    for i < L-1  {
        if nums[i] < nums[i+1] {k=i}
        i++
    }
    if k == -1{
        Reverse(nums,0, L-1)
        return
    }
    // find the number that is greater than nums[k] from k+2 to N
    l := L-1
    for l > k+1 && nums[l] <= nums[k]{
        l--
    }
    //fmt.Println(nums, k,nums[k], l, nums[l])
    // swap them
    nums[k], nums[l] = nums[l], nums[k]
    Reverse(nums,k+1,L-1)
}
func digits(x int) []int {
    ret := make([]int, 0)
    for x > 0 {
        ret = append(ret, x%10)
        x = x / 10
    }
    Reverse(ret,0,len(ret)-1)
    return ret
}

func integer(x []int) int {
    ret := 0
    for _,d := range(x) {
        ret = ret * 10 + d
    }
    return ret
}
func nextGreaterElement(n int) int {
    next := digits(n)
    nextPermutation(next)
    ret := integer(next)
    if ret <= n || ret >= MaxInt {return -1}
    return ret
}
