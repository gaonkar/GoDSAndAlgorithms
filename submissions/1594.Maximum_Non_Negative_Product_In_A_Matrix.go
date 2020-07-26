/*
You are given a rows x cols matrix grid. Initially, you are located at the top-left corner (0, 0), and in each step, you can only move right or down in the matrix.

Among all possible paths starting from the top-left corner (0, 0) and ending in the bottom-right corner
(rows - 1, cols - 1), find the path with the maximum non-negative product. The product of a path is the product of all
integers in the grid cells visited along the path.
Return the maximum non-negative product modulo 109 + 7. If the maximum product is negative return -1.
Notice that the modulo is performed after getting the maximum product.


Example 1:

Input: grid = [[-1,-2,-3],
               [-2,-3,-3],
               [-3,-3,-2]]
Output: -1
Explanation: It's not possible to get non-negative product in the path from (0, 0) to (2, 2), so return -1.
Example 2:

Input: grid = [[1,-2,1],
               [1,-2,1],
               [3,-4,1]]
Output: 8
Explanation: Maximum non-negative product is in bold (1 * 1 * -2 * -4 * 1 = 8).
Example 3:

Input: grid = [[1, 3],
               [0,-4]]
Output: 0
Explanation: Maximum non-negative product is in bold (1 * 0 * -4 = 0).
Example 4:

Input: grid = [[ 1, 4,4,0],
               [-2, 0,0,1],
               [ 1,-1,1,1]]
Output: 2
Explanation: Maximum non-negative product is in bold (1 * -2 * 1 * -1 * 1 * 1 = 2).


Constraints:

1 <= rows, cols <= 15
-4 <= grid[i][j] <= 4
*/

func Make2D(r, c, v int) [][]int {
	ret := make([][]int, r)
	for i := 0; i < len(ret); i++ {
		ret[i] = make([]int, c)
		if v > 0 {
			for j := 0; j < c; j++ {
				ret[i][j] = v
			}
		}
	}
	return ret
}

func Max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
func Min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func Ret(x int) int {
    if x < 0 {return -1}
    return x
}
/*
    You need to keep 2 arrays.
    HP-> MaxProduct
    LP ->MinProduct
    if [i,j] < 0, then get min value [i-1][j] and [i][j-1], else get the max values
    same logic for LP 

*/
func maxProductPath(grid [][]int) int {
    R := len(grid)
    if R == 0 {return 0}
    C := len(grid[0])
    HP := Make2D(R,C,0)
    LP := Make2D(R,C,0)
    MOD := 1000000007

    HP[0][0] =  grid[0][0]
    LP[0][0] =  grid[0][0]

    //fill the top row
    for i:=1;i<R; i++{
        HP[i][0] = grid[i][0] * HP[i-1][0] % MOD
        LP[i][0] = HP[i][0]
    }
    for j:=1;j<C; j++{
        HP[0][j] = grid[0][j] * HP[0][j-1] % MOD
        LP[0][j] = HP[0][j]
    }
    for i:=1;i<R; i++ {
        for j:=1;j<C;j++ {
            if grid[i][j] < 0 {
                HP[i][j] = Min(LP[i-1][j], LP[i][j-1]) * grid[i][j]
                LP[i][j] = Max(HP[i-1][j], HP[i][j-1]) * grid[i][j]
            } else {
                LP[i][j] = Min(LP[i-1][j], LP[i][j-1]) * grid[i][j]
                HP[i][j] = Max(HP[i-1][j], HP[i][j-1]) * grid[i][j]
            }
        }
    }
    //fmt.Println(HP)
    return Ret(HP[R-1][C-1]) % MOD
}
