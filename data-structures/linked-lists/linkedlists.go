package linkedlists

type SList struct {
    Head *Node
    Tail *Node
}

type Node struct {
    Next *Node
    Data interface{}
}

// create an empty list
func New() *SList {
    return &SList{nil, nil}
}

// check if empty
func (l1 *SList) IsEmpty() bool {
    return (l1.Head == nil)
}

func (sl *SList) Append(d interface{}) *SList {
    node := &Node {nil, d}
    if
