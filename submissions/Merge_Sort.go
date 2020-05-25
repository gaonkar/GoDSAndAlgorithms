func mergesort(nums []int) []int {
    if len(nums) == 1 || len(nums) == 0 {
        return nums
    }
    left := nums[0:len(nums)/2]
    right := nums[len(nums)/2:]
    left = mergesort(left)
    right = mergesort(right)
    return merge(left, right)
}

func merge(left, right []int) []int {
    var result []int
    for len(left) != 0 && len(right) != 0 {
        if left[0] < right[0] {
            result = append(result, left[0])
            left = left[1:]
        } else {
            result = append(result, right[0])
            right = right[1:]
        }
    }
    for _, n := range left {
        result = append(result, n)
    }
    for _, n := range right {
        result = append(result, n)
    }
    return result
}
