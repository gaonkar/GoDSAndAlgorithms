/*

Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete at most two transactions.

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

Example 1:

Input: [3,3,5,0,0,3,1,4]
Output: 6
Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
             Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
Example 2:

Input: [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
             Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
             engaging multiple transactions at the same time. You must sell before buying again.
Example 3:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
*/

func Make2D(r,c, v int) [][] int {
    ret := make([][]int, r)
    for i:=0; i < len(ret); i++ {
        ret[i] = make([]int, c)
        if v > 0 {
            for j := 0; j < c; j++ {
                ret[i][j] = v
            }
        }
    }
    return ret
}

func Max(x,y int) int {
    if y < 0 {return x}
    if x > y {return x}
    return y
}

func maxProfit1(prices []int) int {
    L := len(prices)
    if L == 0 {return 0}
    DP := Make2D(2, L,0)
    M  := Make2D(2, L,0)
    // lets first solve for 1 transaction, track the last seen min price, compute the best profit
    min := prices[0]
    for i:=1; i < L;i++ {
        if min > prices[i-1] {
            min = prices[i-1]
        }
        DP[0][i] = Max(DP[0][i-1], prices[i]-min)
    }
    //second iteration
    min = prices[0]
    for i:=1; i < L;i++ {
        s := prices[i] - DP[0][i-1]
        if min > s{
            min = s
        }
        DP[1][i] = Max(DP[1][i-1], prices[i]-min)
    }
    fmt.Println(DP, M)
    return DP[1][L-1]
}

func maxProfit(prices []int) int {
    L := len(prices)
    if L == 0 {return 0}
    K:=2
    DP := make([]int, L)


    //second iteration
    for K > 0 {
        min := prices[0]
        prev := DP[0] // prev tracks the old value before it is over written
        for i:=1; i < L;i++ {
            s := prices[i] - prev
            if min > s{
                min = s
            }
            prev = DP[i]
            DP[i] = Max(DP[i-1], prices[i]-min)
        }
        K--
    }
    fmt.Println(DP)
    return DP[L-1]
}

max(x, y int) int {
    if x > y {return x}
    return y
}

/*
user dietpepsi thinking
The natural states for this problem is the 3 possible transactions : buy, sell, rest. Here rest means no transaction on that day (aka cooldown).

Then the transaction sequences can end with any of these three states.
For each of them we make an array, buy[n], sell[n] and rest[n].
buy[i] means before day i what is the maxProfit for any sequence end with buy.
sell[i] means before day i what is the maxProfit for any sequence end with sell.
rest[i] means before day i what is the maxProfit for any sequence end with rest.
Then we want to deduce the transition functions for buy sell and rest

The fact that buy[i] <= rest[i] which means rest[i] = max(sell[i-1], rest[i-1]).
That made sure [buy, rest, buy] is never occurred.

A further observation is that and rest[i] <= sell[i] is also true therefore

    rest[i] = sell[i-1]
Substitute this in to buy[i] we now have 2 functions instead of 3:

buy[i] = max(sell[i-2]-price, buy[i-1])
sell[i] = max(buy[i-1]+price, sell[i-1])

 */
func maxProfit(prices []int) int {
    L := len(prices)
    buy := make([]int, L)
    sell := make([]int, L)
    rest := make([]int, L)
    buy[0] = -prices[0]

    for i:=1; i < L; i++ {
        buy[i]  = max(rest[i-1]-prices[i], buy[i-1])
        sell[i] = max(buy[i-1]+prices[i], sell[i-1])
        rest[i] = max(sell[i-1], max(buy[i-1], rest[i-1]))
    }
    fmt.Println(buy,sell,rest)
    return max(max(buy[L-1], sell[L-1]), rest[L-1])
}
