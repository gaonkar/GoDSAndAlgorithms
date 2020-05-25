/*
Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in-place.

Example 1:

Input:
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
Output:
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
Example 2:

Input:
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
Output:
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
Follow up:

A straight forward solution using O(mn) space is probably a bad idea.
A simple improvement uses O(m + n) space, but still not the best solution.
Could you devise a constant space solution?


*/

func ZeroRow(M [][]int, r int, C int) {
    for j:=0; j < C;j++ {M[j][r] = 0}
}
func ZeroCol(M [][]int, r int, C int) {
     for j:=0; j < C;j++ {M[r][j] = 0}
}
func setZeroes(matrix [][]int)  {
    R := len(matrix)
    if R == 0 {return}
    C := len(matrix[0])
    if C == 0 {return}
    row0 := false
    col0 := false
    rnd :=0
    for i:=0; i < R ; i++ {
        for j:=0; j < C;j++ {
            if matrix[i][j] == 0 {
                if i == 0 {
                    row0 = true
                } else {
                    matrix[i][0] = -rnd
                }
                if j == 0 {
                    col0 = true
                } else {
                    matrix[0][j] = -rnd
                }
            }
        }
    }
    fmt.Println(row0,col0)
    for i:= 1; i < R; i++ {
        if matrix[i][0] == -rnd {
            ZeroCol(matrix,i,C)
        }
    }
    for i:= 1; i < C; i++ {
        if matrix[0][i] == -rnd {
            ZeroRow(matrix,i, R)
        }
    }
    if row0 {
        ZeroCol(matrix,0, C)
    }
    if col0 {
        ZeroRow(matrix,0, R)
    }

}
