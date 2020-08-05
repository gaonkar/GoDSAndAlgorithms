/*
You are given two integer arrays nums1 and nums2 sorted in ascending order and an integer k.

Define a pair (u,v) which consists of one element from the first array and one element from the second array.

Find the k pairs (u1,v1),(u2,v2) ...(uk,vk) with the smallest sums.

Example 1:

Input: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
Output: [[1,2],[1,4],[1,6]]
Explanation: The first 3 pairs are returned from the sequence:
             [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
Example 2:

Input: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
Output: [1,1],[1,1]
Explanation: The first 2 pairs are returned from the sequence:
             [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
Example 3:

Input: nums1 = [1,2], nums2 = [3], k = 3
Output: [1,3],[2,3]
Explanation: All possible pairs are returned from the sequence: [1,3],[2,3]
*/



type P  struct{
    Val int
    row int
    col int
}
type PHeap []P

func (h PHeap) Len() int           { return len(h) }
func (h PHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h PHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PHeap) Push(x interface{}) {
        // Push and Pop use pointer receivers because they modify the slice's length,
        // not just its contents.
        *h = append(*h, x.(P))
}

func (h *PHeap) Pop() interface{} {
        old := *h
        n := len(old)
        x := old[n-1]
        *h = old[0 : n-1]
        return x
}


func GetP(nums1, nums2 []int, i,j int) P {
    var p P
    p.row = i
    p.col = j
    p.Val = nums1[i] + nums2[j]
    return p
}

/*
 Same as Kth Smallest Element in the Matrix

*/
func kSmallestPairs(nums1 []int, nums2 []int, k int) (r [][]int) {
    R := len(nums1)
    if R == 0 {return nil}
    C := len(nums2)
    if C == 0 {return nil}

    H := &PHeap{}

    heap.Init(H)

    heap.Push(H, GetP(nums1,nums2, 0, 0))
    if k > R*C {
        k = R *C
    }
    for k > 0 {
        x := heap.Pop(H).(P)
        r = append(r, []int{nums1[x.row], nums2[x.col]})
        if x.row == 0 && x.col < C-1 {
            heap.Push(H, GetP(nums1, nums2, x.row, x.col+1))
        }
        if x.row < R-1 {
            heap.Push(H, GetP(nums1, nums2, x.row+1, x.col))
        }
        k--
    }
    return r
}
