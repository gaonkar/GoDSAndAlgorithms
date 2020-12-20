/*
An encoded string S is given.  To find and write the decoded string to a tape, the encoded string is read one character
at a time and the following steps are taken:
If the character read is a letter, that letter is written onto the tape.
If the character read is a digit (say d), the entire current tape is repeatedly written d-1 more times in total.
Now for some encoded string S, and an index K, find and return the K-th letter (1 indexed) in the decoded string.



Example 1:

Input: S = "leet2code3", K = 10
Output: "o"
Explanation:
The decoded string is "leetleetcodeleetleetcodeleetleetcode".
The 10th letter in the string is "o".
Example 2:

Input: S = "ha22", K = 5
Output: "h"
Explanation:
The decoded string is "hahahaha".  The 5th letter is "h".
Example 3:

Input: S = "a2345678999999999999999", K = 1
Output: "a"
Explanation:
The decoded string is "a" repeated 8301530446056247680 times.  The 1st letter is "a".


Constraints:

2 <= S.length <= 100
S will only contain lowercase letters and digits 2 through 9.
S starts with a letter.
1 <= K <= 10^9
It's guaranteed that K is less than or equal to the length of the decoded string.
The decoded string is guaranteed to have less than 2^63 letters.
*/

/*
  **recursive

  Note that K indes from 1, not 0.
  Assume we had a string xyz3 and K = 4, this string would expand to xyzxyzxyz and the 4th character is x.
  That is same as K % len(xyz) = 1. That means that if K < expanded length, it is always one of the character.
  If the string was xyz33 and K = 10, The first expansion makes it 9, the second would make it 27. After the first
  expansion, slentmp = 27, so we recursively find a new K' thats smaller than K,
  K = (K-1) % 9 + 1 = 1

  Time complexity is NLog(N)
 */
func decodeAtIndex(S string, K int) string {
    var DAI func(K uint64) string

    DAI = func(K uint64) string {
        var slen uint64
        fmt.Println(K)
        for i:=0; i < len(S); i++ {
            if 'a' <= S[i] && S[i] <= 'z' {
                if slen == (K - 1) {return string(S[i])}
                slen++
                continue
            }
            slentmp := uint64(S[i]-'0') * slen
            if slentmp >= K {

                return DAI((K-1)%slen+1)
            }
            slen = slentmp
        }
        return ""
    }
    return DAI(uint64(K))
}

// stack
func isDigit(x byte) bool {
    if '0' <= x && x <= '9' {return true}
    return false
}
func decodeAtIndex(S string, K int) string {
    stk := make([]byte, 0)
    var slen uint64

    for i:=0; i < len(S); i++ {
        if !isDigit(S[i]) {
            slen++
        } else {
            slen *= uint64(S[i]-'0')
        }
        stk = append(stk, S[i])
        for slen >= uint64(K) {
            for isDigit(stk[len(stk)-1]) {
                slen /= uint64(stk[len(stk)-1]-'0')
                stk = stk[:len(stk)-1]
            }
            K = K % int(slen)
            if K == 0 { return string(stk[len(stk)-1])}
            slen--
            stk = stk[:len(stk)-1]
        }
    }
    return ""
}
