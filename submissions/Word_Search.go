/*
79. Word Search
Medium

3297

163

Add to List

Share
Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.

Example:

board =
[
  ['A','B','C','E'],
    ['S','F','C','S'],
      ['A','D','E','E']
      ]

      Given word = "ABCCED", return true.
      Given word = "SEE", return true.
      Given word = "ABCB", return false.


      Constraints:

      board and word consists only of lowercase and uppercase English letters.
      1 <= board.length <= 200
      1 <= board[i].length <= 200
      1 <= word.length <= 10^3


*/


func SetM(m map[int] bool, i int, j int, C int) {
    x := i * C + j
    m[x] = true
}
func UnSetM(m map[int] bool, i int, j int, C int) {
    x := i * C + j
    m[x] = false
}
func IsSetM(m map[int] bool, i int, j int, C int) bool {
    x := i *C + j
    return m[x]
}
func EHelper(b [][]byte, w string, i int, j int, m map[int] bool, R int, C int) bool {
    found := false
    W := len(w)

    if W == 0 {
        return found
    }
    //fmt.Println(m,w, i, j)
    //fmt.Println(b[i][j], w[0])
    if w[0] != b[i][j] { return false}
    if W == 1 {
        //fmt.Println("here")
        return true
    }
    SetM(m,i,j,C)
    //fmt.Println(m,w, ",i=", i,",j=", j, string(b[i][j]),"=", string(w[0]), ",I=",i*C+ j)


    if i > 0  &&  IsSetM(m,i-1,j,C)==false {
        //fmt.Println("i>0")
        found = EHelper(b,w[1:],i-1,j,m,R,C)
    }
    if i < R-1  &&  IsSetM(m,i+1,j,C)==false && !found {
        //fmt.Println("i<R-1")
        found = EHelper(b,w[1:],i+1,j,m,R,C)
    }
    if j > 0  &&  IsSetM(m,i,j-1,C)==false && !found{
        //fmt.Println("j>0")
        found = EHelper(b,w[1:],i,j-1,m,R,C)
    }
    if j < C-1  &&  IsSetM(m,i,j+1,C)==false && !found {
        //fmt.Println("j<C-1")
        found = EHelper(b,w[1:],i,j+1,m,R,C)
    }
    UnSetM(m,i,j,C)
    return found
}

func exist(board [][]byte, word string) bool {
    m := make(map[int] bool)
    W := len(word)
    if W == 0 {return true}
    R := len(board)
    if R == 0 {return false}
    C := len(board[0])
    if C == 0 {return false}

    //find first character
    i := 0
    j := 0
    found := false
    //fmt.Println(R,C)
    for i < R  {
        j = 0
        for j < C {
            if word[0] == board[i][j] {
                //m = make(map[int] bool)
                found = EHelper(board, word, i, j, m, R, C)
                //fmt.Println(word, i, j, found, m)
                if found { return true}
            }
            j++
        }
        i++
    }
    return found
}
