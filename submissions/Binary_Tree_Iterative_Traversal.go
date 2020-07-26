/*
Inorder
*/
func inorderTraversal(root *TreeNode) (r []int) {
	if root == nil {
		return nil
	}
	S := make([]*TreeNode, 0)
	s := 0
	for s > 0 || root != nil {
		for root != nil {
			S = append(S, root)
			s++
			root = root.Left
		}
		root = S[s-1]
		S = S[:s-1]
		s--
		r = append(r, root.Val)
		root = root.Right
	}
	return r
}

/*
Preorder
*/

func preorderTraversal(root *TreeNode) (r []int) {
	if root == nil {
		return nil
	}
	S := make([]*TreeNode, 0)
	s := 0
	for s > 0 || root != nil {
		for root != nil {
			r = append(r, root.Val)
			S = append(S, root)
			s++
			root = root.Left
		}
		root = S[s-1]
		S = S[:s-1]
		s--
		root = root.Right
	}
	return r
}

/*
Postorder
*/

type Stack struct {
	a []*TreeNode
}

func MakeStack() Stack {
	var S Stack
	S.a = make([]*TreeNode, 0)
	return S
}

func (this *Stack) Push(r *TreeNode) {
	if r != nil {
		this.a = append(this.a, r)
	}
}

func (this *Stack) Pop() *TreeNode {
	l := len(this.a) - 1
	r := this.a[l]
	this.a = this.a[:l]
	return r
}

func (this *Stack) IsEmpty() bool {
	return len(this.a) == 0
}

func Reverse(r []int) {
	n := len(r) - 1
	for i := 0; i < len(r)/2; i++ {
		r[i], r[n] = r[n], r[i]
		n--
	}
}

/*
   This is also called 2 stack solution. We just reverse r here to get the order.
*/
func postorderTraversal(root *TreeNode) (r []int) {
	if root == nil {
		return nil
	}
	S1 := MakeStack()
	//S2 := MakeStack()
	S1.Push(root)
	for !S1.IsEmpty() {
		root = S1.Pop()
		r = append(r, root.Val)
		S1.Push(root.Left)
		S1.Push(root.Right)
	}
	Reverse(r)
	return r
}

/* Level order zig zag traversal with DQueue
 */
//enqueue front
func EQF(Q *[]*TreeNode, r *TreeNode) {
	if r == nil {
		return
	}
	LQ := []*TreeNode{r}
	(*Q) = append(LQ, (*Q)...)
}

//dqueue front
func DQF(Q *[]*TreeNode) *TreeNode {
	qlen := len(*Q)
	if qlen == 0 {
		return nil
	}
	L := (*Q)[0]
	(*Q) = (*Q)[1:]
	return L
}

//enqueue back
func EQB(Q *[]*TreeNode, r *TreeNode) {
	if r == nil {
		return
	}
	(*Q) = append((*Q), r)
}

//dqueue back
func DQB(Q *[]*TreeNode) *TreeNode {
	qlen := len(*Q)
	if qlen == 0 {
		return nil
	}
	qlen--
	L := (*Q)[qlen]
	(*Q) = (*Q)[:qlen]
	return L
}

func zigzagLevelOrder(root *TreeNode) (r [][]int) {
	Q := make([]*TreeNode, 0)
	EQF(&Q, root)
	zig := false
	n := root
	for len(Q) > 0 {
		qlen := len(Q)
		t := make([]int, 0)
		//fmt.Println(Q)
		for qlen > 0 {
			if zig {
				n = DQF(&Q)
				t = append(t, n.Val)
				EQB(&Q, n.Right)
				EQB(&Q, n.Left)
			} else {
				n = DQB(&Q)
				t = append(t, n.Val)
				EQF(&Q, n.Left)
				EQF(&Q, n.Right)
			}
			qlen--
		}
		r = append(r, t)
		zig = !zig
	}
	return r
}

// Level order Traversal

func LT(root *TreeNode) *Node {
	Q := make([]*Node, 0)
	EQB(&Q, root)
	n := root
	for len(Q) > 0 {
		qlen := len(Q)
		for qlen > 0 {
			n = DQF(&Q)
			qlen--
			EQB(&Q, n.Left)
			EQB(&Q, n.Right)

		}
	}
	return root
}

/*
BST Tree Iterator
*/

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
