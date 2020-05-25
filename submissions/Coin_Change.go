/*
322. Coin Change

You are given coins of different denominations and a total amount of money amount. Write a function to compute the
fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any
combination of the coins, return -1.

Example 1:

Input: coins = [1, 2, 5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
Example 2:

Input: coins = [2], amount = 3
Output: -1
Note:
You may assume that you have an infinite number of each kind of coin.
*/

func coinChange(coins []int, amount int) int {
    S := make([]int, amount+1)
    S[0] = 0
    for i:=1; i <= amount; i++ {
        M := -1
        for _,c:=range(coins) {
            p := i - c
            if p < 0 || S[p] < 0 {continue}
            if M < 1 {
                M = S[p] + 1
            } else if M > S[p] + 1 {
                M = S[p] + 1
            }
        }
        S[i] = M
    }
    fmt.Println(S)
    return S[amount]
}
