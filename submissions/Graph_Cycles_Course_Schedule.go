/*

207. Course Schedule
Medium

3326

166

Add to List

Share
There are a total of numCourses courses you have to take, labeled from 0 to numCourses-1.

Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: [0,1]

Given the total number of courses and a list of prerequisite pairs, is it possible for you to finish all courses?



Example 1:

Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take.
             To take course 1 you should have finished course 0. So it is possible.
Example 2:

Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take.
             To take course 1 you should have finished course 0, and to take course 0 you should
             also have finished course 1. So it is impossible.


Constraints:

The input prerequisites is a graph represented by a list of edges, not adjacency matrices. Read more about how a graph is represented.
You may assume that there are no duplicate edges in the input prerequisites.
1 <= numCourses <= 10^5

*/
type Graph struct {
    N int
    V []int
    AL []map[int] bool
}

func CreateGraph(Vn int, EL [][]int) *Graph {
    var G Graph
    G.N = Vn
    G.V = make([]int, Vn)
    G.AL = make([]map[int] bool, Vn)
    for c:= 0; c < Vn; c++ {
        G.AL[c] = make(map[int] bool)
    }
    for _,p := range(EL) {
        //fmt.Println(Vn,p, G.AL[p[0]])
        G.AL[p[0]][p[1]] = true
    }
    return &G
}

func (this *Graph) DFSCycleFromVertex(v int, C []int) bool {
    C[v] = 1 // visited
    AL := this.AL[v] // Get the AL
    for n:= range(AL) {
        if C[n] == 1 {
            return true
        }
        if C[n] == 0 && this.DFSCycleFromVertex(n, C) {return true}
    }
    C[v] = 2 //cycle
    return false
}

func (this *Graph) HasCycles() bool {
    color := make([]int, this.N) // 0 not visited, 1 visited, 2 cycle

    for i:= range(color) {
        //fmt.Println(i, color)
        if color[i] == 0 {
            if this.DFSCycleFromVertex(i, color)  {
                return true
            }
        }
    }
    //fmt.Println(color)
    return false
}

func canFinish(N int, pq [][]int) bool {
    G := CreateGraph(N,pq)

    //fmt.Println(G)
    ret := G.HasCycles()
    //fmt.Println(ret)
    return !ret
}
