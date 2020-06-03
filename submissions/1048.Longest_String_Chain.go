/*
Given a list of words, each word consists of English lowercase letters.

Let's say word1 is a predecessor of word2 if and only if we can add exactly one letter anywhere in word1 to make it equal to word2.  For example, "abc" is a predecessor of "abac".

A word chain is a sequence of words [word_1, word_2, ..., word_k] with k >= 1, where word_1 is a predecessor of word_2, word_2 is a predecessor of word_3, and so on.

Return the longest possible length of a word chain with words chosen from the given list of words.



Example 1:

Input: ["a","b","ba","bca","bda","bdca"]
Output: 4
Explanation: one of the longest word chain is "a","ba","bda","bdca".


Note:

1 <= words.length <= 1000
1 <= words[i].length <= 16
words[i] only consists of English lowercase letters.

*/

func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

func Match(s string, t string) bool {
    S := len(s)
    T := len(t)
    if T-S != 1 {return false}

    i := 0

    for i < S && s[i] == t[i] {
        i++
    }
    //fmt.Println(s,t,i, S, T)
    if i == S {
        return true
    }
    //fmt.Println(s[i:], t[i+1:])
    return s[i:] == t[i+1:]
}
func longestStrChain(words []string) int {
    L := len(words)
    if L == 0 {return 0}
    C := make([]int, L) // parent pointer

    sort.Strings(words)
    sort.SliceStable(words, func(i, j int) bool {
        return len(words[i]) >= len(words[j])
    })
    C[0] = 1
    M:=1
    words[0] = SortString(words[0])
    max := 1
    for i := 1;i < L; i++ {
        C[i] = 1
        M = 0
        words[i] = SortString(words[i])
        for j:=0; j < i; j++ {
            if M < C[j] && Match(words[i], words[j]) {
                M = C[j]
            }
        }
        C[i] += M
        if max < C[i] { max = C[i]}
    }
    //fmt.Println(words, C)
    return max
}
