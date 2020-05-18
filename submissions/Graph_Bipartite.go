/*
Given an undirected graph, return true if and only if it is bipartite.

Recall that a graph is bipartite if we can split it's set of nodes into two independent subsets A and B
such that every edge in the graph has one node in A and another node in B.

The graph is given in the following form: graph[i] is a list of indexes j for which the edge between nodes i and j
exists.  Each node is an integer between 0 and graph.length - 1.
There are no self edges or parallel edges: graph[i] does not contain i, and it doesn't contain any element twice.
*/

isBipartite(graph [][]int) bool {
    var q []int
    var nodes []int
    var n int
    color := make(map[int] int)
    visited := make(map[int] bool)
    marked := 1
    for i, x:= range(graph) {
        if len(x) > 0 {
            nodes = append(nodes, i)
            if (marked > 0) {
                color[i] = marked
                q = append(q, i)
            }
            marked = 0
        }
    }
    for len(q) > 0 || len(nodes) > 0 { //BFS walk
        if len(q) > 0 {
            n = q[0]
            q = q[1:len(q)]
        } else {
            i := len(nodes) - 1
            for i >=0 &&  visited[nodes[i]] {
                i--
            }
            if (i < 0) { break}
            //fmt.Println(n, nodes)
            n = nodes[i]
            nodes = nodes[:i]
            color[n] = -1
            //fmt.Println(n,nodes)
        }
        visited[n] = true
        for _,j := range(graph[n]) {
            if visited[j] == false {
                q = append(q, j)
            }
            c, ok := color[j]
            if ok {
                if c + color[n] != 0 {
                    //fmt.Println(color, n, j, c, visited)
                    return false
                }
            } else {
                color[j] = -color[n]
            }
        }
    }
    fmt.Println(color)
    return true
}
