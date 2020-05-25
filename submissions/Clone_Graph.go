/*
133. Clone Graph
Medium

1563

1300

Add to List

Share
Given a reference of a node in a connected undirected graph.

Return a deep copy (clone) of the graph.

Each node in the graph contains a val (int) and a list (List[Node]) of its neighbors.

class Node {
    public int val;
    public List<Node> neighbors;
}


Test case format:

For simplicity sake, each node's value is the same as the node's index (1-indexed). For example, the first node with val = 1, the second node with val = 2, and so on. The graph is represented in the test case using an adjacency list.

Adjacency list is a collection of unordered lists used to represent a finite graph. Each list describes the set of neighbors of a node in the graph.

The given node will always be the first node with val = 1. You must return the copy of the given node as a reference to the cloned graph.

*/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func MakeNode(v int) *Node {
    var n Node
    n.Val = v
    n.Neighbors = make([]*Node,0)
    return &n
}


func GetAllNodes(node *Node, M map[*Node] *Node) *Node{
    _,ok := M[node]
    if ok { return nil}
    M[node] = MakeNode(node.Val)
    for _,x:= range(node.Neighbors) {
        GetAllNodes(x, M)
    }
    return M[node]
}

func cloneGraph(node *Node) *Node {
    m := make(map[*Node] *Node)
    if node == nil {return nil}
    clone := GetAllNodes(node, m)
    for o,c := range(m) {

        for _,n := range(o.Neighbors) {
            c.Neighbors = append(c.Neighbors,m[n])
        }
        fmt.Println(o,c)
    }
    return clone
}
