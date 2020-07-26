import hp "container/heap"
/*
    1. Priority Queue Using Heap
    2. Graph and dijkstra algorithm
    3. Represent cell (i,j) as a node with id = i*C + j,
    4. Map relationship between cell(i,j) to cell(i+1,j) and cell(i,j+1) as diff of heights
    5. Use Dijkstra's algorithm to find the short parth between 0 to R*C - 1
    6. the edge cost is modified from usual dijkstras of total cost
        edgeCost = cost + abs(heights[i][j] - heights[i'][j']) to 
        edgeCost = Max(cost, abs(heights[i][j] - heights[i'][j'])  

// adopted dijkstras from here
https://dev.to/douglasmakey/implementation-of-dijkstra-using-heap-in-go-6e3

*/
type path struct {
    value int
    nodes []int
}

type minPath []path

func (h minPath) Len() int           { return len(h) }
func (h minPath) Less(i, j int) bool { return h[i].value < h[j].value }
func (h minPath) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minPath) Push(x interface{}) {
    *h = append(*h, x.(path))
}

func (h *minPath) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

type heap struct {
    values *minPath
}

func newHeap() *heap {
    return &heap{values: &minPath{}}
}

func (h *heap) push(p path) {
    hp.Push(h.values, p)
}

func (h *heap) pop() path {
    i := hp.Pop(h.values)
    return i.(path)
}

type edge struct {
    node   int
    weight int
}

type graph struct {
    nodes map[int][]edge
}

func newGraph() *graph {
    return &graph{nodes: make(map[int][]edge)}
}

func (g *graph) addEdge(origin, destiny int, weight int) {
    //fmt.Println(origin, destiny, weight)
    g.nodes[origin] = append(g.nodes[origin], edge{node: destiny, weight: weight})
    g.nodes[destiny] = append(g.nodes[destiny], edge{node: origin, weight: weight})
}

func (g *graph) getEdges(node int) []edge {
    return g.nodes[node]
}

func Max(x,y int) int {
    if x < y {return y}
    return x
}

func (g *graph) getPath(origin, destiny int) (int, []int) {
    h := newHeap()
    h.push(path{value: 0, nodes: []int{origin}})
    visited := make(map[int]bool)

    for len(*h.values) > 0 {
        // Find the nearest yet to visit node
        p := h.pop()
        node := p.nodes[len(p.nodes)-1]

        if visited[node] {
            continue
        }

        if node == destiny {
            return p.value, p.nodes
        }

        for _, e := range g.getEdges(node) {
            if !visited[e.node] {
                // We calculate the total spent so far plus the cost and the path of getting here
                h.push(path{value: Max(p.value, e.weight), nodes: append([]int{}, append(p.nodes, e.node)...)})
            }
        }

        visited[node] = true
    }

    return 0, nil
}

func Diff(x, y int) int {
    ret := x - y
    if ret < 0 {return -ret}
    return ret
}

func minimumEffortPath(heights [][]int) int {
    R := len(heights)
    if R == 0 {return 0}
    graph := newGraph()
    C:= len(heights[0])
    for i,_:= range(heights) {
        for j:=0; j < C; j++ {
            src := i * C + j        
            if i < R-1 {
                dst := (i+1)*C + j
                graph.addEdge(src,dst,Diff(heights[i][j], heights[i+1][j]) )
            }
            if j < C-1 {
                dst := i*C + j+1
                graph.addEdge(src,dst,Diff(heights[i][j], heights[i][j+1]) )
            }
        }
    }
    ret,_ := graph.getPath(0,R*C-1)
    //fmt.Println(nodes)
    return ret
}
