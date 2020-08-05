/*
In a given array nums of positive integers, find three non-overlapping subarrays with maximum sum.

Each subarray will be of size k, and we want to maximize the sum of all 3*k entries.

Return the result as a list of indices representing the starting position of each interval (0-indexed). If there are multiple answers, return the lexicographically smallest one.

Example:

Input: [1,2,1,2,6,7,5,1], 2
Output: [0, 3, 5]
Explanation: Subarrays [1, 2], [2, 6], [7, 5] correspond to the starting indices [0, 3, 5].
We could have also taken [2, 1], but an answer of [1, 3, 5] would be lexicographically larger.


Note:

nums.length will be between 1 and 20000.
nums[i] will be between 1 and 65535.
k will be between 1 and floor(nums.length / 3).

*/import "fmt"

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

func Print2D(r [][]int) {
	for _, x := range r {
		fmt.Println(x)
	}
}

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	N := len(nums)
	O := Make2D(3, N-k+1, 0)
	KSum := make([]int, N-k+1)
	ret := make([]int, 3)
	sum := 0
	for j := 0; j < k; j++ {
		sum = sum + nums[j]
	}
	// find sum of i to i+k and store it at KSum[i]

	for i := 0; i <= N-k; i++ {
		KSum[i] = sum
		if i == N-k {
			break
		}
		sum += (nums[i+k] - nums[i])
	}
	maxval := 0
	// we have [3 3 3 8 13 12 6 ]
	// suppose we had to only choose 1 subarray,
	for j := 0; j < N-k+1; j++ {
		if maxval < KSum[j] {
			maxval = KSum[j]
		}
		O[0][j] = maxval
	}
	for i, b := 1, k; i < 3; i, b = i+1, b+k {
		maxval = 0
		for j := b; j < N-k+1; j++ {
			// say we need to pick 2nd subarray, we start from position k, we add its sum to previous best sum. if that
			// is better than current sum, make it current
			bsum := O[i-1][j-k] + KSum[j]
			if maxval < bsum {
				maxval = bsum
			}
			O[i][j] = maxval
		}
	}
	idx := len(O[0]) - 1
	// (0)  1   2  (3)   4   (5)   6
	// Table:
	//  0   0   0   0    0    0    0
	// (3)  3   3   8    13   13   13
	//  0   0   6  (11)  16   20   20
	//  0   0   0   0    19  (23)  23
	for i := 2; i > -1; i-- {
		for idx > 0 && O[i][idx] == O[i][idx-1] {
			idx--
		}
		ret[i] = idx
		idx = idx - k
	}
	//fmt.Println(KSum)
	//Print2D(O)
	return ret
}
