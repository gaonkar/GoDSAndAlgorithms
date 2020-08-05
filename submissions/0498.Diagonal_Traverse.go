/*
Given a matrix of M x N elements (M rows, N columns), return all elements of the matrix in diagonal order as shown in the below image.

Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]

Output:  [1,2,4,7,5,3,6,8,9]
*/

func Max(x,y int) int {
    if x < y {return y}
    return x
}
func findDiagonalOrder(matrix [][]int) (r []int) {
    R := len(matrix)
    if R == 0 {return r}
    C := len(matrix[0])
    i, j, sum,k := 0,0, 0,0
    flag := -1  // flag allows us to switch going NW and SE direction
    for sum <= R+C-2 {
        k = 0
        if flag > 0 {
            i, j = 0, sum
            k = Max(k, sum - C+1)  // key here, after sum becomes greater than C-1, we can eliminate
                                   // phantom cells
        } else {
            i, j = sum, 0
            k = Max(k, sum - R+1)
        }
        for k <= sum {
            x, y := i+k*flag, j-k*flag

            if x < R && y < C {
                r = append(r, matrix[x][y])
            } else {
                //fmt.Println(x,y,k, sum)
                // same thing here, but we just break out after first phantom cell 
                break
            }
            k++
        }
        flag = -flag
        sum++
        //fmt.Println(r, sum)
    }
    return r
}
