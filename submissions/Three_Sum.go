/*
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]

*/


func twoSum(nums []int, IDX int, R[][]int) [][]int  {
	l, h := IDX+1, len(nums)-1
	for l < h {
	s := nums[IDX] + nums[l] + nums[h]
	    if s == 0 {
		    R = append(R, []int{nums[IDX], nums[l], nums[h]})
            //get rid of duplicate elements
		    for l < h && nums[l] == nums[l+1] { l++}
		    for l < h && nums[h] == nums[h-1] { h--}
		    l++
		    h--
	    }
	    if s < 0 { l++}
	    if s > 0 { h-- }
	}
    return R
}



func threeSum(nums []int) (r [][]int) {
	sort.Ints(nums)
    L := len(nums)
    if L < 3 {return r}
	for i := 0; i < L-2 ; i++ {
        if nums[i] > 0 { break} // stop at 0
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
        r = twoSum(nums, i, r)
	}
	return r
}
