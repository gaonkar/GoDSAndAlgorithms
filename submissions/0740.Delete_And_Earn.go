/*
Given an array nums of integers, you can perform operations on the array.

In each operation, you pick any nums[i] and delete it to earn nums[i] points. After, you must delete every element equal to nums[i] - 1 or nums[i] + 1.

You start with 0 points. Return the maximum number of points you can earn by applying such operations.

Example 1:

Input: nums = [3, 4, 2]
Output: 6
Explanation:
Delete 4 to earn 4 points, consequently 3 is also deleted.
Then, delete 2 to earn 2 points. 6 total points are earned.


Example 2:

Input: nums = [2, 2, 3, 3, 3, 4]
Output: 9
Explanation:
Delete 3 to earn 3 points, deleting both 2's and the 4.
Then, delete 3 again to earn 3 points, and 3 again to earn 3 points.
9 total points are earned.
*/import "sort"

func Max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func deleteAndEarn(nums []int) int {
	sort.Ints(nums) // n log n
	L := len(nums)
	if L == 0 {
		return 0
	}
	if L == 1 {
		return nums[0]
	}
	C := make([]int, L)
	copy(C, nums)
	for i := 0; i < L; i++ {
		for j := 0; j < i; j++ {
			if nums[j] != nums[i]-1 {
				C[i] = Max(C[i], C[j]+nums[i])
			}
		}
	}
	M := 0
	for i := 0; i < L; i++ {
		M = Max(M, C[i])
	}
	return M
}

/*
 In the group discussion, there is another solution where one converts this into
 the robbing house problem
 Since the value of nums 0 < 10k, allocate a 10k array
 now for each C[nums[i]] += nums[i]
 now run the house robbing algorithm on this new array
 */
