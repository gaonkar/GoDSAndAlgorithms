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





/*

You are given coins of different denominations and a total amount of money.
Write a function to compute the number of combinations that make up that amount.
You may assume that you have infinite number of each kind of coin.



Example 1:

Input: amount = 5, coins = [1, 2, 5]
Output: 4
Explanation: there are four ways to make up the amount:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
Example 2:

Input: amount = 3, coins = [2]
Output: 0
Explanation: the amount of 3 cannot be made up just with coins of 2.
Example 3:

Input: amount = 10, coins = [10]
Output: 1


Note:

You can assume that

0 <= amount <= 5000
1 <= coin <= 5000
the number of coins is less than 500
the answer is guaranteed to fit into signed 32-bit integer

recursive solution(Top- down approach)
time complexity - O(3^(m+n))
space complexity - O(m+n)
*/
func LCSM1(char[] X, char[] Y, int i, int j) {
        if (i <= 0 || j <= 0)
            return 0;
        if (X[i - 1] == Y[j - 1])
            return 1 + LCSM1(X, Y, i - 1, j - 1);
        else
            return Math.max(LCSM1(X, Y, i, j - 1), LCSM1(X, Y, i - 1, j));

    }


func change(amount int, coins []int) int {
    A := make([]int, amount+1)
    A[0]=1
    for j:=0; j < len(coins); j++ {
        for i:= amount; i >= coins[j]; i-- {
            A[i] += A[i-coins[j]]
        }
    }
    return A[amount]
}
