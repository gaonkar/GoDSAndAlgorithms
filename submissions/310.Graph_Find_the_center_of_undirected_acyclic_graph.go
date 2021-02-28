/*
For an undirected graph with tree characteristics, we can choose any node as the root.
The result graph is then a rooted tree. Among all possible rooted trees, those with minimum height are called minimum
height trees (MHTs). Given such a graph, write a function to find all the MHTs and return a list of their root labels.

Format
The graph contains n nodes which are labeled from 0 to n - 1. You will be given the number n and a list of
 undirected edges (each edge is a pair of labels).

You can assume that no duplicate edges will appear in edges.
Since all edges are undirected, [0, 1] is the same as [1, 0] and thus will not appear together in edges.

Example 1 :

Input: n = 4, edges = [[1, 0], [1, 2], [1, 3]]

        0
        |
        1
       / \
      2   3

Output: [1]
Example 2 :

Input: n = 6, edges = [[0, 3], [1, 3], [2, 3], [4, 3], [5, 4]]

     0  1  2
      \ | /
        3
        |
        4
        |
        5

Output: [3, 4]
Note:

According to the definition of tree on Wikipedia: “a tree is an undirected graph in which any two vertices are
connected by exactly one path. In other words, any connected graph without simple cycles is a tree.”
The height of a rooted tree is the number of edges on the longest downward path between the root and a leaf.
*/

type Graph struct {
	N  int
	AL [][]int
}

func CreateUndirectedGraph(Vn int, EL [][]int) *Graph {
	var G Graph
	G.N = Vn
	G.AL = make([][]int, Vn)
	for c := 0; c < Vn; c++ {
		G.AL[c] = make([]int, 0)
	}
	for _, p := range EL {
		G.AL[p[1]] = append(G.AL[p[1]], p[0])
        G.AL[p[0]] = append(G.AL[p[0]], p[1])
	}
	return &G
}

func (this *Graph)LPath(idx int) (r []int) {
    V := make([]bool, this.N) // visited
    D := make([]int, this.N) // distance
    P := make([]int, this.N) // parent node
    Q := make([]int,1)
    Q[0]=idx

    for len(Q) > 0 {
        n:= Q[0]
        Q=Q[1:]
        if V[n] {continue}
        V[n] = true
        for _,x := range(this.AL[n]) {
            if !V[x] {
                Q = append(Q, x)
                P[x] = n
                D[x] = D[n]+1
            }
        }
    }
    // find the farthest node
    max := 0
    pos := 0
    for i:=0; i < len(D); i++ {
        if max < D[i] {
            max = D[i]
            pos = i
        }
    }
    // find the path
    for pos != idx {
        r = append(r, pos)
        pos = P[pos]
    }
    r = append(r, idx)
    //fmt.Println(r)
    return r
}

/*
 Find the longest path from random nodes.
 Take its end. Find the longest path from this.
 Now find the midpoint
 */
func findMinHeightTrees(n int, edges [][]int) []int {
    G := CreateUndirectedGraph(n, edges)
    path0:= G.LPath(0)
    pathN:= G.LPath(path0[0])
    x := len(pathN)/2
    if len(pathN) % 2 == 0 {
        return pathN[x-1:x+1]
    }
    return pathN[x:x+1]
}
