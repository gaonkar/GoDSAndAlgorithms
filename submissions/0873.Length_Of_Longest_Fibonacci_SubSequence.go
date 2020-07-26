/*
A sequence X1, X2, ..., Xn is Fibonacci-like if:

n >= 3
Xi + Xi+1 = Xi+2 for all i + 2 <= n
Given a strictly increasing array arr of positive integers forming a sequence, return the length of the longest
Fibonacci-like subsequence of arr. If one does not exist, return 0.
A subsequence is derived from another sequence arr by deleting any number of elements (including none) from arr,
without changing the order of the remaining elements. For example, [3, 5, 8] is a subsequence of [3, 4, 5, 6, 7, 8].

Example 1:

Input: arr = [1,2,3,4,5,6,7,8]
Output: 5
Explanation: The longest subsequence that is fibonacci-like: [1,2,3,5,8].
Example 2:

Input: arr = [1,3,7,11,12,14,18]
Output: 3
Explanation: The longest subsequence that is fibonacci-like: [1,11,12], [3,11,14] or [7,11,18].


Constraints:

3 <= arr.length <= 1000
1 <= arr[i] < arr[i + 1] <= 109

*/
func Max(i, j int) int {
    if i < j {return j}
    return i
}

func Make2D(r,c, v int) [][] int {
    ret := make([][]int, r)
    for i:=0; i < len(ret); i++ {
        ret[i] = make([]int, c)
        if v > 0 {
            for j := 0; j < c; j++ {
                ret[i][j] = v
            }
        }
    }
    return ret
}


func lenLongestFibSubseq(arr []int) int {
    M := make(map[int] int)
    acc := Make2D(len(arr),len(arr),2)
    for i,x := range(arr) {
        M[x] = i + 1
    }
    ret := 0
    for i := 0; i < len(arr); i++ {
        for j:= 0; j < i; j++ {
            num := arr[i] - arr[j]
            idx := M[num] - 1
            //found the entry and it is smaller than arr[j]
            if idx >= 0 && num < arr[j] {
                acc[j][i] = 1 + acc[idx][j]
                ret = Max(ret, acc[j][i])
            }
        }
    }
    return ret
}
