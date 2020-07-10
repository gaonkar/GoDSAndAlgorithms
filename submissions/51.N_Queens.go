/*
The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.



Given an integer n, return all distinct solutions to the n-queens puzzle.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.


 look at chess board from left to right (a=1, h = 8)
 We can try out all possible values starting from cell(1,1)--> cell(8,8)

 Suppose, we pick cell(0,4), then all columsn=4 is invalid, we use C map to track that
 now [1,5], [2,6] so on are invalid, so 0-4 == 1-5 == 2-6, so we track that in N map
 also [1,3] [2,2] are invalid --> 1+3 == 2+2 == 0+4 is invalid, we track that in M map

 now thats a recursion on what not to choose, since x is always incremented, we dont have to
 worry about it

*/

func diff (x,y int) int {
    if x < y {return y - x}
    return x - y
}

var r [][]int
func NQueens(n int,q []int, C, M, N map[int] bool) {
    x := len(q)
    if x == n {
        tmp := make([]int, n)
        copy(tmp, q)
        r = append(r, tmp)
        return
    }
    //fmt.Println(x, q, C, M, N)
    for y := 0; y < n; y++ {
        //fmt.Println(y, x+y, y-x, y)
        if M[x+y] || N[y-x] || C[y] {continue}
        M[x+y] = true
        N[y-x] = true
        C[y] = true
        q = append(q, y)
        NQueens(n ,q, C, M, N)
        q = q[:x]
        delete(M, x+y)
        delete(N, y-x)
        delete(C,y)
    }
}

func MakeString(p,n int) string {
    b := make([]byte, n)
    for i:=0; i < n ;i++{b[i] = '.'}
    b[p] = 'Q'
    return string(b)
}

func solveNQueens(n int) (x [][]string){
    r = make([][]int, 0)
    if n < 1 {return nil}
    NQueens(n, []int{},  make(map[int] bool), make(map[int] bool), make(map[int] bool))
    fmt.Println(r)
    for i,row := range(r) {
        x = append(x, []string{})
        for _,j := range(row) {
            x[i] = append(x[i],MakeString(j,n))
        }
    }
    return x
}
