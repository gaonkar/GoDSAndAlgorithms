/*
295. Find Median from Data Stream


Median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value. So the median is the mean of the two middle value.

For example,
[2,3,4], the median is 3

[2,3], the median is (2 + 3) / 2 = 2.5

Design a data structure that supports the following two operations:

void addNum(int num) - Add a integer number from the data stream to the data structure.
double findMedian() - Return the median of all elements so far.


Example:

addNum(1)
addNum(2)
findMedian() -> 1.5
addNum(3)
findMedian() -> 2


Follow up:

If all integer numbers from the stream are between 0 and 100, how would you optimize it?
If 99% of all integer numbers from the stream are between 0 and 100, how would you optimize it?
*/


// An MaxHeap is a min-heap of ints.
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


// An MinHeap is a min-heap of ints.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
    minH    *MinHeap    // Right of median
    maxH    *MaxHeap    // Left of median
    ret     float64
    count   int
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
    var M MedianFinder
    M.minH = &MinHeap{}
    M.maxH = &MaxHeap{}
    heap.Init(M.minH)
    heap.Init(M.maxH)
    M.count = 0

    return M
}

/*
    [MIN] >= median >= [MAX] are the heaps
    when len(MIN) == len(MAX), then median is ROOT(MIN)
 */
func (this *MedianFinder) AddNum(num int)  {
    if this.minH.Len() == this.maxH.Len() {
        heap.Push(this.maxH, num);  // push into max
        num = (*this.maxH)[0]       // record its highest value
        heap.Pop(this.maxH)         // pop that
        heap.Push(this.minH, num);  // push into min
    } else {
        heap.Push(this.minH, num);  // push into min
        num = (*this.minH)[0]       // record the lowest
        heap.Pop(this.minH)         // pop the value
        heap.Push(this.maxH, num);  // push into maxheap
    }
}


func (this *MedianFinder) FindMedian() float64 {
    if this.minH.Len() > this.maxH.Len() {
        return float64((*this.minH)[0])
    }
    ret := (*this.maxH)[0] + (*this.minH)[0]
    return float64(ret)/2
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
