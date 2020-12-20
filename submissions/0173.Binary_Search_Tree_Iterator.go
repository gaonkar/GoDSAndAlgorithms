/*

Implement an iterator over a binary search tree (BST). Your iterator will be initialized with the root node of a BST.

Calling next() will return the next smallest number in the BST.

*/
// Python solution
class BSTIterator:

    def __init__(self, root: TreeNode):
        self.stk=[]
        self.GoLeft(root)
        
    def GoLeft(self, root:TreeNode):
        while root != None:
            self.stk.append(root)
            root = root.left
        
    def next(self) -> int:
        """
        @return the next smallest number
        """
        ret = self.stk.pop()
        if ret.right:
            self.GoLeft(ret.right)
        return ret.val    

        
    def hasNext(self) -> bool:
        """
        @return whether we have a next smallest number
        """
        #print(len(self.stk), self.current)
        return len(self.stk) > 0

// Golang solution
type BSTIterator struct {
    stk []*TreeNode
}

func (this *BSTIterator) PushLeft(r *TreeNode){
    for r != nil {
        this.stk = append(this.stk, r)
        r = r.Left
    }
}



func Constructor(root *TreeNode) BSTIterator {
    var iter BSTIterator
    
    iter.stk = make([]*TreeNode, 0)
    iter.PushLeft(root)
    return iter
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
    L := len(this.stk) - 1
    ret := this.stk[L]
    this.stk = this.stk[:L]
    if ret.Right != nil {
        this.PushLeft(ret.Right)
    }

    return ret.Val
}


/** Peek the next  number */
func (this *BSTIterator) Peek() int {
    L := len(this.stk) - 1
    return this.stk[L].Val
}



/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    return len(this.stk) > 0
}


