func BinarySearch(nums []int, target int) int {
    low := 0
    hi := len(nums)-1
    for low <= hi {
        mid := (hi - low) /2 + low ;
        //fmt.Println(low, mid, hi)
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            hi = mid - 1
        } else {
            low = mid + 1
        }

    }
    return low // insert location
}
