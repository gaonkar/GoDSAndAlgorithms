/*
Given a list of airline tickets represented by pairs of departure and arrival airports [from, to], reconstruct the
itinerary in order. All of the tickets belong to a man who departs from JFK. Thus, the itinerary must begin with JFK.

Note:

If there are multiple valid itineraries, you should return the itinerary that has the smallest lexical order when
read as a single string. For example, the itinerary ["JFK", "LGA"] has a smaller lexical order than ["JFK", "LGB"].
All airports are represented by three capital letters (IATA code).
You may assume all tickets form at least one valid itinerary.
One must use all the tickets once and only once.
Example 1:

Input: [["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
Output: ["JFK", "MUC", "LHR", "SFO", "SJC"]
Example 2:

Input: [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
Output: ["JFK","ATL","JFK","SFO","ATL","SFO"]
Explanation: Another possible reconstruction is ["JFK","SFO","ATL","JFK","ATL","SFO"].
             But it is larger in lexical order.

Solution Insight
The constraint states that all ticket entries are valid, but the final destination might not be the source.
Suppose, we had that additional contraint that the person has to be back at the source. Then, this would just a DFS walk
sorted by airport code. Now how to find the open ended ticket. Consider the following path,
JFK->A->B->C->A->D->M->Z->O->M->N
While walking this graph, we end up going JFK->A->B->C->A->M->N, we need to back track back to M until we find a new path.
And thats the solution to this problem. There can be some optimizations done on how we append to make it more optimal.
*/

func Reverse(r []string) {
    n := len(r) - 1
    for i:=0; i < len(r)/2; i++ {
        r[i],r[n] = r[n],r[i]
        n--
    }
}

func findItinerary(tickets [][]string) (r []string) {
    M := make(map[string] []string)
    for _, x:= range(tickets) {
        M[x[0]] = append(M[x[0]], x[1])
    }
    for a,_ := range(M) {
        sort.Strings(M[a])
        //fmt.Println(a, M[a])
    }
    r = append(r, "JFK")
    l := 0
    ret := make([]string, 0)
    for len(r) > 0 {
        //fmt.Println(l, r, ret)
        k := r[l]
        _, ok := M[k]
        if !ok {
            ret = append(ret, r[l])
            r = r[:l]
            l--
            continue
        }
        r = append(r, M[k][0])
        kl := len(M[k])
        if kl == 1 {
            delete(M, k)
        } else {
            M[k] = M[k][1:]
        }
        l++
    }
    //et = append(r, ret...)
    Reverse(ret)
    return ret
}/
