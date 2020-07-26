/*
Given a sequence of n integers a1, a2, ..., an, a 132 pattern is a subsequence ai, aj, ak
such that i < j < k and ai < ak < aj. Design an algorithm that takes a list of n numbers as input and checks whether
there is a 132 pattern in the list.

Note: n will be less than 15,000.

Example 1:
Input: [1, 2, 3, 4]

Output: False

Explanation: There is no 132 pattern in the sequence.
Example 2:
Input: [3, 1, 4, 2]

Output: True

Explanation: There is a 132 pattern in the sequence: [1, 4, 2].
Example 3:
Input: [-1, 3, 2, 0]

Output: True

Explanation: There are three 132 patterns in the sequence: [-1, 3, 2], [-1, 3, 0] and [-1, 2, 0].
*/

// from stack exchange
const (
	MinUint uint = 0 // binary: all zeroes

	// Perform a bitwise NOT to change every bit from 0 to 1
	MaxUint = ^MinUint // binary: all ones

	// Shift the binary number to the right (i.e. divide by two)
	// to change the high bit to 0
	MaxInt = int(MaxUint >> 1) // binary: all ones except high bit

	// Perform another bitwise NOT to change the high bit to 1 and
	// all other bits to 0
	MinInt = ^MaxInt // binary: all zeroes except high bit
)

func Reverse(nums []int) {
	n := len(nums) - 1
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[n] = nums[n], nums[i]
		n--
	}
}

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	ai := MinInt
	//lets find 231 pattern on the reverse list
	Reverse(nums)
	stk := make([]int, 0)
	for _, x := range nums {
		//fmt.Println(ai, stk, x)
		if x < ai {
			return true
		}
		L := len(stk)
		for L > 0 && x > stk[L-1] {
			L--
			ai = stk[L]
			stk = stk[:L]
		}

		stk = append(stk, x)
	}
	return false
}
