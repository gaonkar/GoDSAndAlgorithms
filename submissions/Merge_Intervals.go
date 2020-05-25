/*
56. Merge Intervals
Medium

3765

270

Add to List

Share
Given a collection of intervals, merge all overlapping intervals.

Example 1:

Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
*/

func merge(I [][]int) [][]int {
    var ret [][]int
    sort.SliceStable(I, func(i, j int) bool {
    return I[i][0] < I[j][0]
})
    fmt.Println(I)
    for i:= 0 ; i < len(I); i++ {
        j:=i+1
        U := I[i][1]            // current max interval
        for j < len(I) {
            if U >= I[j][1] {
                j++
                continue
            }
            if I[j][0] <= U && U <= I[j][1] {
                U = I[j][1]     // new max interval
                j++
                continue
            }
            break
        }
        j--
        x := []int{I[i][0], U}
        ret = append(ret,x)
        //fmt.Println(i, j, I[i][0], I[j][1])
        i=j

    }
    return ret
}
