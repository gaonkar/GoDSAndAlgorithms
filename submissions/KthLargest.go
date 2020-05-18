/*
Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:

Input: [3,2,1,5,6,4] and k = 2
Output: 5
Example 2:

Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
Note:
You may assume k is always valid, 1 ≤ k ≤ array's length.


*/

// sol 1 heap
import (
	"container/heap"
	"fmt"
)
// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func findKthLargest(nums []int, k int) int {
    h := &IntHeap{nums[0]}
	heap.Init(h)
    for _,x:= range(nums[1:]) {
        heap.Push(h, x)
        //fmt.Println(h)
    }
    for k > 1 {
        //fmt.Printf("%d ", heap.Pop(h))
        k--
    }
    return heap.Pop(h).(int)
}

// sol 2 quick select
//standard Lomuto partition function
//standard Lomuto partition function
func Partition(nums []int) int {
    L := len(nums)
    p := nums[L-1]
    j :=0
    for i:=0;i< len(nums)-1; i++{
        if nums[i] < p {
            nums[i], nums[j] = nums[j], nums[i]
            j++
        }
    }
    nums[j], nums[L-1] = nums[L-1], nums[j]
    return j
}


func findKthLargest(nums []int, k int) int {
    L := len(nums)
    P := Partition(nums)
    K := L - P
    if (K == k) {
        return nums[P]
    } else if K > k {
        return findKthLargest(nums[P+1:], k)
    }
    return findKthLargest(nums[:P], k-K)
}
