/*
Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place and use only constant extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/
func Reverse(r []int, i,n int) {
    //fmt.Println(r,i,n)
    for i < n {
        r[i], r[n] = r[n], r[i]
        n--
        i++
    }
}
//https://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order
func nextPermutation(nums []int)  {
    i:=0
    L:=len(nums)
    k:=-1
    // find the largest index k such that a[k] < a[k+1]
    for i < L-1  {
        if nums[i] < nums[i+1] {k=i}
        i++
    }
    if k == -1{
        Reverse(nums,0, L-1)
        return
    }
    // find the number that is greater than nums[k] from k+2 to N
    l := L-1
    for l > k+1 && nums[l] <= nums[k]{
        l--
    }
    //fmt.Println(nums, k,nums[k], l, nums[l])
    // swap them
    nums[k], nums[l] = nums[l], nums[k]
    Reverse(nums,k+1,L-1)
}
