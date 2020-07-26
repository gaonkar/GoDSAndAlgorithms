/*
Given an integer n, return the number of strings of length n that consist only of vowels (a, e, i, o, u) and are lexicographically sorted.

A string s is lexicographically sorted if for all valid i, s[i] is the same as or comes before s[i+1] in the alphabet.



Example 1:

Input: n = 1
Output: 5
Explanation: The 5 sorted strings that consist of vowels only are ["a","e","i","o","u"].
Example 2:

Input: n = 2
Output: 15
Explanation: The 15 sorted strings that consist of vowels only are
["aa","ae","ai","ao","au","ee","ei","eo","eu","ii","io","iu","oo","ou","uu"].
Note that "ea" is not a valid string since 'e' comes after 'a' in the alphabet.
Example 3:

Input: n = 33
Output: 66045


Constraints:

1 <= n <= 50
*/

// recursive
func CV(l, n int) int {
    if n == 1 {
        return 5 - l + 1
    }
    // repeat the same character
    c := 0
    for l <= 5 {
        c += CV(l,n-1)
        l++
    }
    return c
}
func countVowelStrings(n int) int {
    return CV(1,n)
}

// bottom up DP
def countVowelStrings(self, n: int) -> int:
     dp = [1] * 5
     for i in range(n):
         dp = accumulate(dp) //https://docs.python.org/3/library/itertools.html
     return list(dp)[-1]


// some smart people have a solution (n+1)C4 :)
