/*
Given an array of integers, find out whether there are two distinct indices i and j in the array such that the absolute difference between nums[i] and nums[j] is at most t and the absolute difference between i and j is at most k.

Example 1:

Input: nums = [1,2,3,1], k = 3, t = 0
Output: true
Example 2:

Input: nums = [1,0,1,1], k = 1, t = 2
Output: true
Example 3:

Input: nums = [1,5,9,1,5,9], k = 2, t = 3
Output: false
*/

// Method 1 range query
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	l := 0
	M := make(map[int][]int)
	//    if t < 0 {return false}
	for i := 0; i < len(nums); i++ {
		for _, x := range M {
			if x[0] <= nums[i] && nums[i] <= x[1] {
				return true
			}
		}
		M[nums[i]] = []int{nums[i] - t, nums[i] + t}
		if i-l >= k {
			delete(M, nums[l])
			l++
		}
	}
	return false
}

// Method 2 Buckets
func BucketID(n, t int) int {
	return n / (t + 1)
}
func Found(M map[int]int, bid, val, t int) bool {
	v, ok := M[bid]
	if !ok {
		return false
	}
	v -= val
	if v < 0 {
		v = -v
	}
	return v <= t
}
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	l := 0
	M := make(map[int]int)
	if t < 0 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		bid := BucketID(nums[i], t)
		if Found(M, bid, nums[i], t) {
			return true
		}
		if Found(M, bid-1, nums[i], t) {
			return true
		}
		if Found(M, bid+1, nums[i], t) {
			return true
		}
		M[bid] = nums[i]
		if i-l >= k {
			bid = BucketID(nums[l], t)
			delete(M, bid)
			l++
		}

	}
	return false
}