/*
Given a set of distinct positive integers, find the largest subset such that every pair
 (Si, Sj) of elements in this subset satisfies:

Si % Sj = 0 or Sj % Si = 0.

If there are multiple solutions, return any subset is fine.

Example 1:

Input: [1,2,3]
Output: [1,2] (of course, [1,3] will also be ok)
Example 2:

Input: [1,2,4,8]
Output: [1,2,4,8]
*/



func largestDivisibleSubset(nums []int) (r []int) {
    L := len(nums)
    if L == 0 {return r}
    P := make([]int, L) // parent pointer
    C := make([]int, L) // count
    P[0] = -1
    C[0] = 1
    M := 0
    sort.Ints(nums)
    for i := 1;i < L; i++ {
        P[i] = -1
        C[i] = 1
        M = 0
        p := 0
        for j:=0; j < i; j++ {
            if nums[i] % nums[j] == 0 && M < C[j] {
                M = C[j]
                p = j
            }
        }
        C[i] +=M
        P[i] = p
    }
    p := 0
    M = C[0]
    for i:= 1; i < L; i++ {
        if M < C[i] {
            M = C[i]
            p = i
        }
    }
    //fmt.Println(M, p, nums, P, C)
    for M > 0 {
        r = append(r, nums[p])
        M--
        p = P[p]
    }
    return r
}
