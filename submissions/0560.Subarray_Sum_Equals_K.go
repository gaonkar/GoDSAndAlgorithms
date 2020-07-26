/*
Given an array of integers nums and an integer k, return the total number of continuous subarrays whose sum equals to k.



Example 1:

Input: nums = [1,1,1], k = 2
Output: 2
Example 2:

Input: nums = [1,2,3], k = 3
Output: 2


Constraints:

1 <= nums.length <= 2 * 104
-1000 <= nums[i] <= 1000
-107 <= k <= 107
*/


func subarraySum(nums []int, k int) int {
    c:= 0
    s:= 0
    H:=make(map[int] int)
    H[0] = 1 //sum of position upto 0
    for _,n:= range(nums) {
        s +=n
        c +=H[s-k] // Sum(0..i) - Sum(0..(i-k))
        H[s]++
    }
    return c
}
