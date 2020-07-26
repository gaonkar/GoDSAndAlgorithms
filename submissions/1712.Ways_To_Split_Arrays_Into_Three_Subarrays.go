/*
A split of an integer array is good if:

The array is split into three non-empty contiguous subarrays - named left, mid, right respectively from left to right.
The sum of the elements in left is less than or equal to the sum of the elements in mid, and the sum of the elements in
mid is less than or equal to the sum of the elements in right.
Given nums, an array of non-negative integers, return the number of good ways to split nums. As the number may be
too large, return it modulo 109 + 7.



Example 1:

Input: nums = [1,1,1]
Output: 1
Explanation: The only good way to split nums is [1] [1] [1].
Example 2:

Input: nums = [1,2,2,2,5,0]
Output: 3
Explanation: There are three good ways of splitting nums:
[1] [2] [2,2,5,0]
[1] [2,2] [2,5,0]
[1,2] [2,2] [5,0]
Example 3:

Input: nums = [3,2,1]
Output: 0
Explanation: There is no good way to split nums.


Constraints:

3 <= nums.length <= 105
0 <= nums[i] <= 104
*/


/*
    [...i...j..k]
    After computing prefix 1
  condition 1
    sum(0..i) <= sum(i+1..j)
    P[i] <= P[j] - P[i]
    2*P[i] <= P[j]

 condition 2
    sum(j+1..k) <= sum(k..N)
    P[k] - P[i] <= SN - P[k]

  Loop
    One could use O(N^3), however note that for a given i, j and k will only be increasing as i is incremented.
    Since there are no negative numbers in this list.


*/
func waysToSplit(nums []int) int {
    P:= make([]int, len(nums))
    for i,_ := range(nums) {
        P[i] = nums[i]
        if i > 0 {
            P[i]+= P[i-1]
        }
    }
    N := len(nums)
    i,j,k := 0,1,2
    count := 0
    for i < len(nums)-2 {
        for j <= i || (P[i] * 2 > P[j] && j < N-1) {j++}
        for k < j || (k < N-1 && 2*P[k] - P[i] <= P[N-1]){k++}
        count = (count + k - j) % 1000000007;
        i++
    }

    return count
}
