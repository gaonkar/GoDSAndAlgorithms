/*
Given a string, find the length of the longest substring without repeating characters.
Status
Accepted	4 ms	2.6 MB	golang
*/
import "fmt"
func Max(x, y int) int{
    if x < y {
        return y
    }
    return x
}

func lengthOfLongestSubstring(s string) int {
    ln := 1
    best := 1
    if s == "" {
        return 0
    }
    for i:= range s {
        
        for j:= i-1; j > i-ln; j-- {
            //fmt.Println(i, j, s[i], s[j], ln, best)
            if s[j] == s[i] {
                ln = i - j
            }
        }
        best = Max(ln, best)
        ln++
    }
    return best
}


