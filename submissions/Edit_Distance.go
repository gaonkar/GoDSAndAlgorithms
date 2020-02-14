import "fmt"

/*

  (i-1,j-1)     (i-1,j)
  (i,j-1)       (i,j)

i is entry in s1 (0 < i< n1)
j is entry in s2 (0 < j< n2)


DP[i,0]= i
DP[0,j]= j
DP[i,j] = Min(A,B,C)
            A = DP[i-1,j] + 1 // insert
            B = DP[i,j-1] + 1 // delete
            C = DP[i-1,j-1] + X and X = 1  if S[i] != S[j] else 0 //replace
*/

func Min(x int, y int, z int) int {
    if x < y {
        if x < z {
            return x
        }
        return z
    } else if y < z {
            return y
    }
    return z
}

func minDistance(s1 string, s2 string) int {
    n1:= len(s1)
    n2:= len(s2)
    var curr = make([]int, n1+1)
    var prev = make([]int, n1+1)
    // makes life easy to move from 1 to N
    s1 ="0" + s1
    s2 = "0" + s2
    //fmt.Println(curr, prev)
    for i:= 0; i <= n1; i++ {
        prev[i] = i
    }
    //fmt.Println(prev)
    for j := 1; j <= n2; j++ {
        curr[0] = j
        for i:= 1; i <= n1; i++ {
            X := 1
            if s1[i] == s2[j] { X = 0}
            A := prev[i] + 1
            B := curr[i-1] + 1
            C := prev[i-1] + X
            //fmt.Println(A,B,C)
            curr[i] = Min(A,B,C)
        }
        prev = curr
        curr = make([]int, n1+1)
        //fmt.Println(prev)
    }
    return prev[n1]
}
