/*
Given an array of events where events[i] = [startDayi, endDayi]. Every event i starts at startDay i and ends at endDay i.
You can attend an event i at any day d where startTimei <= d <= endTimei. Notice that you can only attend one event at any time d.
Return the maximum number of events you can attend.

Example 1:


Input: events = [[1,2],[2,3],[3,4]]
Output: 3
Explanation: You can attend all the three events.
One way to attend them all is as shown.
Attend the first event on day 1.
Attend the second event on day 2.
Attend the third event on day 3.
Example 2:

Input: events= [[1,2],[2,3],[3,4],[1,2]]
Output: 4
Example 3:

Input: events = [[1,4],[4,4],[2,2],[3,4],[1,1]]
Output: 4
Example 4:

Input: events = [[1,100000]]
Output: 1
Example 5:

Input: events = [[1,1],[1,2],[1,3],[1,4],[1,5],[1,6],[1,7]]
Output: 7


Constraints:

1 <= events.length <= 105
events[i].length == 2
1 <= startDayi <= endDayi <= 105

*/


type P  struct{
    s int
    e int
}
type PHeap []P

func (h PHeap) Len() int           { return len(h) }
func (h PHeap) Less(i, j int) bool { return h[i].e < h[j].e }
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


/*
 Solution
    Use a priority heap to track the end day after the list is sorted by start days.
    Initialize d to be I[i][0] if heap is 0. You could have disjoint clusters of overlapping events.
    Add all events that have same or less start date than current d.
    Pull events out from earliest ending date.
*/


func maxEvents(I [][]int) int {
    sort.SliceStable(I, func(i, j int) bool { return I[i][0] < I[j][0] || (I[i][0] == I[j][0] && I[i][1] < I[j][1])})
    H := &PHeap{}
    count:=0
    heap.Init(H)
    fmt.Println(I)

    i, d, count :=0, 0, 0
    for i < len(I) ||  H.Len() > 0 {
        if H.Len()==0 {
            d = I[i][0]
        }
        if i < len(I) && I[i][0] <= d {
            heap.Push(H, P{I[i][0], I[i][1]})
            i++
            continue
        }
        x := heap.Pop(H).(P)
        count++
        d++
        for H.Len() > 0 {
            x = heap.Pop(H).(P)
            if x.e >= d {
                heap.Push(H,x)
                break
            }
        }
    }

    return count
}
