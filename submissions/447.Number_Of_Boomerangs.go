/*
Given n points in the plane that are all pairwise distinct, a "boomerang" is a tuple of points (i, j, k) such that the distance between i and j equals the distance between i and k (the order of the tuple matters).

Find the number of boomerangs. You may assume that n will be at most 500 and coordinates of points are all in the range [-10000, 10000] (inclusive).

Example:

Input:
[[0,0],[1,0],[2,0]]

Output:
2

Explanation:
The two boomerangs are [[1,0],[0,0],[2,0]] and [[1,0],[2,0],[0,0]]
*/

func Distance (x, y []int) int {
    a := x[0]-y[0]
    b := x[1]-y[1]
    return (a * a + b * b)
}
func numberOfBoomerangs(points [][]int) (r int) {
    for i:=0; i < len(points); i++ {
        M := make(map[int] int)
        for j:=0; j < len(points); j++ {
            if i == j {continue}
            d := Distance(points[i], points[j])
            M[d]++ //find equidistant points from i and count them
        }
        for _,x := range(M) {
            r += x*(x-1) // NC2 ways
        }
    }
    return r
}
