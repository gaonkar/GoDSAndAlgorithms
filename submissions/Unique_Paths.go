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
