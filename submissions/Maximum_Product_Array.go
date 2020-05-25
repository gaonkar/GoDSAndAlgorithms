/*
152. Maximum Product Subarray
Medium

3706

149

Add to List

Share
Given an integer array nums, find the contiguous subarray within an array (containing at least one number) which has the largest product.

Example 1:

Input: [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
Example 2:

Input: [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

*/
func maxProduct(nums []int) int {
    prod := 1

    max := nums[0]
    for i:= range(nums) {
        prod *= nums[i]
        if prod > max {
            max = prod
        }
        if prod == 0 {prod = 1}
    }
    prod = 1
    for i:= len(nums)-1; i >=0; i-- {
        prod *= nums[i]
        if prod > max {
            max = prod
        }
        if prod == 0 {prod = 1}
    }
    return max
}:
