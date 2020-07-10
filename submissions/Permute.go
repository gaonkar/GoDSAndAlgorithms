/*
46: Permutation problem

*/
func permuteUnique(nums []int) [][]int {
    if len(nums) == 1 {
        return [][]int{nums}
    }
    var result [][]int
    dict := make(map[int] bool)
    for i,n := range nums {
        _,ok := dict[n];
        if ok {
                continue
        }
        dict[n] = true
        nums2 := make([]int,len(nums))
        copy(nums2,nums)
        temp := permuteUnique(append(nums2[:i],nums2[i+1:]...))
        for _,m := range temp {
            m = append(m,n)
            result = append(result,m)
        }
    }
    return result
}
