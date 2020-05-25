/*
Input: matrix =
[
  [0,1,1,1],
  [1,1,1,1],
  [0,1,1,1]
]
Output: 15
Explanation:
There are 10 squares of side 1.
There are 4 squares of side 2.
There is  1 square of side 3.
Total number of squares = 10 + 4 + 1 = 15.
Example 2:

Input: matrix =
[
  [1,0,1],
  [1,1,0],
  [1,1,0]
]
Output: 7
Explanation:
There are 6 squares of side 1.
There is 1 square of side 2.
Total number of squares = 6 + 1 = 7.

*/
func countSquares(matrix [][]int) int {
    R := len(matrix)
    if R == 0 {return 0}
    C := len(matrix[0])
    for i:= range(matrix) {
        for j:= range(matrix[0]) {
            if j > 0 { matrix[i][j] += matrix[i][j-1]}
            if i > 0 {matrix[i][j] += matrix[i-1][j]}
            if i > 0 && j > 0 { matrix[i][j] -= matrix[i-1][j-1]}
        }
    }
    K := R
    if K > C {K = C}
    fmt.Println(matrix)
    count := matrix[R-1][C-1]
    for k := 1; k < K; k++ {
        mm := (k+1) * (k+1)
        for i:=0; i < len(matrix) -k; i++ {
            for j:=0; j <  len(matrix[0])-k; j++ {
                BR := matrix[i+k][j+k]
                TL := 0
                if i > 0 && j > 0 { TL = matrix[i-1][j-1]}
                TR := 0
                if i > 0 { TR = matrix[i-1][j+k]}
                BL := 0
                if j > 0 { BL = matrix[i+k][j-1]}
                diff := BR +TL - BL -TR
                //fmt.Println("k=", k+1, i,j, TL, TR, BL, BR, diff, mm)
                if (mm == diff) {count++}
            }
        }
    }

    return count
}

/* DP solution
    for i = 0 or j = 0, the maximum matrix that can be is size = 1
    now for i > 0 and j > 0 and the matrix[i][j] == 1, count if you can make  any larger square by looking at [i-1,j-1], [i,j-1] and
    [i-1,j]
 */
func countSquares(matrix [][]int) int {
    row := len(matrix)
    col := len(matrix[0])
    if row == 0 || col == 0 {
        return 0
    }
    ans := 0
    var min int
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            if matrix[i][j] > 0 && i > 0 && j > 0 {
                min = matrix[i-1][j-1]
                if min > matrix[i-1][j] {
                    min = matrix[i-1][j]
                }
                if min > matrix[i][j-1] {
                    min = matrix[i][j-1]
                }
                matrix[i][j] = min + 1
            }
            ans += matrix[i][j]
        }
    }
    fmt.Println(matrix)
    return ans
}
