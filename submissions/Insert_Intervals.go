/*
57. Insert Interval
Hard

1452

172

Add to List

Share
Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).

You may assume that the intervals were initially sorted according to their start times.

Example 1:

Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]
Example 2:

Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.

*/

func Min(x int, y int) int {
    if x < y {return x}
    return y
}
func Max(x int, y int) int {
    if x < y {return y}
    return x
}

func insert(I [][]int, N []int) [][]int {
    var r [][]int
    i := 0
    L := len(I)
    l := N[0]
    h := N[1]
    //add all intervals whose high range is greater than l
    for i < L  &&  l > I[i][1]{
        r = append(r, I[i])
        i++
    }
    // we have now an entry that is intersecting the l, keep merging them
    // until we have h < lower range of the series
    for i < L && h >= I[i][0] {
        l = Min(l, I[i][0])
        h = Max(h, I[i][1])
        i++
    }
    r = append(r, []int{l,h})
    for i < L {
        r = append(r, I[i])
        i++
    }
    fmt.Println(i)
    return r
}
