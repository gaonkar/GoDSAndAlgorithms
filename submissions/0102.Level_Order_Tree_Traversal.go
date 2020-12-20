/*
102. Binary Tree Level Order Traversal
Medium

2675

65

Add to List

Share
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:

[
  [3],
  [9,20],
  [15,7]
]

*/import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TL struct {
	r *TreeNode
	l int
}

func levelOrder(root *TreeNode) (r [][]int) {
	var q []*TL
	l := 0
	L := 0
	if root == nil {
		return r
	}
	q = append(q, &TL{root, 0})
	r = make([][]int, 1)
	for len(q) > 0 {
		root = q[0].r
		l = q[0].l
		fmt.Println(root.Val, l, r)
		r[l] = append(r[l], root.Val)
		q = q[1:]
		l = l + 1
		if L < l {
			r = append(r, make([]int, 0))
			L = l
		}
		if root.Left != nil {
			q = append(q, &TL{root.Left, l})
		}
		if root.Right != nil {
			q = append(q, &TL{root.Right, l})

		}
	}
	return r[:len(r)-1]
}

/*
 * Problem 958 Complete Binary Tree, Two additional constraints.
   1. Make sure when you see null child, its the last in the current level.
   2. Make sure that L-1 levels have 2 * L-1 nodes.
*/
func LevelOrder(root *TreeNode) (r [][]*TreeNode) {
	q := []*TreeNode{root}
	for len(q) > 0 {
		r = append(r, q)
		t := []*TreeNode{}
		for i := 0; i < len(q); i++ {
			if q[i].Left != nil {
				t = append(t, q[i].Left)
			}
			if q[i].Right != nil {
				t = append(t, q[i].Right)
			}
		}
		q = t
	}
	return r
}
