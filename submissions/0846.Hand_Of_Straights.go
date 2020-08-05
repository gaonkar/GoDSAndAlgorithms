/*
Alice has a hand of cards, given as an array of integers.

Now she wants to rearrange the cards into groups so that each group is size W, and consists of W consecutive cards.

Return true if and only if she can.



Example 1:

Input: hand = [1,2,3,6,2,3,4,7,8], W = 3
Output: true
Explanation: Alice's hand can be rearranged as [1,2,3],[2,3,4],[6,7,8].
Example 2:

Input: hand = [1,2,3,4,5], W = 4
Output: false
Explanation: Alice's hand can't be rearranged into groups of 4.


Note:

1 <= hand.length <= 10000
0 <= hand[i] <= 10^9
1 <= W <= hand.length

*/

func isPossible(nums []int, s int) bool {
    M := make(map[int] int)
    L:= len(nums)
    if L < s { return false}
    for _,n := range(nums) {
        M[n]++
    }
    lo:= 0

    for lo < L { // O(L * L / s) N square
        if M[nums[lo]] == 0 {
            lo++
            continue
        }
        // start with lowest card, now find s cards and eliminate them, then
        M[nums[lo]]--
        hi := nums[lo] + 1

        for M[hi] > 0  && hi < nums[lo] + s{
            M[hi]--
            hi++
        }
        //fmt.Println(M)
        if hi - nums[lo] != s {return false}
    }
    return true
}
func isNStraightHand(hand []int, W int) bool {
    sort.Ints(hand) // n log n
    fmt.Println(hand)
    return  isPossible(hand, W)
}
