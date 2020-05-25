/*
76. Minimum Window Substring
Hard

4089

284

Add to List

Share
Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

Example:

Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"
Note:

If there is no such window in S that covers all characters in T, return the empty string "".
If there is such window, you are guaranteed that there will always be only one unique minimum window in S.
 */
func minWindow(s string, t string) (res string) {
    T := make(map[byte] int)
    S := make(map[byte] int)
    var pos []int
    //  Create a hashtable of T string that counts occurance of each byte
    for i:=0; i < len(t);i++ {
        T[t[i]]++
    }
    // Note the length of the hashtable (uniq char count)
    Tn := len(T)
    Ts := 0
    l := 0      // back pointer
    h := 0      // front pointer
    res = ""
    ok := false
    M := 2*len(s) // just to catch the base case
    for  h < len(s) {
        _,ok = T[s[h]]
        if false == ok { // ignore characters not in the T string
            h++
            continue
        }
        S[s[h]]++ // add it to the S hashtable
        pos = append(pos, h) // keep track of the position we found it, FIFO queue
        if T[s[h]] == S[s[h]]{
            Ts++ // track the uniq characters in the S hashtable
            //fmt.Println("INC")
        }

        l = pos[0]      // lets start moving the backpointer to see if we can do better (shorter string)
        for l <= h && Ts == Tn {
            if M > h-l + 1 { // we found a string that matches the requirement thats smaller in length
                M = h-l+1
                res =s[l:h+1]
            }

            S[s[l]]--   // remove that
            if T[s[l]] > S[s[l]]{
                Ts--
            }
            pos = pos[1:]   //pop the oldest
            if len(pos) == 0 {break}
            l = pos[0]      // move to next character
        }
        h++
    }
    return res
}
