/*
1044. Longest Duplicate Substring Rabin Karp
Given a string S, consider all duplicated substrings: (contiguous) substrings of S that occur 2 or more times.  (The occurrences may overlap.)

Return any duplicated substring that has the longest possible length.  (If S does not have a duplicated substring, the answer is "".)



Example 1:

Input: "banana"
Output: "ana"
Example 2:

Input: "abcd"
Output: ""


Note:

2 <= S.length <= 10^5
S consists of lowercase English letters.
*/
func Power(n int, M int64) []int64 {
    p := int64(1)
    r := make([]int64, 0)
    for i:=0; i < n;i++{
        r = append(r, p)
        p = (p * 26)%M
    }
    return r
}

func RKRollingHash(s string, M int64) (r int64) {

    r = 0
    for i:=0; i < len(s); i++ {
        r = ((r * 26) % M + int64(s[i]-'a')) % M
    }
    return r
}




/*
Find if duplicate of length n exists
*/
func DupN(S string, n int, pow int64, M int64) int {
    found := make(map[int64] int) // map hash to index
    hash := RKRollingHash(S[:n], M)
    found[hash] = 0
    for i:=n; i < len(S);i++ {
        //fmt.Println(i,S[i-n:i], hash, RKRollingHash(S[i-n:i], M), S[i], S[i-n], found )
        e := int64(S[i]-'a')
        s := int64(S[i-n]-'a')
        hash = (hash + M - (pow * s)%M)%M
        hash = ((hash * 26) %M + e) %M
        pos,ok := found[hash]
        if ok {
            //need a string compare, but for now ok

            if S[pos:pos+n] == S[i-n+1:i+1] {
                return pos
            }
        }
        found[hash] = i-n+1
    }
    return -1
}


func longestDupSubstring(S string) string {
    M := int64(100003559) // prime number
    pow := Power(len(S)+1, M)
    //fmt.Println(pow)
    lo := 1
    hi := len(S)
    fpos := -1
    flen := 0
    for lo < hi {
        mi := lo + (hi-lo)/2

        pos := DupN(S, mi, pow[mi-1], M)
        //fmt.Println(fpos,flen,mi,pos)
        if pos == -1 {
            //not found
            hi = mi
        } else {
            if mi > flen {
                fpos = pos
                flen = mi
            }
            lo = mi+1
        }
    }
    if fpos < 0 {return ""}
    //fmt.Println(S[fpos:fpos+flen])
    return S[fpos:fpos+flen]
}

