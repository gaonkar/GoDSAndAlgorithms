/*
Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

The update(i, val) function modifies nums by updating the element at index i to val.

Example:

Given nums = [1, 3, 5]

sumRange(0, 2) -> 9
update(1, 2)
sumRange(0, 2) -> 8
Note:

The array is only modifiable by the update function.
You may assume the number of calls to update and sumRange function is distributed evenly

Notes
Apparently, this is called the binary index tree.  The node is equal to sume of all its children on this left and itself.
Log(n) update and query

*/import "fmt"

type Tree struct {
	Idx   int
	Val   int
	Sum   int
	Left  *Tree
	Right *Tree
}
type NumArray struct {
	root *Tree
	sum  []int
	nums []int
}

func PrintTree(root *Tree) {
	if root == nil {
		return
	}
	fmt.Println(root)
	PrintTree(root.Left)
	PrintTree(root.Right)
}

func BuildTree(nums, sum []int, lo, hi int) *Tree {
	var root Tree
	if lo > hi {
		return nil
	}
	mi := lo + (hi-lo)/2
	//fmt.Println(mi, lo,hi, sum[lo], sum[hi], sum[mi])
	if lo == hi {
		root.Val = nums[lo]
		root.Idx = lo
		root.Sum = nums[lo]
		return &root
	}

	root.Val = nums[mi]
	root.Idx = mi
	root.Sum = sum[mi] - sum[lo] + nums[lo]

	root.Left = BuildTree(nums, sum, lo, mi-1)
	root.Right = BuildTree(nums, sum, mi+1, hi)

	return &root
}

func Regular(s, n []int) {
	s[0] = n[0]
	for i := 1; i < len(n); i++ {
		s[i] = s[i-1] + n[i]
	}
}

func Constructor(nums []int) NumArray {
	var N NumArray
	if len(nums) == 0 {
		return N
	}
	N.sum = make([]int, len(nums))
	N.nums = make([]int, len(nums))
	N.sum = make([]int, len(nums))
	copy(N.nums, nums)
	Regular(N.sum, nums)
	N.root = BuildTree(nums, N.sum, 0, len(nums)-1)
	//PrintTree(N.root)
	return N
}

func Sum(root *Tree, idx, psum int) int {
	if root.Idx == idx {

		return psum + root.Sum
	}
	if idx < root.Idx {
		return Sum(root.Left, idx, psum)
	}
	return Sum(root.Right, idx, psum+root.Sum)
}

func Update(root *Tree, idx, val int) int {
	if root.Idx == idx {
		diff := val - root.Val
		//fmt.Println(idx, val, root.Val, root.Sum)
		root.Val = val
		root.Sum += diff
		return diff
	}
	if idx < root.Idx {
		diff := Update(root.Left, idx, val)
		root.Sum += diff
		return diff
	}
	return Update(root.Right, idx, val)
}

func (this *NumArray) Update(i int, val int) {
	//fmt.Println(i,val)
	Update(this.root, i, val)
	this.nums[i] = val
	//Regular(this.sum,this.nums)
}

func (this *NumArray) SumRange(i int, j int) int {
	x := Sum(this.root, j, 0)
	if i > 0 {
		x -= Sum(this.root, i-1, 0)
	}
	return x
}
