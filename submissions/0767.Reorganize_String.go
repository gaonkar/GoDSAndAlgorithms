/*
Given a string S, check if the letters can be rearranged so that two characters that are adjacent to each other are not the same.

If possible, output any possible result.  If not possible, return the empty string.

Example 1:

Input: S = "aab"
Output: "aba"
Example 2:

Input: S = "aaab"
Output: ""
Note:

S will consist of lowercase letters and have length in range [1, 500].
*/




type B struct {
    c int
    x rune
}

func reorganizeString(S string) (r string) {
    F := make([]B, 26)
    ret := make([]rune, len(S))
    for _, x := range(S) {
        idx := int(x - 'a')
        F[idx].c++
        F[idx].x = x
    }
    sort.SliceStable(F, func(i, j int) bool { return F[i].c > F[j].c})
    if F[0].c *2 - 1 > len(S) {return ""}
    j := 0
    for i:=0; i < len(F);i++ {
        for F[i].c > 0 {
            ret[j] = F[i].x
            F[i].c--
            j+=2
            if j >= len(S) {j = 1} //flipt to fill odd locatons
        }
    }
    return string(ret)

}
