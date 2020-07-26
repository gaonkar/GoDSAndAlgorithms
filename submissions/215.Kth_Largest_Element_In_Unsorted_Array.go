/*
Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:

Input: [3,2,1,5,6,4] and k = 2
Output: 5
Example 2:

Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
Note:
You may assume k is always valid, 1 ≤ k ≤ array's length.
*/
//standard Lomuto partition function
func Partition(nums []int) int {
    L := len(nums)
    p := nums[L-1]
    j :=0
    for i:=0;i< len(nums)-1; i++{
        if nums[i] < p {
            nums[i], nums[j] = nums[j], nums[i]
            j++
        }
        //fmt.Println(nums, "i=", i,"j=", j, p)
    }
    nums[j], nums[L-1] = nums[L-1], nums[j]
    //fmt.Println(nums, "j=", j, p)
    return j
}


func findKthLargest(nums []int, k int) int {
    //fmt.Println(nums, k)
    L := len(nums)
    P := Partition(nums)
    K := L - P
    //fmt.Println(nums, k, K)
    if (K == k) {
        return nums[P]
    } else if K > k {
        return findKthLargest(nums[P+1:], k)
    }
    return findKthLargest(nums[:P], k-K)
}
