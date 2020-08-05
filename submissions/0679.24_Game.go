/*

You have 4 cards each containing a number from 1 to 9. You need to judge whether they could operated through *, /, +, -, (, ) to get the value of 24.

Example 1:

Input: [4, 1, 8, 7]
Output: True
Explanation: (8-4) * (7-1) = 24
Example 2:

Input: [1, 2, 1, 2]
Output: False
Note:

The division operator / represents real division, not integer division. For example, 4 / (1 - 2/3) = 12.
Every operation done is between two numbers. In particular, we cannot use - as a unary operator. For example, with [1, 1, 1, 1] as input, the expression -1 - 1 - 1 - 1 is not allowed.
You cannot concatenate numbers together. For example, if the input is [1, 2, 1, 2], we cannot write this as 12 + 12.

Permutations
*/
PUnique(arr []int, n int)[][]int{
    var helper func([]int, int, []int, []bool)
    res := [][]int{}
    tmp := []int{}
    usd := make([]bool, len(arr))
    helper = func(arr []int, N int, val []int, u []bool){
        if len(val) == N{
            tmp := make([]int, N)
            copy(tmp, val)
            res = append(res, tmp)
        } else {
            for i:=0; i < len(arr);i++ {
                l := len(val)
                if u[i] {continue}
                u[i] = true
                val = append(val, arr[i])
                helper(arr, N, val, u)
                val = val[:l]
                u[i] = false
            }
        }
    }
    helper(arr, n, tmp, usd)
    return res
}
func PDuplicates(arr []int, n int)[][]int{
    var helper func([]int, int, []int)
    res := [][]int{}
    tmp := []int{}
    helper = func(arr []int, N int, val []int){
        if len(val) == N{
            tmp := make([]int, N)
            copy(tmp, val)
            res = append(res, tmp)
        } else {
            for i:=0; i < len(arr);i++ {
                l := len(val)
                val = append(val, arr[i])
                helper(arr, N, val)
                val = val[:l]
            }
        }
    }
    helper(arr, n, tmp)
    return res
}

func Op(x,y float64, o int) float64 {
    if o == 1 { return x + y}
    if o == 2 { return x - y}
    if o == 3 { return x * y}
    if o == 4 { return x / y}
    return 0
}

func Eval(p, o []int) bool {
    a := Op(float64(p[0]), float64(p[1]), o[0])
    b := Op(float64(p[2]), float64(p[3]), o[1])
    c := Op(a,b,o[2])
    if 23.99 <= c && c <= 24.02 {return true}
    c = Op(float64(p[0]), float64(p[1]), o[0])
    c = Op(c, float64(p[2]), o[1])
    c = Op(c, float64(p[3]),o[2])

    if 23.99 <= c && c <= 24.02 {return true}
    c = Op(float64(p[0]), float64(p[1]), o[0])
    c = Op(float64(p[2]),c, o[1])
    c = Op(float64(p[3]),c, o[2])
    if 23.99 <= c && c <= 24.02 {return true}
    return false
}
func judgePoint24(nums []int) bool {
    O := []int{1,2,3,4}
    p := PUnique(nums, 4)
    o := PDuplicates(O, 3)
    //fmt.Println(PUnique(nums,4))
    for _, x := range(o) {
        for _, y := range(p) {
            if Eval(y,x) {return true}
        }
    }
    return false
}

