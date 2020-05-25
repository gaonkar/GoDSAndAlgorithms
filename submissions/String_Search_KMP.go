/*
1 2 3 4 5 6 7 8 9
a a b a a b a b b

0 1 0 1 2 3 4 0 0

f - failure function (by definition, this is the length of the longest prefix of the string which is a suffix also)

Here how I built it step by step:
f(a) = 0 (always = 0 for one letter)
f(aa) = 1 (one letter 'a' is both a prefix and suffix)
f(aab) = 0 (there is no the same suffixes and prefixes: a != b, aa != ab)
f(aaba) = 1 ('a' is the same in the beginning and the end, but if you take 2 letters, they won't be equal: aa != ba)
f(aabaa) = 2 ( you can take 'aa' but no more: aab != baa)
f(aabaab) = 3 ( you can take 'aab')
f(aabaaba) = 4 ( you can take 'aaba')
f(aabaabab) = 0 ( 'a' != 'b', 'aa' != 'ab' and so on, it can't be = 5, so as 'aabaa' != 'aabab')
f(aabaababb) = 0 ( the same situation)
*/
func KMPF(p string) []int {
    b :=make([]int, len(p))
    b[0]=0
    i:=1
    L := 0
    for i < len(p) {
        if p[i] == p[L] {
            L++
            b[i] = L
            i++
        } else if L > 0 {
            L = b[L-1]
        } else {
            b[i] = 0
            i++
        }
    }
    return b
}

func strStr(h string, n string) int {
    i := 0
    j := 0
    L := len(h)
    N := len(n)
    if (N==0) {return 0}
    f := KMPF(n)
    for i < L {
        for (i < L && j < N && h[i] == n[j]) {
            i++
            j++
        }
        if j == N { return i-j}
        if j > 0 {
            j = f[j-1]
        } else {
            i++
        }
        //fmt.Println(i,j)
    }
    return -1
}
