/*
Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted in ascending from left to right.
Integers in each column are sorted in ascending from top to bottom.
Example:

Consider the following matrix:

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
Given target = 5, return true.

Given target = 20, return false.
 if you look at any rectangle, top left is the lowest value and bottom right is the highest. so you just need to box it in
 O(M+N)
 A lot of time has been spent on this problem
 https://stackoverflow.com/questions/2457792/how-do-i-search-for-a-number-in-a-2d-array-sorted-left-to-right-and-top-to-botto/2458113#2458113
Its also called saddleback search
*/
func searchMatrix(matrix [][]int, target int) bool {
	R := len(matrix)
	if R == 0 {
		return false
	}
	C := len(matrix[0])
	i := 0
	j := C - 1
	for i < R && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}
