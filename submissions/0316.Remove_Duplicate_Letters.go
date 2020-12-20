/*
iven a string s, remove duplicate letters so that every letter appears once and only once.
You must make sure your result is the smallest in lexicographical order among all possible results.

Note: This question is the same as 1081: https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/



Example 1:

Input: s = "cdadabcc"
Output: "adbc"
Example 2:

Input: s = "abcd"
Output: "abcd"
Example 3:

Input: s = "ecbacba"
Output: "eacb"
Example 4:

Input: s = "leetcode"
Output: "letcod"


Constraints:

1 <= s.length <= 104
s consists of lowercase English letters.

Other stack problems
    1130 Minimum Cost Tree From Leaf Values
    1081 same problem
    907 Sum of Subarray Minimums
    901 Online Stock Span
    856 Score of Parentheses
        503 Next Greater Element II
        496 Next Greater Element I
    84 Largest Rectangle in Histogram
    42 Trapping Rain Water

*/


func removeDuplicateLetters(s string) string {
    carr := make([]rune, 0)
    last := make(map[rune] int)
    for i,x := range(s) {
        if last[x] < i {last[x]=i}
    }
    for i,x := range(s) {
        found := false
        for _, y := range(carr) {
            if x == y {
                found = true
                break
            }
        }
        if found {continue}
        L := len(carr)
        for L > 0 && carr[L-1] > x && i < last[carr[L-1]] {
            L--
        }
        carr = carr[:L]
        carr = append(carr, x)
    }

    return string(carr)
}
