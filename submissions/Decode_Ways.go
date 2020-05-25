/*
A message containing letters from A-Z is being encoded to numbers using the following mapping:

'A' -> 1
'B' -> 2
...
'Z' -> 26
Given a non-empty string containing only digits, determine the total number of ways to decode it.

Example 1:

Input: "12"
Output: 2
Explanation: It could be decoded as "AB" (1 2) or "L" (12).
Example 2:

Input: "226"
Output: 3
Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
*/

func numDecodings(s string) int {
    arr := make([]int, 0)
    L := len(s)
    if L == 0 {return 0}
    if s[0] == '0' {
        arr = append(arr, 0)
        arr = append(arr, 0)
    } else {
        arr = append(arr, 1)
        arr = append(arr, 1)
    }
    for i:=1 ; i < L; i++ {
        a := 0
        if s[i] != '0' {
            a = arr[i]
        }

        if(s[i-1]=='1'||(s[i]<'7'&& s[i-1]=='2')) {
            a += arr[i-1]
        }
        arr = append(arr, a)
    }
    fmt.Println(arr)
    return arr[len(arr)-1]
}:
