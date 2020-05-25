/*

213. House Robber II
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed.
All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one.
Meanwhile, adjacent houses have security system connected and it will automatically contact the police if two
adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house, determine the maximum
amount of money you can rob tonight without alerting the police.

Example 1:

Input: [2,3,2]
Output: 3
Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2),
             because they are adjacent houses.
Example 2:

Input: [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
             Total amount you can rob = 1 + 3 = 4.
Hint 1
Since House[1] and House[n] are adjacent, they cannot be robbed together. Therefore, the problem becomes to rob either
House[1]-House[n-1] or House[2]-House[n], depending on which choice offers more money.
Now the problem has degenerated to the House Robber, which is already been solved.
*/
func Max(x int, y int) int {
    if x < y { return y}
    return x
}
func rob1(nums []int) int {
    L := len(nums)
    if L < 1 { return 0}
    if L < 2 { return nums[0]}
    prev := nums[0]
    cur := Max(prev, nums[1])
    for i := 2; i < L; i++ {
        //fmt.Println(prev, cur, nums[i])
        new_ := nums[i] + prev
        if new_ < cur {
            new_ = cur
        }
        prev = cur
        cur = new_
    }
    return Max(cur, prev)
}

func rob(nums []int) int {
    N := len(nums)
    if N == 0 { return 0}
     if N == 1 { return nums[0]}
    if N == 2 { return Max(nums[0],nums[1])}
    return Max(rob1(nums[1:]), rob1(nums[:N-1]))
}
