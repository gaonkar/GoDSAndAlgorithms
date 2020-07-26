/*
A sequence of numbers is called a wiggle sequence if the differences between successive numbers strictly alternate between positive and negative. The first difference (if one exists) may be either positive or negative. A sequence with fewer than two elements is trivially a wiggle sequence.

For example, [1,7,4,9,2,5] is a wiggle sequence because the differences (6,-3,5,-7,3) are alternately positive and negative. In contrast, [1,4,7,2,5] and [1,7,4,5,5] are not wiggle sequences, the first because its first two differences are positive and the second because its last difference is zero.

Given a sequence of integers, return the length of the longest subsequence that is a wiggle sequence. A subsequence is obtained by deleting some number of elements (eventually, also zero) from the original sequence, leaving the remaining elements in their original order.

Example 1:

Input: [1,7,4,9,2,5]
Output: 6
Explanation: The entire sequence is a wiggle sequence.
Example 2:

Input: [1,17,5,10,13,15,10,5,16,8]
Output: 7
Explanation: There are several subsequences that achieve this length. One is [1,17,10,13,10,16,8].
Example 3:

Input: [1,2,3,4,5,6,7,8,9]
Output: 2
*/

func Max(x, y int) int {
 if x > y {
   return x
 }
 return y
}
/*
    U tracks the subsequence where the last 2 numbers ended with positive value
    D tracks the last 2 numbers with negative value
 */
func wiggleMaxLength(nums []int) int {
    L := len(nums)
    if L <= 1 {return L}
    U := make([]int, L)
    D := make([]int, L)
    u,d :=1,1
    U[0], D[0] = 1, 1
    for i:=1; i < L; i++ {
        if nums[i-1] < nums[i] {
            // add 1 to down sequence and assign it to u as it ends with positive value
            u = d + 1
            U[i] = D[i-1]+1
            D[i] = U[i-1]
        } else if nums[i-1] > nums[i] {
            d = u + 1
            D[i] = U[i-1]+1
            U[i] = D[i-1]
        } else {
            // ignore as it adds no value
            U[i] = U[i-1]
            D[i] = D[i-1]

        }
    }
    /*
    1. Learnt from the discussion that these are cumulative, so one does not need to keep track in array, but the last value is good.
    2. Spent considerable time figuring out that we need 2 arrays to track Up or Down.
    */

    return Max(u,d)
}
