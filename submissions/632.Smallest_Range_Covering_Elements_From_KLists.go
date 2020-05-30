/*
You have k lists of sorted integers in ascending order. Find the smallest range that includes at least one number from each of the k lists.

We define the range [a,b] is smaller than range [c,d] if b-a < d-c or a < c if b-a == d-c.



Example 1:

Input: [[4,10,15,24,26], [0,9,12,20], [5,18,22,30]]
Output: [20,24]
Explanation:
List 1: [4, 10, 15, 24,26], 24 is in range [20,24].
List 2: [0, 9, 12, 20], 20 is in range [20,24].
List 3: [5, 18, 22, 30], 22 is in range [20,24].


Note:

The given list may contain duplicates, so ascending order means >= here.
1 <= k <= 3500
-105 <= value of elements <= 105.

Add nums[i][0] into K min heap, track row and col of the value, remove the lowest and replace with next element from the same row.
Track max as you add into the list, stop when one col reaches K list

*/


// An IntHeap is a min-heap of ints.
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


func smallestRange(nums [][]int) []int {
    KN := len(nums)
    H := &PHeap{}
    heap.Init(H)


    max := nums[0][0]
    for i:= 0; i < KN; i++ {
        x := P{nums[i][0], i, 0}
        heap.Push(H,x)
        if max < nums[i][0] {
            max = nums[i][0]
        }
    }
    RMAX:=max
    RMIN:=-1000000
    min := 0
    for true {

        x:= heap.Pop(H).(P)
        min = x.Val
        if RMIN == -1000000 {
            RMIN = min
        } else if max - min < RMAX - RMIN {
            RMAX = max
            RMIN = min
        }
        i := x.row
        j := x.col
        //fmt.Println(x, nums[i][j])
        if j == len(nums[i]) -1 {break}
        j++
        x = P{nums[i][j], i, j}
        if max < x.Val {max = x.Val}
        heap.Push(H,x)
    }

    return []int{RMIN, RMAX}
}
