/*
Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.

Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.

The order of output does not matter.

Example 1:

Input:
s: "cbaebabacd" p: "abc"

Output:
[0, 6]

Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
Example 2:

Input:
s: "abab" p: "ab"

Output:
[0, 1, 2]

Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".
*/
func Match(x []int, y []int) bool {
    for i:=0; i < len(x); i++ {
        if x[i] != y[i] {return false}
    }
    return true
}

func findAnagrams(s string, p string) []int {
    var r []int
    P := make([]int, 26)
    L := len(p)
    // make a map
    for i:=range(p) {
        P[(int(p[i]) - int('a'))]++
    }

    if L > len(s) {return r}
    i := 0
    S := make([]int, 26)
    // at s[0], make a map of first L chars
    for i < L {
        S[(int(s[i]) - int('a'))]++
        i++
    }
    j := 0
    for i < len(s) {
        //Match the map
        if Match(S,P) {
            r = append(r, j)
        }
        // remove i character and add the jth
        S[(int(s[i]) - int('a'))]++
        S[(int(s[j]) - int('a'))]--
        i++
        j++
    }
    if Match(S,P) {
        r = append(r, j)
    }
    return r
}
