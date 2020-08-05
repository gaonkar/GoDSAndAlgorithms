/*
Given a string s, partition s such that every substring of the partition is a palindrome.

Return the minimum cuts needed for a palindrome partitioning of s.

Example:

Input: "aab"
Output: 1
Explanation: The palindrome partitioning ["aa","b"] could be produced using 1 cut.
*/

func Min(x,y int) int {
    if x < y {return x}
    return y
}

func minCut(s string) int {
    L := len(s)
    if (L < 2) {return 0}
    P :=make([][]bool, L)
    C := make([]int, L+1)
    for i,_ :=range(P) {
        P[i] = make([]bool, L)
    }
    for i := 0; i <= L; i++ {
       C[i] = i-1
    }

    for j := 1; j < L;j++  {
        C[j+1] = j
        for i := j; i >= 0; i-- {
            if s[i] == s[j]{
                if j - i < 2 || P[i+1][j-1] == true {
                    //fmt.Println(i,j, s[i], s[j])
                    P[i][j] = true
                    C[j+1] = Min(C[j+1], 1 + C[i])
                }
            }
        }
    }
    return C[L]
}
