/*
Given n balloons, indexed from 0 to n-1. Each balloon is painted with a number on it represented by array nums.
You are asked to burst all the balloons. If the you burst balloon i you will get nums[left] * nums[i] * nums[right]
coins. Here left and right are adjacent indices of i. After the burst, the left and right then becomes adjacent.

Find the maximum coins you can collect by bursting the balloons wisely.

Note:

You may imagine nums[-1] = nums[n] = 1. They are not real therefore you can not burst them.
0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
Example:

Input: [3,1,5,8]
Output: 167
Explanation: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
             coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167

For multiplication replace Max by Min

*/


func Max(x, y int) int {
    if x < y {return y}
    return x
}

/*
 Memoization
 */
func MinT(A []int, D [][]int, l, h int) int{
    if h == l {
        return 0
    }
    if D[l][h] > 0 {return D[l][h]}

    m := 0
    for i:=l; i < h; i++ {
        P := MinT(A,D, l, i)
        Af := MinT(A, D, i+1, h)
        coins := A[l-1] * A[i] *A[h]

        x :=  P + Af +  coins
        m = Max(m, x)
        //fmt.Println(l,i,h, P, Af, coins, m)
    }

    D[l][h] = m
    return m
}


func maxCoins(A []int) int {
    A = append([]int{1}, A...)
    A = append(A, 1)
    L:= len(A)
    D:=make([][]int, L+1)
    for i:= range(D) {
        D[i] = make([]int, L+1)
    }

    y := MinT(A, D, 1, L-1)
    //for i:=0; i < L+1;i++ {    fmt.Println(D[i])}
    return y
}

/*
 DP
 [3,2,5,7]--> [1,3,2,5,7,1] L = 6
 the way to think of DP is as follows,
 [1,3,2,)5,7(,1] Suppose we are looking at should we remove 5 first or 7 first.
 Here l = 3, h = 4
 We would have already computed M[3][3,2,1] Note 2, 1 dont exist all values from diagonal upto 3
 We would have already computed M[5][4] all rows > 4
 these represents the array on left before ) and on the right after (
 We have to decide what will provide better value, removing 5 first, then the array would be
 2*7 + M[5][4] or 2.*5 + M[4][4]


 We start filling the matrix from bottom right [4][4] which looks at elements 5,7,1 index (3, 4, 5)
 Next iteration, we look at [3][3] elements index (2, 3, 4)
 and then followed y [3][4] here we look at (3,3,4)
 */
func MinDP(A []int, M [][]int) int{

    L:= len(A)
    m := -1
    for l:=L-2; l>0; l-- {
        for h:=l; h < L-1; h++ {
            m = M[l][h]
            for i := l ; i <= h; i++ {
                x := M[l][i-1] + M[i+1][h] + A[i] * A[l-1] * A[h+1]
                fmt.Println(l,i,h, x, "M[l][i-1]", M[l][i-1], "M[i+1][h]", M[i+1][h])
                m = Max(m, x)
            }
            M[l][h] = m
            //for i:=0; i < L;i++ {    fmt.Println(M[i])}
            //fmt.Println()
        }
    }
    //fmt.Println(M)
    return M[1][L-2]
}


