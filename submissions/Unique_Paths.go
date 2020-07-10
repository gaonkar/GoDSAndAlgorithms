/*
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
The robot can only move either down or right at any point in time.
The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?
*/

func uniquePathsIter2(m int, n int) int {
    arr := make([]int, n)
    for j := 0; j < n; j++ {
        arr[j]=1
    }
    for i := 1; i < m; i++ {
        for j := 1;j < n; j++ {
            arr[j] += arr[j-1]
        }
    }
    return arr[n-1]
}

func uniquePaths(m int, n int) int {
    arr := make([][]int, m)
    arr[0] = make([] int, n)
    //only 1 way to go in the row 0
    for j := 0; j < n; j++ {
        arr[0][j]=1
    }
    for i := 1; i < m; i++ {
        arr[i] = make([]int, n)
        arr[i][0] = 1
        for j := 1;j < n; j++ {
            arr[i][j] = arr[i-1][j] + arr[i][j-1]
        }
    }
    return arr[m-1][n-1]
}

/*
 Unique Paths 2
 An obstacle and empty space is marked as 1 and 0 respectively in the grid.

Note: m and n will be at most 100.

Example 1:

Input:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
Output: 2
Explanation:
There is one obstacle in the middle of the 3x3 grid above.
There are two ways to reach the bottom-right corner:
1. Right -> Right -> Down -> Down
2. Down -> Down -> Right -> Right
*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])
    arr := make([][]int, m)
    arr[0] = make([] int, n)
    for j := 0; j < n; j++ {
        if obstacleGrid[0][j] == 0 && j == 0 || obstacleGrid[0][j] == 0 && j > 0 && arr[0][j-1]==1 {
            arr[0][j]=1
        }
    }
    for i := 1; i < m; i++ {
        arr[i] = make([]int, n)
        if obstacleGrid[i][0] == 0 && i == 0 || obstacleGrid[i][0] == 0 && i > 0 && arr[i-1][0]==1 {
            arr[i][0] = 1
        }
        for j := 1;j < n; j++ {
            if obstacleGrid[i][j] == 0 {
                arr[i][j] = arr[i-1][j] + arr[i][j-1]
            }
        }
    }
    //fmt.Println(arr)
    return arr[m-1][n-1]
}
