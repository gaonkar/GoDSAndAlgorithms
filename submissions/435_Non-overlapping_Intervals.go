/*
Given a collection of intervals, find the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.



Example 1:

Input: [[1,2],[2,3],[3,4],[1,3]]
Output: 1
Explanation: [1,3] can be removed and the rest of intervals are non-overlapping.
Example 2:

Input: [[1,2],[1,2],[1,2]]
Output: 2
Explanation: You need to remove two [1,2] to make the rest of intervals non-overlapping.
Example 3:

Input: [[1,2],[2,3]]
Output: 0
Explanation: You don't need to remove any of the intervals since they're already non-overlapping.


Note:

You may assume the interval's end point is always bigger than its start point.
Intervals like [1,2] and [2,3] have borders "touching" but they don't overlap each other.

*/
func eraseOverlapIntervals(I [][]int) (r int) {
    sort.SliceStable(I, func(i, j int) bool {
    return I[i][1] < I[j][1]})
    p := I[0]
    r = 0
    fmt.Println(I)
    for i:=1 ; i < len(I); i++ {
        //fmt.Println(p, I[i])
        if p[1] <= I[i][0] {
            p = I[i]
        } else {
            r++
            if p[1] > I[i][1] {
                p = I[i]
            }
        }
    }
    return r
}
