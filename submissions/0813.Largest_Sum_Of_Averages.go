/*
813. Largest Sum of Averages
Medium

958

46

Add to List

Share
We partition a row of numbers A into at most K adjacent (non-empty) groups, then our score is the sum of the average of each group. What is the largest score we can achieve?

Note that our partition must use every number in A, and that scores are not necessarily integers.

Example:
Input:
A = [9,1,2,3,9]
K = 3
Output: 20
Explanation:
The best choice is to partition A into [9], [1, 2, 3], [9]. The answer is 9 + (1 + 2 + 3) / 3 + 9 = 20.
We could have also partitioned A into [9, 1], [2], [3, 9], for example.
That partition would lead to a score of 5 + 2 + 6 = 13, which is worse.


Note:

1 <= A.length <= 100.
1 <= A[i] <= 10000.
1 <= K <= A.length.
Answers within 10^-6 of the correct answer will be accepted as correct.

*/

/* Without Memoization, TLE
*/
func Make2D(r,c int, v float64) [][] float64 {
    ret := make([][]float64, r)
    for i:=0; i < len(ret); i++ {
        ret[i] = make([]float64, c)
        if v > 0 {
            for j := 0; j < c; j++ {
                ret[i][j] = v
            }
        }
    }
    return ret
}
func largestSumOfAverages(A []int, K int) float64 {
    var dphelper func(i,k int) float64
    L := len(A)
    sum := make([]float64, L+1)
    for i:=0; i < L; i++ {sum[i+1] = sum[i]+float64(A[i])}
    //fmt.Println(sum)
    memo := Make2D(L+1,K+1,-1.0)
    dphelper = func(i,k int) float64 {
        if k == 1 {
            rsum := sum[L] - sum[i-1]
            alen := float64(L-i+1)
            return rsum / alen
        }
        if memo[i][k] > 0 {return memo[i][k]}
        max := 0.0
        for j :=i; j <= L-k+1; j++ {
            alen := float64(j-i+1)
            lmax := dphelper(j+1, k-1) + (sum[j] - sum[i-1]) / alen
            if lmax > max {max = lmax}
        }
        memo[i][k] = max
        return max
    }
    return dphelper(1, K)
}

// DP bottom up
func Make2D(r,c int, v float64) [][] float64 {
    ret := make([][]float64, r)
    for i:=0; i < len(ret); i++ {
        ret[i] = make([]float64, c)
        if v > 0 {
            for j := 0; j < c; j++ {
                ret[i][j] = v
            }
        }
    }
    return ret
}
func largestSumOfAverages(A []int, K int) float64 {
    L := len(A)
    sum := make([]float64, L+1)
    for i:=0; i < L; i++ {sum[i+1] = sum[i]+float64(A[i])}
    //fmt.Println(sum)
    memo := Make2D(K,L+1,-1.0)
    // fill for k = 1
    for i:=0; i < L; i++ {
        alen := float64(L-i)
        memo[0][i+1] = (sum[L] - sum[i])/alen
    }
    for k:=1; k < K; k++ {
        for i:=1; i <= L-k; i++ {
            for j:=i; j <= L-k; j++ {
                avg := (sum[j] - sum[i-1]) / float64(j-i+1)
                avg += memo[k-1][j+1]
                if avg > memo[k][i] {memo[k][i] = avg}
            }
        }
    }
    return memo[K-1][1]
}
