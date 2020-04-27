/*
Given a 2D binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.

Example:

Input:

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

Output: 4

Intuition
    1. Ignore Row 0 and col 0
    2. For M[i,j] = if M[i,j] == 1 then min(M[i,j-1],M[i-1,j-1], M[i-1][j]) + 1 else 0

*/

func min (x byte, y byte) byte {
    if (x < y) { return x}
    return y
}
func max (x int, y byte) int {
    if (x > int(y)) { return x}
        return int(y)
}
func maximalSquare(matrix [][]byte) int {
    L:= len(matrix)
    if (L == 0) {
        return 0
    }
    W:= len(matrix[0])
    if (W == 0) {
        return 0
    }
    R:= int(matrix[0][0] -'0')

    for i:=0 ; i < L; i++ {
        matrix[i][0] = matrix[i][0] - '0'
        if (matrix[i][0] > 0) {R = 1}
        for j:=1; j < W; j++ {
            matrix[i][j] -= '0'
            if (matrix[i][j] > 0) {
                if (i > 0) {
                    matrix[i][j] += matrix[i][j-1]
                }
                R = 1
            }
        }
    }
    //fmt.Println(matrix)
    for i:=1 ; i < L; i++ {
        for j:=1; j < W; j++ {
            if (matrix[i][j] > 0) {
                matrix[i][j] = min(min(matrix[i][j], matrix[i-1][j] + 1), matrix[i-1][j-1] + 1)
            }
            R = max(R, matrix[i][j])
        }
    }
    //fmt.Println(matrix)

    return R*R
}
