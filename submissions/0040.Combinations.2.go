/*
Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

Each number in candidates may only be used once in the combination.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
Example 1:

Input: candidates = [10,1,2,7,6,1,5], target = 8,
A solution set is:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
Example 2:

Input: candidates = [2,5,2,1,2], target = 5,
A solution set is:
[
  [1,2,2],
  [5]
]

*/
var r [][]int
func CHelper(c []int, t int, idx int, tmp []int)  {
    if t <= 0 {
        if t == 0 {
            var A []int
            //fmt.Println("X:",tmp)
            for _,x:= range(tmp) {
                A = append(A, x)
            }
            r = append(r, A)
        }
        //fmt.Println("<",r)
        return
    }
    for i := idx ; i < len(c); i++ {
        if i > idx && c[i] == c[i-1] || c[i] > t {
            //ignore duplicates and if sum goes beyond
            continue
        }
        tmp = append(tmp, c[i])
        CHelper(c,t - c[i], i+1, tmp)
        tmp = tmp[:len(tmp)-1]
    }

}

func combinationSum2(c []int, t int) [][]int {
    var tmp []int
    r = make([][]int, 0)
    //b := make([]bool, len(c))
    sort.Ints(c)
    //fmt.Println(c)
    CHelper(c,  t, 0, tmp)
    //fmt.Println(r)
    return r
}


