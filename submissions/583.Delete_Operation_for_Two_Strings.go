/*
Given two words word1 and word2, find the minimum number of steps required to make word1 and word2 the same, where in each step you can delete one character in either string.

Example 1:

Input: "sea", "eat"
Output: 2
Explanation: You need one step to make "sea" to "ea" and another step to make "eat" to "ea".
Note:

The length of given words won't exceed 500.
Characters in given words can only be lower-case letters.
*/

/*
    DP[i][j] -->    number of steps to make W1[0..i-1] same as W2[0..i-2]
    DP[i][0] --> i  Delete all the i from W1 to make it null
    DP[0][j] --> j  Delete all the i from W2 to make it null
    DP[i][j] --> D[i-1][j-1] if W1[i-1] == W[j-1]
                 Min(D[i][j-1], D[i-1][j]) + 1


*/

func Min(x, y int) int{
    if x < y {
        return x
    }
    return y
}

func MinDel(W1 string, W2 string) int {
    var curr = make([]int, len(W1)+1)
    var prev = make([]int, len(W1)+1)
    //fmt.Println(curr, prev)
    for i := len(W1); i >= 0; i-- { // filling W2="" string scenario
        prev[i] = i
    }
    for j := 1; j <= len(W2); j++ {
        curr[len(W1)] = j+1
        for i := 1; i <= len(W1); i++ {
            if W1[i-1] == W2[j-1] {
               curr[i] = curr[i-1]
            } else {
                curr[i] = 1 + Min(curr[i-1], prev[i])
            }
        }
        //fmt.Println(curr)
        prev = curr
        curr = make([]int, len(W1)+1)
    }
    //fmt.Println(prev)
    return prev[len(W1)]
}


func minDistance(word1 string, word2 string) int {
    return MinDel(word1, word2)
}
