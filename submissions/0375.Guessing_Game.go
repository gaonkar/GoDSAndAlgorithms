/*
We are playing the Guess Game. The game is as follows:

I pick a number from 1 to n. You have to guess which number I picked.

Every time you guess wrong, I'll tell you whether the number I picked is higher or lower.

However, when you guess a particular number x, and you guess wrong, you pay $x. You win the game when you guess the number I picked.

Example:

n = 10, I pick 8.

First round:  You guess 5, I tell you that it's higher. You pay $5.
Second round: You guess 7, I tell you that it's higher. You pay $7.
Third round:  You guess 9, I tell you that it's lower. You pay $9.

Game over. 8 is the number I picked.

You end up paying $5 + $7 + $9 = $21.
Given a particular n ≥ 1, find out how much money you need to have to guarantee a win.
*/

/*
 Solution:
  Lets define DP[i][j] be equal to the money required for guessing a number thats in the range i to j
  We pick a number x, we end up paying x, but now what is the worst we have to pay that would be
    Max(DP[i][x-1], DP[x+1][j]) + x
  So for we now know thats the cost of choosing x in range i to j. We now try all values from i to j
  to come up with minimum that we could choose.
  for x = i; x <= j; x++ {
     //find value of choosing x,
     pick the minimum value
  }
}

*/

// from stack exchange
const (
	MinUint uint = 0 // binary: all zeroes

	// Perform a bitwise NOT to change every bit from 0 to 1
	MaxUint = ^MinUint // binary: all ones

	// Shift the binary number to the right (i.e. divide by two)
	// to change the high bit to 0
	MaxInt = int(MaxUint >> 1) // binary: all ones except high bit

	// Perform another bitwise NOT to change the high bit to 1 and
	// all other bits to 0
	MinInt = ^MaxInt // binary: all zeroes except high bit
)

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

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func DP(dp [][]int, l, h int) int {
	if l >= h {
		return 0
	}
	if dp[l][h] > 0 {
		return dp[l][h]
	}
	ret := MaxInt
	for i := l; i <= h; i++ {
		tmp := Max(DP(dp, l, i-1), DP(dp, i+1, h)) + i
		ret = Min(ret, tmp)
	}
	dp[l][h] = ret
	return ret
}

func getMoneyAmount(n int) int {
	dp := Make2D(n+1, n+1, 0)
	ret := DP(dp, 1, n)
	//fmt.Println(dp)
	return ret
}
