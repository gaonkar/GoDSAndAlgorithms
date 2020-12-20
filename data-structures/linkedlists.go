/*
Design your implementation of the linked list. You can choose to use the singly linked list or the doubly linked list.
A node in a singly linked list should have two attributes: val and next. val is the value of the current node, and next
is a pointer/reference to the next node. If you want to use the doubly linked list, you will need one more attribute prev
to indicate the previous node in the linked list. Assume all nodes in the linked list are 0-indexed.

Implement these functions in your linked list class:

get(index) : Get the value of the index-th node in the linked list. If the index is invalid, return -1.

addAtHead(val) : Add a node of value val before the first element of the linked list. After the insertion, the new node 
will be the first node of the linked list.

addAtTail(val) : Append a node of value val to the last element of the linked list.

addAtIndex(index, val) : Add a node of value val before the index-th node in the linked list. If index equals to the
length of linked list, the node will be appended to the end of linked list. If index is greater than the length,
the node will not be inserted.

deleteAtIndex(index) : Delete the index-th node in the linked list, if the index is valid.


Example:

Input:
["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]
[[],[1],[3],[1,2],[1],[1],[1]]
Output:
[null,null,null,null,2,null,3]

Explanation:
MyLinkedList linkedList = new MyLinkedList(); // Initialize empty LinkedList
linkedList.addAtHead(1);
linkedList.addAtTail(3);
linkedList.addAtIndex(1, 2);  // linked list becomes 1->2->3
linkedList.get(1);            // returns 2
linkedList.deleteAtIndex(1);  // now the linked list is 1->3
linkedList.get(1);            // returns 3
*/

type Node struct {
    Val int
    Next *Node
}
type MyLinkedList struct {
    head, tail *Node
    C    int
}

func (this *MyLinkedList) Print(s string, v int) {
    h := this.head
    return
    fmt.Print(s, " ", v, ":")
    for h != nil {
        fmt.Print(h.Val, ",")
        h = h.Next
    }
    fmt.Println("")
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
    var M MyLinkedList
    M.head = nil
    M.tail = nil
    return M
}

func (this *MyLinkedList) GetPtr(index int) *Node {
    if index+1 > this.C { return nil}
    h := this.head
    for index > 0 {
        h = h.Next
        index--
    }
    return h
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
    this.Print(" Get ", index)
    h := this.GetPtr(index)
    if h == nil { return -1}
    return h.Val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
    this.head = &Node{val, this.head}
    if this.C == 0 {
        this.tail = this.head
    }
    this.Print("AHead", val)
    this.C++
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
    tmp := &Node{val, nil}
    if this.tail == nil {
        this.tail = tmp
        this.head = tmp
    } else {
        this.tail.Next = tmp
        this.tail = tmp
    }
    this.Print("ATail", val)
    this.C++
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list,
the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    fmt.Println(index, val, this.C)
    if index > this.C {return}
    if index <= 0 {
        this.AddAtHead(val)
        return
    }
    if index == this.C {
        this.AddAtTail(val)
        return
    }
    tmp := this.GetPtr(index-1)
    tmp.Next = &Node{val, tmp.Next}
    this.C++
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index+1 > this.C {return}
    if this.C == 1 {
        this.head = nil
        this.tail = nil
        this.C = 0
        return
    }
    if index == 0 {
        this.head = this.head.Next
        this.C--
        return
    }
    tmp := this.GetPtr(index-1)
    tmp.Next = tmp.Next.Next
    if index+1 == this.C {
        this.tail = tmp
    }
    this.C--
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
