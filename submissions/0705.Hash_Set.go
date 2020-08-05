/*
Design a HashSet without using any built-in hash table libraries.

To be specific, your design should include these functions:

add(value): Insert a value into the HashSet.
contains(value) : Return whether the value exists in the HashSet or not.
remove(value): Remove a value in the HashSet. If the value does not exist in the HashSet, do nothing.

Example:

MyHashSet hashSet = new MyHashSet();
hashSet.add(1);
hashSet.add(2);
hashSet.contains(1);    // returns true
hashSet.contains(3);    // returns false (not found)
hashSet.add(2);
hashSet.contains(2);    // returns true
hashSet.remove(2);
hashSet.contains(2);    // returns false (already removed)

Note:

All values will be in the range of [0, 1000000].
The number of operations will be in the range of [1, 10000].
Please do not use the built-in HashSet library.
*/
type Node struct {
    key int
    Next *Node

}
type MyHashSet struct {
    N int
    B []*Node
}

func (this *MyHashSet)Find(key int) (*Node, *Node) {
    find := this.B[key % this.N]
    prev := find
    for find != nil  && find.key != key{
        prev = find
        find = find.Next
    }
    return find, prev
}
/** Initialize your data structure here. */
func Constructor() MyHashSet {
    var H MyHashSet
    H.N = 30011
    H.B = make([]*Node, 30011)
    return H
}

func (this *MyHashSet) Add(key int)  {
    var v Node
    head := this.B[key % this.N]
    find, _ := this.Find(key)
    if find != nil {
        return
    }
    v.key = key
    v.Next = head
    this.B[key % this.N] = &v
}


func (this *MyHashSet) Remove(key int)  {
    find,prev := this.Find(key)
    if find == nil {
        return
    }
    if find == prev {
        // head
        this.B[key % this.N] = find.Next
        return
    }
    prev.Next = find.Next
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
    find, _ := this.Find(key)
    if find != nil {
        return true
    }
    return false
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
