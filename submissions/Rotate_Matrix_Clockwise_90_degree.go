/*
You are given an n x n 2D matrix representing an image.

Rotate the image by 90 degrees (clockwise).

Note:

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

Example 1:

Given input matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

rotate the input matrix in-place such that it becomes:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
 */

func SwapCell(matrix [][]int, r1 int, c1 int, r2 int, c2 int) {
    temp:= matrix[r1][c1]
    matrix[r1][c1] = matrix[r2][c2]
    matrix[r2][c2] = temp
}


func Swap(matrix [][]int, N int) {
    for i:=0;i < N/2; i++ {
        for j:= range matrix {
            //fmt.Println(i,j,matrix[j][i], matrix[j][N-i-1])
            SwapCell(matrix,j,i, j, N-i-1)
        }
    }
}

func Transpose(matrix [][]int, N int) {
     for i:= range matrix {
        for j:= 0; j < i; j++ {
            //fmt.Println(i,j,matrix[i][j], matrix[j][i])
            SwapCell(matrix,i,j, j, i)
        }
    }
}

func rotate(matrix [][]int)  {
    N := len(matrix)
    Transpose(matrix,N)
    //fmt.Println(matrix)
    Swap(matrix,N)
    //fmt.Println(matrix)
}
