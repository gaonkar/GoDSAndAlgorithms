/*
Given N, consider a convex N-sided polygon with vertices labelled A[0], A[i], ..., A[N-1] in clockwise order.

Suppose you triangulate the polygon into N-2 triangles.  For each triangle, the value of that triangle is the
product of the labels of the vertices, and the total score of the triangulation is the sum of these values over
all N-2 triangles in the triangulation.

Return the smallest possible total score that you can achieve with some triangulation of the polygon.



Example 1:

Input: [1,2,3]
Output: 6
Explanation: The polygon is already triangulated, and the score of the only triangle is 6.
Example 2:



Input: [3,7,4,5]
Output: 144
Explanation: There are two triangulations, with possible scores: 3*7*5 + 4*5*7 = 245, or 3*4*5 + 3*4*7 = 144.  The minimum score is 144.
Example 3:

Input: [1,3,1,4,1,5]
Output: 13
Explanation: The minimum score triangulation has score 1*1*3 + 1*1*4 + 1*1*5 + 1*1*1 = 13.

*/

func Min(x, y int) int {
    if x < 0 {return y}
    if x < y {return x}
    return y
}

/*
 Memoization
 */
func MinT(A []int, M [][]int, l, h int) int{
    if h - l <=1 {return 0}
    if M[l][h] >= 0 {return M[l][h]}
    m := -1
    for i:=l+1; i < h; i++ {
        x := MinT(A,M, l, i) + MinT(A, M, i, h) + A[i] * A[l] * A[h]
        //fmt.Println(l,i,h, x)
        m = Min(m, x)
    }
    M[l][h] = m
    return m
}
func minScoreTriangulation(A []int) int {
    L:= len(A)
    M:=make([][]int, L+1)
    for i:= range(M) {
        M[i] = make([]int, L+1)
        for j:=range(M[i]) {
            M[i][j] = -1
        }
    }

    return MinT(A, M, 0, L-1)
}


/*
 DP
 */
func MinDP(A []int, M [][]int) int{

    L:= len(A)
    m := -1
    for l:=L-1; l>=0; l-- {
    // Note [i][h] requires the column all the way from h to i, so l better start from bottom to fill that up
	for h:=l+1; h < L; h++ {
	// Note that [l][i] needs column from low to i. so h goes h++
	// this fill the upper right triange and the solution is in the row 0, column L-1
            m = M[l][h]
            for i := l + 1; i < h; i++ {
                x := M[l][i] + M[i][h] + A[i] * A[l] * A[h]
                //fmt.Println(l,i,h, x)
                m = Min(m, x)
            }
            M[l][h] = m
        }
    }
    //fmt.Println(M)
    return M[0][L-1]
}


