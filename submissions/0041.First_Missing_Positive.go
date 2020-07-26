/*
Given an unsorted integer array nums, find the smallest missing positive integer.

Follow up: Could you implement an algorithm that runs in O(n) time and uses constant extra space.?



Example 1:

Input: nums = [1,2,0]
Output: 3
Example 2:

Input: nums = [3,4,-1,1]
Output: 2
Example 3:

Input: nums = [7,8,9,11,12]
Output: 1


Constraints:

0 <= nums.length <= 300
-231 <= nums[i] <= 231 - 1
*/
// key insight is that number should be 1.. len(nums)-1

func Partition(v int, nums []int) int {
	L := len(nums)
    for L > 0 && nums[L-1] <= v {L--}
    if L == 0 {return 0}
	j := 0
	for i := 0; i < L; i++ {
		if nums[i] > v {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return j
}

func Abs(x int) int {
    if x < 0 {return -x}
    return x
}


func FMP(nums []int) int {
    for _,v := range(nums) {
        v = Abs(v) - 1
        if v < len(nums) && nums[v] > 0  {
            nums[v] = -nums[v]
        }
    }
    i := 1
    for _,v := range(nums) {
        if v > 0 {
            return i
        }
        i++
    }
    return i
}

func firstMissingPositive(nums []int) int {
    i := Partition(0, nums)

    return FMP(nums[:i])
}
