/*
You are given n pairs of numbers.
In every pair, the first number is always smaller than the second number.

Now, we define a pair (c, d) can follow another pair (a, b) if and only if b < c.
Chain of pairs can be formed in this fashion.

Given a set of pairs, find the length longest chain which can be formed.
You needn't use up all the given pairs. You can select pairs in any order.

Example 1:
Input: [[1,2], [2,3], [3,4]]
Output: 2
Explanation: The longest chain is [1,2] -> [3,4]
Note:
The number of given pairs will be in the range [1, 1000].
*/

func findLongestChain(nums [][]int) int {
    L := len(nums)
    if L == 0 {return 0}
    C := make([]int, L) // parent pointer
    C[0] = 1
    M := 0
    sort.SliceStable(nums, func(i, j int) bool {
    return nums[i][0] < nums[j][0]
    })
    for i := 1;i < L; i++ {
        C[i] = 1
        M = 0
        for j:=0; j < i; j++ {
            if nums[j][1] < nums[i][0] && M < C[j] { M = C[j]}
        }
        C[i] +=M
    }
    for i := 0;i < L; i++ {
        if M < C[i] {M = C[i]}
    }
    //fmt.Println(C)
    return M

}
