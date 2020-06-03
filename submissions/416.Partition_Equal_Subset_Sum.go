/*
416. Partition Equal Subset Sum
Medium

2313

64

Add to List

Share
Given a non-empty array containing only positive integers, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.

Note:

Each of the array element will not exceed 100.
The array size will not exceed 200.


Example 1:

Input: [1, 5, 11, 5]

Output: true

Explanation: The array can be partitioned as [1, 5, 5] and [11].


Example 2:

Input: [1, 2, 3, 5]

Output: false

Explanation: The array cannot be partitioned into equal sum subsets.
*/

Max(x,y int) int {
    if x < y {return y}
    return x
}
/*
 Recursive time limit exceeded
 */
func KnapRecursive(N []int, S int, idx int) bool{
    if idx >= len(N) || S < N[idx] {return false}
    if S == N[idx] {return true}
    return KnapRecursive(N, S, idx+1) || KnapRecursive(N, S-N[idx], idx+1)
}

/*
 DP 1 Knapsack
// Time Complexity: O(nW)
// Space Complexity: O(nW)

Accepted
*/

func KnapDP(N []int, W int) bool {
    L := len(N)
    D := make([][]bool, L+1)
    for i,_ :=range(D) {
        D[i] = make([]bool, W+1)
    }
    for i:=0; i<=L; i++ {
        D[i][0] = true;
    }
    for i:=1; i <= L; i++ {
        for j := 1; j <= W; j++ {
            if N[i-1] > j {
                D[i][j] = D[i-1][j]
            } else {
                D[i][j] = D[i-1][j] || D[i-1][j-N[i-1]]
            }
        }
    }
    //fmt.Println(D)
    return D[L][W]
}

/*
 DP 2 Knapsack
// Time Complexity: O(nW)
// Space Complexity: O(W)

Accepted
*/

func KnapDP2(N []int, W int) bool {
    L := len(N)
    D := make([]bool, W+1)


    D[0] = true;
    for i:=0; i < L; i++ {
        for j := W; j >= N[i]; j-- {
            // traversing reverse will prevent re-use of the weight more than once
            D[j] = D[j] || D[j-N[i]];
        }
    }
    //fmt.Println(D)
    return D[W]
}

func canPartition(nums []int) bool {
    S := 0
    for _,x := range(nums) { S +=x}
    if S % 2 == 1 || S==0 { return false} // odd sum false
    S = S / 2
    sort.Ints(nums)

    return KnapDP2(nums, S)
}
