/*
659. Split Array into Consecutive Subsequences

Share
Given an array nums sorted in ascending order, return true if and only if you can split it into 1 or more
subsequences such that each subsequence consists of consecutive integers and has length at least 3.



Example 1:

Input: [1,2,3,3,4,5]
Output: True
Explanation:
You can split them into two consecutive subsequences :
1, 2, 3
3, 4, 5

Example 2:

Input: [1,2,3,3,4,4,5,5]
Output: True
Explanation:
You can split them into two consecutive subsequences :
1, 2, 3, 4, 5
3, 4, 5

Example 3:

Input: [1,2,3,4,4,5]
Output: False


Constraints:

1 <= nums.length <= 10000
*/

func isPossible(nums []int) bool {
    M := make(map[int] int)
    L:= len(nums)
    if L < 2 { return false}
    // Generate frequency
    for _,n := range(nums) {
        M[n]++
    }
    lo:= 0
    for lo < L {
        if M[nums[lo]] == 0 { // object has been consumed
            lo++
            continue
        }
        p := M[nums[lo]]
        M[nums[lo]]--
        hi := nums[lo] + 1
        // add all the increasing numbers into this subsequence until you find it does not exist in list or frequency of
        // the previous > current number because by taking the current, you might end up leaving previous orphan
        for M[hi] > 0 {
            if M[hi] >= p {
                p = M[hi]
                M[hi]--
            } else {
                break
            }
            hi++
        }
        //fmt.Println(M)
        if hi - nums[lo] <= 2 {return false}
        lo++
    }
    return true
}

