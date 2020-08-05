/*
347. Top K Frequent Elements
Medium

Share
Given a non-empty array of integers, return the k most frequent elements.

Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
Example 2:

Input: nums = [1], k = 1
Output: [1]
Note:

You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
It's guaranteed that the answer is unique, in other words the set of the top k frequent elements is unique.
You can return the answer in any order.

*/

type Freq struct {
    Val int
    count int
}

func topKFrequent(nums []int, k int) []int {
    M := make(map[int] int)
    F := []Freq{}
    for _,n := range(nums) {
        V, ok := M[n]
        if ok {
            F[V].count++
        } else {
            V := Freq{n, 1}
            F = append(F, V)
            M[n] = len(F) - 1
        }
    }
    sort.SliceStable(F, func(i, j int) bool {
        return F[i].count > F[j].count
    })
    r := []int{}
    for i:=0; i < k; i++ {
        r = append(r, F[i].Val)
    }
    return r
}
