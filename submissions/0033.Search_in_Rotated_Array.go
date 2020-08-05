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
*/import "fmt"

func BinSearch(nums []int, target int, low int, high int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)/2
	//fmt.Println(low, mid,high, nums[low],nums[mid],nums[high], target)
	if nums[mid] < target {
		return BinSearch(nums, target, mid+1, high)
	}
	if nums[mid] > target {
		return BinSearch(nums, target, low, mid-1)
	}

	return mid
}
func FindPivot(nums []int, low int, high int) int {
	mid := low + (high-low)/2
	fmt.Println(low, mid, high)

	if low >= high {
		return low
	}
	if nums[mid] > nums[mid+1] {
		return mid + 1
	}
	if nums[mid] > nums[low] {
		return FindPivot(nums, mid, high)
	}
	return FindPivot(nums, low, mid)
}
func search(nums []int, target int) int {
	L := len(nums)
	if L == 0 {
		return -1
	} else if L == 1 {
		if target == nums[0] {
			return 0
		}
		return -1
	}
	P := FindPivot(nums, 0, len(nums)-1)
	fmt.Println(P)
	X := BinSearch(nums, target, 0, P-1)
	if X < 0 && P < L {
		X = BinSearch(nums, target, P, len(nums)-1)
	}
	return X
}

// when duplicates are allowed
func FindPivot(nums []int, lo int, hi int) int {

	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else if nums[mid] < nums[hi] {
			hi = mid
		} else {
			//fmt.Println("*")
			if nums[hi-1] > nums[hi] {
				lo = hi
				break
			}
			hi--
		}
	}
	return lo
}
func search(nums []int, target int) bool {
	L := len(nums)
	if L == 0 {
		return false
	}
	pos = FindPivot(nums, 0, L-1)

	fmt.Println(pos)
	if BinSearch(nums, target, 0, pos-1) {
		return true
	}
	if nums[pos] == target {
		return true
	}
	return BinSearch(nums, target, pos+1, L-1)
}

/*
A peak element is an element that is greater than its neighbors.

Given an input array nums, where nums[i] ≠ nums[i+1], find a peak element and return its index.

The array may contain multiple peaks, in that case return the index to any one of the peaks is fine.

You may imagine that nums[-1] = nums[n] = -∞.

Example 1:

Input: nums = [1,2,3,1]
Output: 2
Explanation: 3 is a peak element and your function should return the index number 2.
Example 2:

Input: nums = [1,2,1,3,5,6,4]
Output: 1 or 5
Explanation: Your function can return either index number 1 where the peak element is 2,
             or index number 5 where the peak element is 6.
Follow up: Your solution should be in logarithmic complexity.
*/
func findPeakElement(nums []int) int {
    lo, hi := 0, len(nums) - 1
    for lo < hi {
        mi := lo + (hi - lo)/2
        //fmt.Println(lo,mi,hi)
        if nums[mi] < nums[mi+1] {
            lo = mi+1
        } else {
            hi = mi
        }
    }
    return lo
}
