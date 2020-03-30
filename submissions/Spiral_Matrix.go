/*
Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

Example 1:

Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
Output: [1,2,3,6,9,8,7,4,5]

*/
func Debug(DEBUG bool, a ...interface{}) {
    if DEBUG {
        fmt.Println(a)
    }
}

func spiralOrder(matrix [][]int) []int {
    var ret []int
    R:= len(matrix)
    if R == 0 {
        return ret
    }
    C:= len(matrix[0])
    debug:= false
    M := R * C

    for s:=0; s <= R/2+1; s++ { // each spiral
        // right
        i:=s
        j:=s
        Debug(debug, s, R, C, ret)
        for j < C { //right
            ret = append(ret,matrix[i][j])
            j++
            if len(ret) == M {
                return ret
            }
        }
        j--
        i++
        Debug(debug, "R", i,j, ret)
        // down
        for  i < R {
            ret = append(ret,matrix[i][j])
            i++
            if len(ret) == M {
                return ret
            }
        }
        i--
        j--
        Debug(debug, "D", i,j, ret)
        // left
        for j >= s {
            ret = append(ret,matrix[i][j])
            j--
            if len(ret) == M {
                return ret
            }
        }
        j++
        i--
        Debug(debug, "L", i,j, ret)
        // up
        for i > s {
            ret = append(ret, matrix[i][j])
            i--
            if len(ret) == M {
                return ret
            }
        }
        Debug(debug, "-", i,j, ret)
        R--
        C--
    }
    return ret
}
