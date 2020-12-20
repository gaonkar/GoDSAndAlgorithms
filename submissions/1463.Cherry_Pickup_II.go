/*
Given a rows x cols matrix grid representing a field of cherries. Each cell in grid represents the number of cherries that
you can collect. You have two robots that can collect cherries for you, Robot #1 is located at the top-left corner (0,0),
and Robot #2 is located at the top-right corner (0, cols-1) of the grid.
Return the maximum number of cherries collection using both robots  by following the rules below:

From a cell (i,j), robots can move to cell (i+1, j-1) , (i+1, j) or (i+1, j+1).
When any robot is passing through a cell, It picks it up all cherries, and the cell becomes an empty cell (0).
When both robots stay on the same cell, only one of them takes the cherries.
Both robots cannot move outside of the grid at any moment.
Both robots should reach the bottom row in the grid.
*/

// Just a dictionary to map 3 dimension to 1 dimension
func Map(row, c1, c2 int) int {
    ret := (row + 1) * 10000
    ret += (c1 + 1) * 100
    ret += c2
    return ret
}

func cherryPickup(grid [][]int) int {
    var DFS func(row,c1, c2 int) int
    R, C := len(grid), len(grid[0])
    cache := make(map[int]int)
    DFS = func(row, c1, c2 int) int {
        if c1 < 0 || c2 < 0 || c1 >= C || c2 >= C { return -10000}
        ret, ok := cache[Map(row, c1, c2)]
        if ok {
            return ret
        }

        ret = grid[row][c1]
        if c1 != c2 {
            ret += grid[row][c2]
        }

        if row >= R-1 {return ret}
        next,tmp :=0,0
        for i:=c1-1; i <= c1+1; i++ {
            for j:=c2-1; j <= c2+1; j++ {
                tmp = DFS(row+1, i, j)
                if tmp > next {next = tmp}
            }
        }
        ret += next
        cache[Map(row, c1, c2)] = ret
        return ret
    }
    return DFS(0,0,C-1)
 }
