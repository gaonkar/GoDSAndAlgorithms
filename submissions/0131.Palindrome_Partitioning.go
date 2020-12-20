
// Go
func IsPalindrome(s string) bool {
    L := len(s)
    for i:=0; i < L; i++ {
        if s[i] != s[L-1-i] {return false}
    }
    return true
}

func RecPal(s string, tmp []string, res *[][]string) {
    L := len(s)
    if L == 0 {
        cpy := make([]string, len(tmp))
        // I assumed that assignment is a copy operator, spent considerable time debugging this issue
        // unit test case abababb
        copy(cpy, tmp)
        *res = append(*res, cpy)
        return
    }
    for i:=1; i < L+1; i++ {
        pre := s[:i]
        if IsPalindrome(pre) {
            ntmp := append(tmp, pre)
            RecPal(s[i:], ntmp, res)
        }
    }
}

func partition(s string) [][]string {
    ret := make([][]string, 0)
    tmp := make([]string, 0)
    RecPal(s,tmp,&ret)
    return ret
}


// Python 3
class Solution:
    def palindrome(self, s: str) -> bool:
        return s == s[::-1]

    def partition(self, s: str) -> List[List[str]]:
        dp = [[] for _ in range(len(s) + 1)]
        print(dp)
        dp[-1] = [[]]
        print(dp)
        for i in range(len(s) - 1, -1, -1):
            for j in range(i + 1, len(s) + 1):
                if self.palindrome(s[i:j]):
                    for each in dp[j]:
                        dp[i].append([s[i:j]] + each)
                    print(i,j,dp)
        return dp[0]
