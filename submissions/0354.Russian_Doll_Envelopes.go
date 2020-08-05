/*
You have a number of envelopes with widths and heights given as a pair of integers (w, h).
One envelope can fit into another if and only if both the width and height of one envelope is greater
than the width and height of the other envelope.

What is the maximum number of envelopes can you Russian doll? (put one inside other)

Note:
Rotation is not allowed.

Example:

Input: [[5,4],[6,4],[6,7],[2,3]]
Output: 3
Explanation: The maximum number of envelopes you can Russian doll is 3 ([2,3] => [5,4] => [6,7]).
*/
func maxEnvelopes(nums [][]int) int {
    L := len(nums)
    if L == 0 {return 0}
    C := make([]int, L) // parent pointer
    C[0] = 1
    M := 0
    sort.SliceStable(nums, func(i, j int) bool {
    return nums[i][0] < nums[j][0]
    })
    fmt.Println(nums)
    for i := 1;i < L; i++ {
        C[i] = 1
        M = 0
        for j:=0; j < i; j++ {
            if nums[j][1] < nums[i][1] && nums[j][0] < nums[i][0] && M < C[j] { M = C[j]}
        }
        C[i] +=M
    }
    for i := 0;i < L; i++ {
        if M < C[i] {M = C[i]}
    }
    //fmt.Println(C)
    return M

}
