/*
139. Word Break
Medium

3925

212

Add to List

Share
Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, determine if s can be segmented into a space-separated sequence of one or more dictionary words.

Note:

The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.
Example 1:

Input: s = "leetcode", wordDict = ["leet", "code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".
Example 2:

Input: s = "applepenapple", wordDict = ["apple", "pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
             Note that you are allowed to reuse a dictionary word.
Example 3:

Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
Output: false
*/



func wordBreak(s string, wordDict []string) bool {
    M := make(map[string] bool)
    L := len(s)
    D := make([][]bool, L)
    for _,w := range(wordDict) {
        M[w] = true
    }
    for j := range(s) {
        D[j] = make([]bool, L+1)
        if M[string(s[j])] {
            D[j][j+1] = true // the off by 0 requires us to not worry so we represent 1 char s[0:1]
        }
    }
    for l := 2; l <= L; l++ {
        for i:=0; i < L - l+1; i++ {
            h := i + l
            str := s[i:h]
            //fmt.Println(str, i, h)
            if M[str] {
                D[i][h] = true		// We found the string directly in the between S[i:h]
                continue
            }
	    // Now for any position k  between i and h, check if we already found i:k, k+1:i in the D array
            for k := i+1; k < h; k++ {
                if D[i][k] && D[k][h] {
                    D[i][h] = true
                    break
                }
            }
        }
    }
    return D[0][L]
}
