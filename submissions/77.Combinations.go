/*
Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.

Example:

Input: n = 4, k = 2
Output:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/
func combine(n int, k int) [][]int {
    var helper func([]int, int, int, []int, []bool)
    res := [][]int{}
    tmp := []int{}
    arr := make([]int, n)
    for i:=0; i < n ; i++ {arr[i] = i+1}
    usd := make([]bool, len(arr))
    helper = func(arr []int, N, pos int, val []int, u []bool){
        if len(val) == N{
            tmp := make([]int, N)
            copy(tmp, val)
            res = append(res, tmp)
            return
        }

        for i:=pos; i < len(arr);i++ {
            l := len(val)
            val = append(val, arr[i])
            helper(arr, N,i+1, val, u)
            val = val[:l]
        }
    }
    helper(arr, k,0, tmp, usd)
    return res
}
