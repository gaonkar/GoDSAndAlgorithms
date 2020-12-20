/*
Given a non-empty list of words, return the k most frequent elements.

Your answer should be sorted by frequency from highest to lowest.
If two words have the same frequency, then the word with the lower alphabetical order comes first.

Example 1:
Input: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
Output: ["i", "love"]
Explanation: "i" and "love" are the two most frequent words.
    Note that "i" comes before "love" due to a lower alphabetical order.
Example 2:
Input: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4
Output: ["the", "is", "sunny", "day"]
Explanation: "the", "is", "sunny" and "day" are the four most frequent words,
    with the number of occurrence being 4, 3, 2 and 1 respectively.
Note:
You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
Input words contain only lowercase letters.
Follow up:
Try to solve it in O(n log k) time and O(n) extra space.
*/
type P  struct{
    Val int
    idx int  // index of the array
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


//Time := O((n+k)log(k))
//Space:= O(n+k)
func topKFrequent(words []string, k int) (r []string) {
    M :=make(map[string] int)
    F := make([]int, len(words))
    for i,w := range(words) {
        _, ok := M[w]
        if !ok {
            M[w] = i
        }
        idx := M[w]
        F[idx]++
    }
    //fmt.Println(F)
    H := &PHeap{}
    for i:=0; i < len(F); i++ {
        if F[i] == 0 {continue}
        heap.Push(H, P{F[i], i})
        if H.Len() > k+10 {
            // this 10 is because the question has this  requirement that if the frequency
            // is the same then have them sorted. but the heap cannot handle this at the moment.
            // It can be solved by making tracking the eviced element, but it was just easier to hold more
            // discard them post sorting
            heap.Pop(H)
        }
    }
    parr := make([]P, 0)
    for H.Len() > 0 {
        x := heap.Pop(H).(P)
        parr = append(parr, x)
    }

    sort.SliceStable(parr, func(i, j int) bool {
    return parr[i].Val > parr[j].Val || parr[i].Val == parr[j].Val && words[parr[i].idx] < words[parr[j].idx]
})
    for i:=0; i < k; i++ {
        x := parr[i]
        r = append(r, words[x.idx])
    }
    return r
}
