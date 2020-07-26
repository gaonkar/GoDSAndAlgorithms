/*
You have an inventory of different colored balls, and there is a customer that wants orders balls of any color.
The customer weirdly values the colored balls. Each colored ball's value is the number of balls of that color you
currently have in your inventory. For example, if you own 6 yellow balls, the customer would pay 6 for the first yellow ball.
After the transaction, there are only 5 yellow balls left, so the next yellow ball is then valued at 5 (i.e.,
the value of the balls decreases as you sell more to the customer).

You are given an integer array, inventory, where inventory[i] represents the number of balls of the ith color that you
initially own. You are also given an integer orders, which represents the total number of balls that the customer wants.
You can sell the balls in any order.

Return the maximum total value that you can attain after selling orders colored balls. As the answer may be too large, return it modulo 109 + 7.


Example 1:


Input: inventory = [2,5], orders = 4
Output: 14
Explanation: Sell the 1st color 1 time (2) and the 2nd color 3 times (5 + 4 + 3).
The maximum total value is 2 + 5 + 4 + 3 = 14.
Example 2:

Input: inventory = [3,5], orders = 6
Output: 19
Explanation: Sell the 1st color 2 times (3 + 2) and the 2nd color 4 times (5 + 4 + 3 + 2).
The maximum total value is 3 + 2 + 5 + 4 + 3 + 2 = 19.
Example 3:

Input: inventory = [2,8,4,10,6], orders = 20
Output: 110
Example 4:

Input: inventory = [1000000000], orders = 1000000000
Output: 21
Explanation: Sell the 1st color 1000000000 times for a total value of 500000000500000000. 500000000500000000 modulo 109 + 7 = 21.


Constraints:

1 <= inventory.length <= 105
1 <= inventory[i] <= 109
1 <= orders <= min(sum(inventory[i]), 109)
*/

func SUMiTOj(i,j int) int {
    ret := j * (j+1) /2
    ret -= i * (i+1) /2
    return ret
}

/*
    You have bunch of common inventory, say 10 balls of 4 colors that we need to distribute.
    But the distribution could be partial such as 3. so we need to pick 10,10,10,
*/
func PartialSum(n,c,o, MOD int) int {   //number, # of items, order
    if o == 0 {
        return 0
    }
    dec := o/c      // how many full sets
    //fmt.Println(n,c,o,dec)
    ret := (SUMiTOj(n - dec, n) * c) % MOD
    //fmt.Println(ret)
    // now partial use mod
    ret += ((n-dec) * (o%c)) % MOD
    return ret % MOD
}

func maxProfit(A []int, orders int) int {
    sort.Sort(sort.Reverse(sort.IntSlice(A)))

    i:=1 // keeps track of inventory with same count // reverse sort handles the list Ma
    cost := 0
    MOD:= 1000000007
    for i < len(A) && orders > 0 {
        inv:= A[i-1] - A[i]
        if inv == 0 {
            i++
            continue
        }
        if orders < inv * i {
            break
        }
        cost += i * SUMiTOj(A[i], A[i-1])
        orders -= inv * i
        i++
        cost %= MOD
    }
    cost += PartialSum(A[i-1],i,orders, MOD)

    return cost%MOD
}
