/*
In a directed graph, we start at some node and every turn, walk along a directed edge of the graph.
If we reach a node that is terminal (that is, it has no outgoing directed edges), we stop.
Now, say our starting node is eventually safe if and only if we must eventually walk to a terminal node.
More specifically, there exists a natural number K so that for any choice of where to walk, we must have stopped at a
terminal node in less than K steps.

Which nodes are eventually safe?  Return them as an array in sorted order.
The directed graph has N nodes with labels 0, 1, ..., N-1, where N is the length of graph.  The graph is given in the
following form: graph[i] is a list of labels j such that (i, j) is a directed edge of the graph.

Example:
Input: graph = [[1,2],[2,3],[5],[0],[5],[],[]]
Output: [2,4,5,6]
Note:

graph will have length at most 10000.
The number of edges in the graph will not exceed 32000.
Each graph[i] will be a sorted list of different integers, chosen within the range [0, graph.length - 1].
*/

/*
    All nodes leading to cycle are unsafe. Else they are safe. 
    Color
    0   - not visited
    1   - visited
    2   - unsafe
*/

func eventualSafeNodes(graph [][]int) (ret []int) {
    c :=make([]int, len(graph))
    var Safe func(n int) bool
    Safe = func(n int) bool {
        if c[n] == 2 {return false}
        if c[n] == 1 {return true}
        c[n] = 2 // mark as unsafe, start exploring
        for _, x := range(graph[n]) {
            if Safe(x) == false {
                return false
            }
        }
        c[n] = 1 // all paths are safe, so mark this node safe
        return true
    }
    for i:=0; i < len(graph); i++ {
        if Safe(i) {ret = append(ret, i)}
    }
    return ret
}
