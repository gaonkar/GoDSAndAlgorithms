/*
We have two types of tiles: a 2x1 domino shape, and an "L" tromino shape. These shapes may be rotated.

XX  <- domino

XX  <- "L" tromino
X
Given N, how many ways are there to tile a 2 x N board? Return your answer modulo 10^9 + 7.

(In a tiling, every square must be covered by a tile. Two tilings are different if and only if there are two 4-directionally adjacent cells on the board such that exactly one of the tilings has both squares occupied by a tile.)

//difficult to draw them in ASCII
D[N] = D[N-1] + D[N-2] + 2 * sum(D[i]) for i = 0 to N-3

1 way to translate n-1, and n-2 to N and 2 ways for the rest

Solving the relation we get
D[N] = 2 * D[N-1] + D[n-1]

this is like Fibonnaci series, time spent 2 weeks
*/

func numTilings(N int) int {
    mod := int(math.Pow(10,9)) + 7
    D := make([]int, N+1)
    D[0] = 1
    D[1] = 1
    D[2] = 2
    for i:=3; i <=N; i++ {
        D[i] = D[i-1]*2 + D[i-3]
        D[i] %=mod
    }
    return D[N]
}

