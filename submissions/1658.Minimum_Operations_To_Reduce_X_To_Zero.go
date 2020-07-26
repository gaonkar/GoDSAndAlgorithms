/*
You are given an integer array nums and an integer x. In one operation, you can either remove the leftmost or the
rightmost element from the array nums and subtract its value from x. Note that this modifies the array for future operations.

Return the minimum number of operations to reduce x to exactly 0 if it's possible, otherwise, return -1.



Example 1:

Input: nums = [1,1,4,2,3], x = 5
Output: 2
Explanation: The optimal solution is to remove the last two elements to reduce x to zero.
Example 2:

Input: nums = [5,6,7,8,9], x = 4
Output: -1
Example 3:

Input: nums = [3,2,20,1,1,3], x = 10
Output: 5
Explanation: The optimal solution is to remove the last three elements and the first two elements (5 operations in total) to reduce x to zero.


Constraints:

1 <= nums.length <= 105
1 <= nums[i] <= 104
1 <= x <= 109

*/

func SALen(nums []int, k int) int {
    c:= 0
    s:= 0
    H:=make(map[int] int)
    H[0] = -1 //sum of position upto 0
    for i,n:= range(nums) {
        s +=n
        v, ok := H[s-k]
        if ok && c < i - v {c = i -v}
        H[s]=i // store the index
    }
    return c
}

// another variant of sub array sum  wont work for negative numbers 

func minOperations(nums []int, x int) int {
    sum := 0
    for _,x := range(nums) {sum +=x}
    if sum < x {return -1}
    if sum == x {return len(nums)} // hack to distinguish no solution vs delete all
    T := sum - x
    l := SALen(nums, T)
    if l == 0 {return -1}
    return len(nums) - l
}
