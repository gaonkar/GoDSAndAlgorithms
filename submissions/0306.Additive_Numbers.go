/*
Additive number is a string whose digits can form additive sequence.

A valid additive sequence should contain at least three numbers. Except for the first two numbers, each subsequent number in the sequence must be the sum of the preceding two.

Given a string containing only digits '0'-'9', write a function to determine if it's an additive number.

Note: Numbers in the additive sequence cannot have leading zeros, so sequence 1, 2, 03 or 1, 02, 3 is invalid.



Example 1:

Input: "112358"
Output: true
Explanation: The digits can form an additive sequence: 1, 1, 2, 3, 5, 8.
             1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8
Example 2:

Input: "199100199"
Output: true
Explanation: The additive sequence is: 1, 99, 100, 199.
             1 + 99 = 100, 99 + 100 = 199


Constraints:

num consists only of digits '0'-'9'.
1 <= num.length <= 35
*/

func VS(s string) bool {
    if len(s) == 1 {return true}
    return s[0] != '0'
}
func IsValidIntSeq(f,s int, r string) bool {
    if !VS(r) && f+s != 0 {return false}
    total := f + s
    //fmt.Println(f,s, r, len(r))
    flag := true
    Rlen := len(r)
    for i:=1; i <= Rlen && flag;i++ {
        z := StoI(r[:i])
        if (total == z) {
            if (i == Rlen) {return true}
            return IsValidIntSeq(s,z, r[i:])
        } else if total < z || (i == Rlen) {
            flag = false
        }
    }
    return flag
}

func StoI(s string) int {
    x, err := strconv.Atoi(s)
    if err == nil {
        return x
    }
    return 0
}

func IsValidSeq(f,s,r string) bool {
    if (!VS(f)||!VS(s)) {return false}
    if len(f) > len(r) || len(s) > len(r) {return false}
    x := StoI(f)
    y := StoI(s)
    //fmt.Println(f,s,r)
    return IsValidIntSeq(x,y, r)
}

func isAdditiveNumber(num string) bool {
    L := len(num)
    if (L <= 2) {return false}
    H := (2*L)/3 + L%2 // just an optimization to just use upto 2/3+1 digits
    for i:=1; i <= H; i++ {
        for j:=i+1; j <= H; j++ {
            if (IsValidSeq(num[:i], num[i:j], num[j:])) {
                return true
            }
        }
    }
    return false
}
