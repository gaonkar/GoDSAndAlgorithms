/*
You are given the number of rows n_rows and number of columns n_cols of a 2D binary matrix where all values are initially 0.
Write a function flip which chooses a 0 value uniformly at random, changes it to 1,
and then returns the position [row.id, col.id] of that value. Also, write a function reset which sets all values back to 0.
Try to minimize the number of calls to system's Math.random() and optimize the time and space complexity.

Note:

1 <= n_rows, n_cols <= 10000
0 <= row.id < n_rows and 0 <= col.id < n_cols
flip will not be called when the matrix has no 0 values left.
the total number of calls to flip and reset will not exceed 1000.
Example 1:

Input:
["Solution","flip","flip","flip","flip"]
[[2,3],[],[],[],[]]
Output: [null,[0,1],[1,2],[1,0],[1,1]]
Example 2:

Input:
["Solution","flip","flip","reset","flip"]
[[1,2],[],[],[],[]]
Output: [null,[0,0],[0,1],null,[0,0]]
Explanation of Input Syntax:

The input is two lists: the subroutines called and their arguments.
Solution's constructor has two arguments, n_rows and n_cols. flip and reset have no arguments.
Arguments are always wrapped with a list, even if there aren't any.

Read https://en.wikipedia.org/wiki/Fisher–Yates_shuffle#Examples

*/

type Solution struct {
    R,C,T int

    M map[int] int
}


func Constructor(n_rows int, n_cols int) Solution {
    var S Solution
    S.R = n_rows
    S.C = n_cols
    S.Reset()
    return S
}



func GetOrDefaultSame(M map[int] int, val int) int {
    v, ok := M[val]
    if ok {
        return v
    }
    return val
}

func (this *Solution) Flip() []int {
    r := rand.Intn(this.T)
    cols := this.R
    v := GetOrDefaultSame(this.M, r)

    this.T--
    this.M[r] = GetOrDefaultSame(this.M,this.T)
    fmt.Println(v, r, this.T+1, this.M)
    return []int{v/cols, v%cols}
}


func (this *Solution) Reset()  {
    this.T = this.R * this.C
    this.M = map[int] int{}
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n_rows, n_cols);
 * param_1 := obj.Flip();
 * obj.Reset();
 */
