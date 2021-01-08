/*
We have a two dimensional matrix A where each value is 0 or 1. A move consists of choosing any row or column, and
toggling each value in that row or column: changing all 0s to 1s, and all 1s to 0s. After making any number of moves,
every row of this matrix is interpreted as a binary number, and the score of the matrix is the sum of these numbers.
Return the highest possible score.

Example 1:

Input: [[0,0,1,1],[1,0,1,0],[1,1,0,0]]
Output: 39
Explanation:
Toggled to [[1,1,1,1],[1,0,0,1],[1,1,1,1]].
0b1111 + 0b1001 + 0b1111 = 15 + 9 + 15 = 39


Note:

1 <= A.length <= 20
1 <= A[0].length <= 20
A[i][j] is 0 or 1.
*/

func FlipRow(A [][]int, r, C int) {
    // first column is 1
    if A[r][0] == 1 {return}
    // nope flip it
    for c:=0; c < C; c++ {
        if A[r][c] == 0 {
            A[r][c] = 1
        } else {
            A[r][c] = 0
        }
    }
}

func FlipCol(A [][]int, c, R int) int{
    sum := 0
    for r := 0; r < R; r++ {
        sum += A[r][c]
    }
    // more columns are 1, so just return
    if 2*sum > R {return sum}
    return R - sum
}


func matrixScore(A [][]int) int {
    ret := 0
    R := len(A)
    if R == 0 {return 0}
    C := len(A[0])
    // Col0 is worth more than all columns 1 to N
    for r := 0; r < R; r++ {
        FlipRow(A,r,C)
    }
    mult := 1
    for c := C-1; c >= 0; c-- {
        ret += FlipCol(A,c,R) * mult
        mult *= 2
    }
    return ret
}
