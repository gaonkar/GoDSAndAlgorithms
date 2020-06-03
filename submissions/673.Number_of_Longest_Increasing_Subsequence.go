/*

Given an unsorted array of integers, find the number of longest increasing subsequence.

Example 1:

Input: [1,3,5,4,7]
Output: 2
Explanation: The two longest increasing subsequence are [1, 3, 4, 7] and [1, 3, 5, 7].
Example 2:

Input: [2,2,2,2,2]
Output: 5
Explanation: The length of longest continuous increasing subsequence is 1, and there are 5 subsequences' length is 1, so output 5.
Note: Length of the given array will be not exceed 2000 and the answer is guaranteed to be fit in 32-bit signed int.

*/

func findNumberOfLIS(nums []int) int {
    S := []int{}
    L := len(nums)
    if L == 0 {return 0}
    S = append(S, nums[0])
    P := make([]int, L) // parent pointer
    C := make([]int, L) // parent pointer
    P[0] = 1
    C[0] = 1
    M := 1
    for i := 1;i < L; i++ {
        P[i] = 1
        C[i] = 1
        if (nums[i] == nums[i-1]) {
            // take care of consecutive numbers that are equal
            P[i] = P[i-1]
            C[i] = C[i-1]
            if M < P[i] { // find the max
                M = P[i]
            }
            continue
        }
        for j:=0; j < i; j++ {
            if nums[j] < nums[i] { // valid increasing subsequence
                if P[i] == P[j] + 1 { //finding similar sequences, so add them to count
                    C[i] += C[j]
                } else if P[i] < P[j]+1 { //found a squence of length X > current, reset our counters
                    C[i] = C[j]
                    P[i] = P[j] + 1
                }
            }
        }
        if M < P[i] { //find the max seq
            M = P[i]
        }

    }
    r:= 0
    for i:=0; i < L ; i++ {
        if M == P[i] {
            r+= C[i]
        }
    }

    return r

}
