/*
Given a n x n matrix where each of the rows and columns are sorted in ascending order, find the kth smallest
element in the matrix.

Note that it is the kth smallest element in the sorted order, not the kth distinct element.

Example:

matrix = [
   [ 1,  5,  9],
   [10, 11, 13],
   [12, 13, 15]
],
k = 8,

return 13.
Note:
You may assume k is always valid, 1 ≤ k ≤ n2.
*/

//Binary Search idea from discussion
/*
 Count the number of elements that are <= val in the array
 */

func BinarySearch(nums []int, target int) int {
    low := 0
    hi := len(nums)
    for low < hi {
        mid := (hi - low) /2 + low
        if nums[mid] > target {
            hi = mid
        } else {
            low = mid + 1
        }

    }
    return low // insert location
}

// O(Rlog(C)) * log(hi)
func kthSmallest(matrix [][]int, k int) int {
    R := len(matrix)
    if R == 0 {return 0}
    C := len(matrix[0])
    lo,hi := matrix[0][0], matrix[R-1][C-1]
    for lo < hi {
        mi := lo + (hi - lo)/2
        //fmt.Println(lo,hi, mi)

        // find number of values less than mi
        c := 0
        for _,row := range(matrix) { // R Log (C)
            c = c + BinarySearch(row, mi)
        }
        if c < k {
            lo = mi + 1
        } else {
            hi = mi
        }
    }
    return lo
}


// My HeapQ solution

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


// O(k log(k)) now k = R*C and this becomes n^2 log n algorithm
func kthSmallest(matrix [][]int, k int) int {
    R := len(matrix)
    if R == 0 {return 0}
    C := len(matrix[0])

    H := &PHeap{}

    heap.Init(H)

    heap.Push(H, P{matrix[0][0],0,0})
    ret := matrix[0][0]
    //push elements from the matrix into the heap k times. Each time pull the value that is minimum
    for k > 0 {
        x := heap.Pop(H).(P)
        //fmt.Println(x)
        ret = x.Val
        if x.row == 0 && x.col < C-1 {
            heap.Push(H, P{matrix[x.row][x.col+1], x.row, x.col+1})
        }
        if x.row < R-1 {
            heap.Push(H, P{matrix[x.row+1][x.col], x.row+1, x.col})
        }
        k--
    }
    return ret
}

/*
 Ther
*/
