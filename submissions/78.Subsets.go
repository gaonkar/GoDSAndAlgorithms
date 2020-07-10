/*
Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

Input: nums = [1,2,3]
Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]

The solution is more like a generating function. Each iteration, we generate next
combination of numbers of size n.

*/import "sort"

func subsets(nums []int) (r [][]int) {
	r = append(r, []int{})
	sort.Ints(nums)
	for _, x := range nums {
		for i, _ := range r {
			t := make([]int, len(r[i]))
			copy(t, r[i])
			t = append(t, x)
			r = append(r, t)
		}
		//fmt.Println(r)
	}
	return r
}

func subsets(nums []int) (r [][]int) {
	var helper func(stk []int, idx int)

	helper = func(stk []int, idx int) {
		tmp := make([]int, len(stk))
		copy(tmp, stk)
		r = append(r, tmp)
		for i := idx; i < len(nums); i++ {
			stk = append(stk, nums[i])
			helper(stk, i+1)
			stk = stk[:len(stk)-1]
		}
	}
	helper([]int{}, 0)
	return r
}

func subsetsWithDup(nums []int) (r [][]int) {
	var helper func(stk []int, idx int)

	sort.Ints(nums)
	helper = func(stk []int, idx int) {
		tmp := make([]int, len(stk))
		copy(tmp, stk)
		r = append(r, tmp)
		for i := idx; i < len(nums); i++ {
			if i > idx && nums[i] == nums[i-1] {
				continue
			}
			stk = append(stk, nums[i])
			helper(stk, i+1)
			stk = stk[:len(stk)-1]
		}
	}
	helper([]int{}, 0)
	return r
}
