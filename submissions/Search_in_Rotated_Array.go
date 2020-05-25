/*
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
*/
BinSearch(nums []int, target int, low int, high int) int {
    if (low > high) {
        return -1
    }
    mid := low + (high - low) /2
    //fmt.Println(low, mid,high, nums[low],nums[mid],nums[high], target)
    if (nums[mid] < target) {
        return BinSearch(nums, target, mid+1, high)
    }
    if (nums[mid] > target) {
        return BinSearch(nums, target, low, mid-1)
    }

    return mid
}
func FindPivot(nums []int, low int, high int) int {
    mid := low + (high - low) /2
    //fmt.Println(low, mid,high)
    if (low + 1 == high && nums[low] > nums[high]) {
        return high
    }
    if (low >= high) {
        return low
    }

    if (nums[mid] > nums[low]) {
        return FindPivot(nums, mid, high)
    }
    return FindPivot(nums, low, mid)
}
func search(nums []int, target int) int {
    L := len(nums)
    if (L == 0) {
        return -1
    } else if (L==1) {
        if target == nums[0] {
            return 0
        }
        return -1
    }
    P:= FindPivot(nums, 0, len(nums)-1)
    fmt.Println(P)
    X:= BinSearch(nums, target, 0, P-1)
    if (X < 0 && P < L) {
        X = BinSearch(nums, target, P, len(nums)-1)
    }
    return X
}

