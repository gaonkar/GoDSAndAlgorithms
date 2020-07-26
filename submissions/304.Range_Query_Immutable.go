/*
Given a 2D matrix matrix, find the sum of the elements inside the rectangle defined by its upper left corner (row1, col1) and lower right corner (row2, col2).

Range Sum Query 2D
The above rectangle (with the red border) is defined by (row1, col1) = (2, 1) and (row2, col2) = (4, 3), which contains sum = 8.

Example:
Given matrix = [
  [3, 0, 1, 4, 2],
  [5, 6, 3, 2, 1],
  [1, 2, 0, 1, 5],
  [4, 1, 0, 1, 7],
  [1, 0, 3, 0, 5]
]

sumRegion(2, 1, 4, 3) -> 8
sumRegion(1, 1, 2, 2) -> 11
sumRegion(1, 2, 2, 4) -> 12
*/
type NumMatrix struct {
    SM [][]int
}


func Sum2D(r,c int, M [][]int) [][] int {
    ret := make([][]int, r+1)
    for i:=0; i < len(ret); i++ {
        ret[i] = make([]int, c+1)
    }
    o := 1 // we offset everything by 1 to make life easier
    ret[1][1] = M[0][0]
    for i:=1; i < r; i++ {ret[i+o][0+o] = M[i][0] + ret[i-1+o][0+o]}
    for j:=1; j < c; j++ {ret[0+o][j+o] = M[0][j] + ret[0+o][j-1+o]}

    for i:=1; i < r; i++ {
        for j:=1; j < c; j++ {
            ret[i+o][j+o] = M[i][j] + ret[i+o][j-1+o] + ret[i-1+o][j+o] - ret[i-1+o][j-1+o]
        }
    }


    return ret
}



func Constructor(matrix [][]int) NumMatrix {
    var NM NumMatrix
    R := len(matrix)
    if (R == 0) {
        return NM
    }
    C := len(matrix[0])
    NM.SM = Sum2D(R,C, matrix)
    //fmt.Println(NM.SM)
    return NM
}


func (this *NumMatrix) SumRegion(x int, y int, A int, B int) int {
    A++
    B++
    ret := this.SM[A][B] - this.SM[A][y] - this.SM[x][B] + this.SM[x][y]
    return ret
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
