/*
Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the
very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.
Return the max sliding window.

Follow up:
Could you solve it in linear time?

Example:

Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
Output: [3,3,5,5,6,7]
Explanation:

Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
1 [3  -1  -3] 5  3  6  7       3
1  3 [-1  -3  5] 3  6  7       5
1  3  -1 [-3  5  3] 6  7       5
1  3  -1  -3 [5  3  6] 7       6
1  3  -1  -3  5 [3  6  7]      7


Constraints:

1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
1 <= k <= nums.length
*/


func maxSlidingWindow(nums []int, k int) (r []int) {
    mq := make([]int, 0) // keep the index else we will have to keep 2 arrays
    n := 0
    for i,x := range(nums) {
        for n > 0 && x >= nums[mq[n-1]] { // you saw 5 thats greater than all we saw, pop them out
            n--
            mq = mq[:n]
        }
        mq = append(mq, i)
        n++
        if mq[0] == i-k {
            mq = mq[1:]
            n--
        }
        if i >= k-1 {
            r = append(r, nums[mq[0]])
        }
        //fmt.Println(mq)
    }
    return r
}
