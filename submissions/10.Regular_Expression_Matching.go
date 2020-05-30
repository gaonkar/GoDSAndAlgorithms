/*
Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

Note:

s could be empty and contains only lowercase letters a-z.
p could be empty and contains only lowercase letters a-z, and characters like . or *.
Example 1:

Input:
s = "aa"
p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
Example 2:

Input:
s = "aa"
p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
Example 3:

Input:
s = "ab"
p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
Example 4:

Input:
s = "aab"
p = "c*a*b"
Output: true
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".
Example 5:

Input:
s = "mississippi"
p = "mis*is*p*."
Output: false
*/

//recursive with memoization
// recursive is faster
func IM(s string, p string, Map map[string] bool, l int) bool {
    if len(p) == 0 {return len(s) == 0}
    key := s + "|" + p
    _, ok := Map[key]
    if ok {
        //fmt.Println("M", s, p, l)
        return Map[key]
    }

    M := (len(s) > 0) && (s[0] == p[0] || p[0] == '.')
    //fmt.Println(l, M, s, p)
    if len(p) >= 2 && p[1] == '*'{
        M1 := IM(s,p[2:], Map, l+1)
        Map[s + ":"+ p[2:]] = M1
        if M1 { return M1}
        //fmt.Println("*",l, M1, s[1:], p[2:], M)
        if M {
            //fmt.Println("x")
            M1 = IM(s[1:], p,Map, l+1)
            Map[s[1:] + ":" + p] = M
            //fmt.Println("_",l, M1, s[1:], p[2:],M)
            return M1
        }
    } else {
        if M {
            M =  IM(s[1:], p[1:], Map, l+1)
            Map[s[1:] + ":" + p[1:]] = M
        }
    }
    //fmt.Println("P", key, M)
    Map[key] = M
    return M
}
func isMatch(s string, p string) bool {
    M := make(map[string] bool)
    ret := IM(s,p, M,0)
    //fmt.Println(M)
    return ret
}


