/*
Given an integer matrix, find the length of the longest increasing path.

From each cell, you can either move to four directions: left, right, up or down. You may NOT move diagonally or move outside of the boundary (i.e. wrap-around is not allowed).

Example 1:

Input: nums =
[
  [9,9,4],
  [6,6,8],
  [2,1,1]
]
Output: 4
Explanation: The longest increasing path is [1, 2, 6, 9].
Example 2:

Input: nums =
[
  [3,4,5],
  [3,2,6],
  [2,2,1]
]
Output: 4
Explanation: The longest increasing path is [3, 4, 5, 6]. Moving diagonally is not allowed.
*/
func Max(x,y int) int {
    if x > y {return x}
    return y
}
func LIP(M, Mem [][]int, i, j, R, C int) (r int){
    if i < 0 || j < 0 || i >= R || j >=C {return 0}
    if Mem[i][j] > 0 {return Mem[i][j]}
    //fmt.Println(i,j)
    r = 1
    if i > 0 && M[i-1][j] > M[i][j]{
        r = LIP(M, Mem, i-1, j, R, C) + 1
    }
    if j > 0 && M[i][j-1] > M[i][j]{
        r = Max(LIP(M, Mem, i, j-1, R, C) + 1, r)
    }
    if j < C-1 && M[i][j+1] > M[i][j]{
        r = Max(LIP(M, Mem, i, j+1, R, C) + 1, r)
    }
    if i < R-1 && M[i+1][j] > M[i][j]{
        r = Max(LIP(M, Mem, i+1, j, R, C) + 1, r)
    }
    Mem[i][j] = r
    return r
}

func longestIncreasingPath(matrix [][]int) (r int) {
    R := len(matrix)
    if R == 0 {return 0}
    C := len(matrix[0])
    Mem := make([][]int, R)
    for i:= range(Mem) {
        Mem[i] = make([]int, C)
        for j:=0; j < C; j++ {
            Mem[i][j] = 0
        }
    }
    r = 0
    for i:=0; i < R; i++ {
        for j:=0; j < C; j++ {
            r = Max(r, LIP(matrix, Mem, i,j,R,C))
        }
    }
    //fmt.Println(Mem)
    return r

/* Topological sort
*/
func AddDag(M [][]int, i, j, x, y, R, C int, m map[int] []int, I []int) {
    if x < 0 || y < 0 || x == R || y == C {return}
    a := M[i][j]
    b := M[x][y]
    //fmt.Println(i,j, x,y,R,C, x*C+y, i*C+j)
    if a < b {
        m[i*C+j] = append(m[i*C+j], x*C+y)
        I[x*C+y]++
    }

}

func TopSort(M map[int] []int, I []int, N []int) (r int) {
    Q := make([]int, 0)
    Q1 := make([]int, 0)
    for i:=0; i < len(I);i++ {
        if I[i] == 0 {
            Q = append(Q, i)
        }
    }
    //fmt.Println(M,I,Q)

    for true {
        if len(Q) == 0 {
            r = r + 1
            if len(Q1) == 0 {
                break
            }
            Q = Q1
            Q1 = make([]int, 0)
        }
        n := Q[len(Q)-1]
        Q = Q[:len(Q)-1]
        for _,x:= range(M[n]) {
            I[x]--
            if I[x] == 0 {
                Q1 = append(Q1, x)
            }
        }
        //fmt.Println(n, Q, I)
    }
    return r
}

func longestIncreasingPath(matrix [][]int) (r int) {
    R := len(matrix)
    if R == 0 {return 0}
    C := len(matrix[0])
    M := make(map[int] []int)
    // Build DAG
    N := []int{}
    I := make([]int, R*C) // Incoming degree
    // Build a directed graph
    //fmt.Println(R,C)
    for i:=0; i < R;i++ {
        for j :=0; j < C ; j++ {
            //M[i*R+j] = []int{}
            AddDag(matrix,i,j,i-1,j,R,C,M, I)
            AddDag(matrix,i,j,i+1,j,R,C,M, I)
            AddDag(matrix,i,j,i,j-1,R,C,M, I)
            AddDag(matrix,i,j,i,j+1,R,C,M, I)
            N = append(N, i*R+j)
        }
    }

    return TopSort(M,I,N)
}}
