/*
Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.

Example 1:

Input: [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
Example 2:

Input: [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum
             jump length is 0, which makes it impossible to reach the last index
*/

/*
 * Solution iteration 1:
 * 	Start from the last element, keep calculating the longest jump. Once you reach 0, if that is greater than the
 * 	length of the array, then its true
 */
func min (x int, y int) int {
    if (x < y) { return x}
    return y
}
func MaxJump(nums []int, pos int) int {
    K:= nums[pos]
    M:= K
    F:= min(pos+K, len(nums)-1)
    j:=0
    for i := pos; i <= F ; i++ {
        mjump := nums[i] + j
        if (M < mjump) {
            M = mjump
        }
        //fmt.Println(pos, nums[pos], M, mjump, K, i)
        j++
    }
    return M
}
func canJump(nums []int) bool {
    L:= len(nums)
    i:=L-1
    for i >=0 {
        nums[i] = MaxJump(nums, i)
        i--
    }
    //mt.Println( nums, L)
    return nums[0] >= (L-1)
}
/*
 *
 *  Solution iteration 2:
 *  Start from Len - 2 to 0, check if you can jump over that, if yes then the next hurdle is the current index. Once you
 *  reach 0, if the hurdle is 0, then we are good
 */


func canJump(nums []int) bool {
    L:= len(nums)
    curHurdle := L-1
    for i := L-2; i >= 0; i-- {
        if curHurdle <= i + nums[i] {
            curHurdle = i
        }
    }
    return curHurdle == 0
}
