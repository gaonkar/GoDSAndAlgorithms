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

*/import "fmt"

func Debug(DEBUG bool, a ...interface{}) {
	if DEBUG {
		fmt.Println(a)
	}
}

func spiralOrder(matrix [][]int) []int {
	var ret []int
	R := len(matrix)
	if R == 0 {
		return ret
	}
	C := len(matrix[0])
	debug := false
	M := R * C

	for s := 0; s <= R/2+1; s++ { // each spiral
		// right
		i := s
		j := s
		Debug(debug, s, R, C, ret)
		for j < C { //right
			ret = append(ret, matrix[i][j])
			j++
			if len(ret) == M {
				return ret
			}
		}
		j--
		i++
		Debug(debug, "R", i, j, ret)
		// down
		for i < R {
			ret = append(ret, matrix[i][j])
			i++
			if len(ret) == M {
				return ret
			}
		}
		i--
		j--
		Debug(debug, "D", i, j, ret)
		// left
		for j >= s {
			ret = append(ret, matrix[i][j])
			j--
			if len(ret) == M {
				return ret
			}
		}
		j++
		i--
		Debug(debug, "L", i, j, ret)
		// up
		for i > s {
			ret = append(ret, matrix[i][j])
			i--
			if len(ret) == M {
				return ret
			}
		}
		Debug(debug, "-", i, j, ret)
		R--
		C--
	}
	return ret
}

/*
Spiral Matrix III : 885
On a 2 dimensional grid with R rows and C columns, we start at (r0, c0) facing east.

Here, the north-west corner of the grid is at the first row and column, and the south-east corner of the grid is at the last row and column.

Now, we walk in a clockwise spiral shape to visit every position in this grid.

Whenever we would move outside the boundary of the grid, we continue our walk outside the grid (but may return to the grid boundary later.)

Eventually, we reach all R * C spaces of the grid.

Return a list of coordinates representing the positions of the grid in the order they were visited.
Input: R = 5, C = 6, r0 = 1, c0 = 4
Output: [[1,4],[1,5],[2,5],[2,4],[2,3],[1,3],[0,3],[0,4],[0,5],[3,5],[3,4],[3,3],[3,2],[2,2],[1,2],[0,2],[4,5],[4,4],[4,3],[4,2],[4,1],[3,1],[2,1],[1,1],[0,1],[4,0],[3,0],[2,0],[1,0],[0,0]]


*/

func spiralMatrixIII(R int, C int, r0 int, c0 int) (r [][]int) {
	moves := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} //directions
	index, steps := 0, 0
	r = append(r, []int{r0, c0})
	for len(r) < R*C {
		steps = index/2 + 1
		// index is generating numbers 0,1,2,3,4,5,6,...
		// steps takes the values      1,1,2,2,3,3 and that is what we need
		// moves[index%4] will choose east/south/west/north for each step
		for i := 0; i < steps; i++ {
			r0 += moves[index%4][0]
			c0 += moves[index%4][1]
			if 0 <= r0 && r0 < R && 0 <= c0 && c0 < C {
				r = append(r, []int{r0, c0})
			}
		}
		index++
	}
	return r
}
