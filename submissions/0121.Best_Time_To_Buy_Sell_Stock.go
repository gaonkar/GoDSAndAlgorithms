/*
121. Best Time to Buy and Sell Stock
Easy

4659

208

Add to List

Share
Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:

Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.
Example 2:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.

*/

func maxProfit(prices []int) int {
    max:= 0
    N:= len(prices)
    i:= 1
    diff := 0
    if (N == 0) {return 0}
    min:= prices[0]
    for i < N {
        if (prices[i] < min) {
            min = prices[i]
        }
        diff = (prices[i] - min)
        if (diff > max) {
            max = diff
        }

        //fmt.Println(min, max, prices[i])
        i++
    }
    return max
}
