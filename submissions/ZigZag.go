/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
(you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I
*/


func convert(s string, numRows int) string {
    a := make([][]rune, numRows)

    if numRows == 1 {
        return s
    }
    for i := range a {
        a[i] = make([]rune, 0)
    }
    r := 0
    dir := true
    for _,x := range s {
        //fmt.Println(x)
        a[r] = append(a[r], x)
        if dir {
            r++
            if (r == numRows) {
                r = r - 2
                dir = false
            }
        } else {
            r--
            if (r < 0) {
                r = 1
                dir = true
            }
        }
    }
    out := ""
    for i := range a {
        s = string(a[i])
        out = out + s
    }
    return out
}
