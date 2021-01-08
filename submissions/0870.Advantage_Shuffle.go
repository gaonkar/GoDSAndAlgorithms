/*
Given two arrays A and B of equal size, the advantage of A with respect to B is the number of indices i for which A[i] > B[i].

Return any permutation of A that maximizes its advantage with respect to B.



Example 1:

Input: A = [2,7,11,15], B = [1,10,4,11]
Output: [2,11,7,15]
Example 2:

Input: A = [12,24,8,32], B = [13,25,32,11]
Output: [24,32,8,12]


Note:

1 <= A.length = B.length <= 10000
0 <= A[i] <= 10^9
0 <= B[i] <= 10^9
*/
type BIDX struct {
    i,v int
}
// O(N log N) time and O(N) space
func advantageCount(A []int, B []int) []int {
    sort.Ints(A)
    barr := make([]BIDX, 0)
    for i, v := range(B) {
        barr = append(barr, BIDX{i,v})
    }
    sort.SliceStable(barr, func(i, j int) bool {
    return barr[i].v > barr[j].v
    })

    alen := len(A)-1
    ret := make([]int, alen+1)
    for i,_ := range(barr) {
        if barr[i].v < A[alen] {
            // assign A to get advantage at this index
            ret[barr[i].i] = A[alen]
            alen--
            continue
        }
        // mark the entry can take any value
        ret[barr[i].i] = -1
    }
    for i,_ := range(barr) {
        if ret[barr[i].i] < 0 {
            //fill any value
            ret[barr[i].i]  = A[alen]
            alen--
            continue
        }
    }
    return ret
}
