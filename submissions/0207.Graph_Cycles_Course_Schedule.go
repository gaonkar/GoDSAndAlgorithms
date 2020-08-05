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
	N  int
	D  []int
	AL [][]int
}

func CreateGraph(Vn int, EL [][]int) *Graph {
	var G Graph
	G.N = Vn
	G.D = make([]int, Vn)
	G.AL = make([][]int, Vn)
	for c := 0; c < Vn; c++ {
		G.AL[c] = make([]int, 0)
	}
	for _, p := range EL {
		G.AL[p[1]] = append(G.AL[p[1]], p[0])
		G.D[p[0]]++
	}
	return &G
}

func (this *Graph) DFSCycleFromVertex(v int, C []int) bool {
	C[v] = 1 // visited
	AL := this.AL[v]
	for _, n := range AL {
		if C[n] == 1 {
			return true
		}
		if C[n] == 0 && this.DFSCycleFromVertex(n, C) {
			return true
		}
	}
	C[v] = 2 //cycle
	return false
}

func (this *Graph) HasCycles() bool {
	color := make([]int, this.N) // 0 not visited, 1 visited, 2 cycle

	for i := range color {
		//fmt.Println(i, color)
		if color[i] == 0 {
			if this.DFSCycleFromVertex(i, color) {
				return true
			}
		}
	}
	//fmt.Println(color)
	return false
}

func canFinish(N int, pq [][]int) bool {
	G := CreateGraph(N, pq)

	//fmt.Println(G)
	ret := G.HasCycles()
	//fmt.Println(ret)
	return !ret
}

/*
210 Course Schedule 2
*/

func findOrder(numCourses int, prerequisites [][]int) (r []int) {
	G := CreateGraph(numCourses, prerequisites)
	//fmt.Println(G)
	if G.HasCycles() {
		return nil
	}
	V := make(map[int]bool)
	// lets add all 0 dependent courses to our list

	for true {
		for i, x := range G.D {
			if V[i] || x > 0 {
				continue
			}
			V[i] = true
			r = append(r, i)
			for _, y := range G.AL[i] {
				G.D[y]--
			}
		}
		if len(V) == G.N {
			for i, x := range G.D {
				if !V[i] && x == 0 {
					r = append(r, i)
				}
			}
			break
		}
	}

	return r
}
