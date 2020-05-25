/*
Given two lists of closed intervals, each list of intervals is pairwise disjoint and in sorted order.

Return the intersection of these two interval lists.

(Formally, a closed interval [a, b] (with a <= b) denotes the set of real numbers x with a <= x <= b.
The intersection of two closed intervals is a set of real numbers that is either empty, or can be represented
as a closed interval.  For example, the intersection of [1, 3] and [2, 4] is [2, 3].)
Note:

0 <= A.length < 1000
0 <= B.length < 1000
0 <= A[i].start, A[i].end, B[i].start, B[i].end < 10^9


*/

Max(x int, y int) int {
    if x > y {return x}
    return y
}
func Min(x int, y int) int {
    if x < y {return x}
    return y
}



func intervalIntersection(A [][]int, B [][]int) [][]int {
    ret := make([][]int, 0)
    i := 0
    j := 0
    for i < len(A) && j < len(B) {
        lo := Max(A[i][0],B[j][0])
        hi := Min(A[i][1],B[j][1])
        //mt.Println(A[i],B[j], i,j, lo, hi)
        if lo <= hi {
            r := make([]int, 2)
            r[0] = lo
            r[1] = hi
            ret = append(ret, r)
            //fmt.Println(r)
        }
        if A[i][1] < B[j][1] {
            i++
        } else {
            j++
        }
    }
    return ret
}
