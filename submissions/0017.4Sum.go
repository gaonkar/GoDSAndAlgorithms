import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target, idx1, idx2 int) (r [][]int) {

	if idx1 > 0 && nums[idx1] == nums[idx1-1] {
		return nil
	}
	if idx2 < len(nums)-1 && nums[idx2] == nums[idx2+1] {
		return nil
	}

	target = target - nums[idx1] - nums[idx2]
	idx1++
	idx2--
	for idx1 < idx2 {
		//fmt.Println(idx1, idx2, target)
		if nums[idx1]+nums[idx2] == target {
			//fmt.Println(target, idx1, idx2)
			r = append(r, []int{nums[idx1], nums[idx2]})
			for nums[idx1+1] == nums[idx1] && idx1 < idx2 {
				idx1++
			}
			for nums[idx2-1] == nums[idx2] && idx1 < idx2 {
				idx2--
			}
		}

		if nums[idx1]+nums[idx2] > target {
			idx2--
		} else {
			idx1++
		}
	}
	return r
}

func Append(r [][]int, a [][]int, x, y int) [][]int {
	if len(a) == 0 {
		return r
	}
	for _, z := range a {
		tmp := make([]int, 4)
		tmp[0] = x
		tmp[1] = z[0]
		tmp[2] = z[1]
		tmp[3] = y
		r = append(r, tmp)
	}
	return r
}

func fourSum(nums []int, target int) (r [][]int) {
	sort.Ints(nums)

	fmt.Println(nums)
	for l := 0; l < len(nums)-1; l++ {
		for h := len(nums) - 1; h > l+2; h-- {
			x := twoSum(nums, target, l, h)
			r = Append(r, x, nums[l], nums[h])
		}
	}
	return r
}
