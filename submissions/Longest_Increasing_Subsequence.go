/*
300. Longest Increasing Subsequence

Share
Given an unsorted array of integers, find the length of longest increasing subsequence.

Example:

Input: [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
Note:

There may be more than one LIS combination, it is only necessary for you to return the length.
Your algorithm should run in O(n2) complexity.
Follow up: Could you improve it to O(n log n) time complexity?

*/

func lengthOfLIS(nums []int) int {
    S := []int{}
    L := len(nums)
    if L == 0 {return 0}
    S = append(S, nums[0])
    P := make([]int, L) // parent pointer
    P[0] = 1
    M := 0
    for i := 1;i < L; i++ {
        P[i] = 1
        M = 0
        for j:=0; j < i; j++ {
            if nums[j] < nums[i] && M < P[j] { M = P[j]}
        }
        P[i] +=M
    }
    for i := 0;i < L; i++ {
        if M < P[i] {M = P[i]}
    }
    //fmt.Println(P)
    return M
}
