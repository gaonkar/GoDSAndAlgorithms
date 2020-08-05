/*
Given an array of scores that are non-negative integers. Player 1 picks one of the numbers from either end of the array followed by the player 2 and then player 1 and so on. Each time a player picks a number, that number will not be available for the next player. This continues until all the scores have been chosen. The player with the maximum score wins.

Given an array of scores, predict whether player 1 is the winner. You can assume each player plays to maximize his score.

Example 1:
Input: [1, 5, 2]
Output: False
Explanation: Initially, player 1 can choose between 1 and 2.
If he chooses 2 (or 1), then player 2 can choose from 1 (or 2) and 5. If player 2 chooses 5, then player 1 will be left with 1 (or 2).
So, final score of player 1 is 1 + 2 = 3, and player 2 is 5.
Hence, player 1 will never be the winner and you need to return False.
Example 2:
Input: [1, 5, 233, 7]
Output: True
Explanation: Player 1 first chooses 1. Then player 2 have to choose between 5 and 7. No matter which number player 2 choose, player 1 can choose 233.
Finally, player 1 has more score (234) than player 2 (12), so you need to return True representing player1 can win.
Note:
1 <= length of the array <= 20.
Any scores in the given array are non-negative integers and will not exceed 10,000,000.
If the scores of both players are equal, then player 1 is still the winner.


Recursive with Memoization
*/

func Max(x,y int) int {
    if x < y {return y}
    return x
}

func Make2D(r, c, v int) [][]int {
	ret := make([][]int, r)
	for i := 0; i < len(ret); i++ {
		ret[i] = make([]int, c)
		if v > 0 {
			for j := 0; j < c; j++ {
				ret[i][j] = v
			}
		}
	}
	return ret
}

func CanIWin(nums []int, s,e,t int, M [][]int) int {
    if s == e {
        return (t * nums[s])
    }
    if M[s][e] != 0 {return M[s][e]}
    choice1 := nums[s] * t + CanIWin(nums,s+1,e,-1*t, M) // other player plays
    choice2 := nums[e] * t + CanIWin(nums,s,e-1,-1*t, M) // other player plays

    ret := Max(t*choice1,t*choice2)*t
    M[s][e] = ret
    return ret
}
func PredictTheWinner(nums []int) bool {
    L := len(nums)
    return CanIWin(nums,0, L-1, 1, Make2D(L+1,L+1,0)) >= 0
}

/*
DP solution 1

*/

func DPCanIWin(nums []int,L int, DP[][]int) int {
    for i:=0; i < L; i++ {
        // easy to understand, each cell represents the value, and thats the only value that can be chosen
        DP[i][i] = nums[i]
    }
    // upper right triangle, fill column by column starting with first cell above the diagonal
    for j:=1; j < L;j++ {
        for i:=j-1; i >= 0; i-- {
            // substract and if the last cell is >= 0 player 1 wins
            DP[i][j] = Max(nums[j]-DP[i][j-1], nums[i]-DP[i+1][j])
        }
    }
    return DP[0][L-1]
}

func DPCanIWin2(nums []int,L int) bool {
    DP1 := Make2D(L,L,0)
    DP2 := Make2D(L,L,0)
    for i:=0; i < L; i++ {
        // player 1 gets chance 1
        DP1[i][i] = nums[i]
    }
    for j:=1; j < L;j++ {
        for i:=j-1; i >= 0; i-- {
            picki := nums[i] + DP2[i][j-1]
            pickj := nums[j] + DP2[i+1][j]
            if picki > pickj {
               DP1[i][j] = picki
               DP2[i][j] = DP1[i+1][j]
            } else {
               DP1[i][j] = pickj
               DP2[i][j] = DP1[i][j-1]
            }

        }
    }

    return DP1[0][L-1] > DP2[0][L-1]
}
