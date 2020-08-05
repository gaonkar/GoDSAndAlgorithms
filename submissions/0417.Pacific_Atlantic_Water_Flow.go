/*
Given an m x n matrix of non-negative integers representing the height of each unit cell in a continent,
the "Pacific ocean" touches the left and top edges of the matrix and the "Atlantic ocean" touches the right and bottom edges.

Water can only flow in four directions (up, down, left, or right) from a cell to another one with height equal or lower.

Find the list of grid coordinates where water can flow to both the Pacific and Atlantic ocean.

Note:

The order of returned grid coordinates does not matter.
Both m and n are less than 150.


Example:

Given the following 5x5 matrix:

  Pacific ~   ~   ~   ~   ~
       ~  1   2   2   3  (5) *
       ~  3   2   3  (4) (4) *
       ~  2   4  (5)  3   1  *
       ~ (6) (7)  1   4   5  *
       ~ (5)  1   1   2   4  *
          *   *   *   *   * Atlantic

Return:

[[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]] (positions with parentheses in above matrix).
*/


func DFS(m [][]int, v [][]bool,  re [][]int, i int, j int, R int, C int) {
    if v[i][j] {
        return
    }
    v[i][j] = true
    re[i][j]++
    if i > 0  && m[i][j] <= m[i-1][j]{
        DFS(m, v, re, i-1, j, R, C)
    }
    if i < R - 1 && m[i][j] <= m[i+1][j]{
        DFS(m, v, re, i+1, j, R, C)
    }
    if j > 0 && m[i][j] <= m[i][j-1]{
        DFS(m, v, re, i, j-1, R, C)
    }
    if j < C - 1 && m[i][j] <= m[i][j+1]{
        DFS(m, v, re, i, j+1, R, C)
    }
}
func pacificAtlantic(matrix [][]int) (r [][]int) {
    R :=len(matrix)
    if R == 0 {return nil}
    C := len(matrix[0])
    if C == 0 {return nil}
    v := make([][]bool, R)
    re := make([][]int, R)
    rows := make([]int, R*C)
    for i := 0; i < R; i++ {
        re[i] = rows[i*C : (i+1)*C]
        v[i] = make([]bool, C)
    }
    for i := 0; i < R; i++ {
        DFS(matrix, v, re, i, 0, R, C)
    }
    for j := 0; j < C; j++ {
        DFS(matrix, v, re, 0, j, R, C)
    }

    for i := 0; i < R; i++ {
        v[i] = make([]bool, C)
    }
    for i := 0; i < R; i++ {
        DFS(matrix, v, re, i, C-1, R, C)
    }
    for j := 0; j < C; j++ {
        DFS(matrix, v, re, R-1, j, R, C)
    }
    for i := 0; i < R; i++ {
        for j := 0; j < C; j++ {
            if re[i][j] > 1 {
                r = append(r, []int{i,j})
            }
        }
    }
    return r
}
