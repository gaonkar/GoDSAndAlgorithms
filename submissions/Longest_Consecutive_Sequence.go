/*
128. Longest Consecutive Sequence
Hard

2987

167

Add to List

Share
Given an unsorted array of integers, find the length of the longest consecutive elements sequence.

Your algorithm should run in O(n) complexity.

Example:

Input: [100, 4, 200, 1, 3, 2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.


*/


func longestConsecutive(nums []int) int {
    M := make(map[int] bool)
    for _,x:=range(nums) {
        M[x] = true
    }
    max := 0
    for _,x:=range(nums) {
        if M[x-1] { continue}
        cur := 0
        for M[x] {
            x++
            cur++
        }
        if cur > max { max = cur}
    }
    return max
}
