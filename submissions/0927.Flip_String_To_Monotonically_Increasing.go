/*
A string of '0's and '1's is monotone increasing if it consists of some number of '0's (possibly 0), followed by some number of '1's (also possibly 0.)

We are given a string S of '0's and '1's, and we may flip any '0' to a '1' or a '1' to a '0'.

Return the minimum number of flips to make S monotone increasing.



Example 1:

Input: "00110"
Output: 1
Explanation: We flip the last digit to get 00111.
Example 2:

Input: "010110"
Output: 2
Explanation: We flip to get 011111, or alternatively 000111.
Example 3:

Input: "00011000"
Output: 2
Explanation: We flip to get 00000000.


Note:

1 <= S.length <= 20000
S only consists of '0' and '1' characters.
*/


/*
  Case 1 and 2
    Find number of 0 and 1. Just pick min of those. that will transform the string to
    just 000s or 111s.
  If thats not the best, then
  The string has to be transformed to 0000....01....1111
  We need to find that pivot that minimizes the flips such that everything on the
  left of the pivot is flipped to 0 and everything on right is flipped to 1.
  What we can do is find the flip required if we choose a pivot i and then we
  iterate from 0 to len(S)-1 and choose mininum
*/
func Min(x,y int) int {
    if x < y {return x}
    return y
}

func minFlipsMonoIncr(S string) int {
    c0,c1 := 0,0
    // Find the number of 0
    for _,x := range(S) {
        if x == '0' {
            c0++
        }
    }
    ret := Min(c0,len(S)-c0)
    for _,x := range(S) {
        if x == '0' {
            // we dont need to flip here
            c0--
        } else {
            c1++
        }
        ret = Min(ret, c0+c1)
    }
    return ret
}
