/*
Given a string containing only digits, restore it by returning all possible valid IP address combinations.

A valid IP address consists of exactly four integers (each integer is between 0 and 255) separated by single points.

Example:

Input: "25525511135"
Output: ["255.255.11.135", "255.255.111.35"]
*/
func GetStr(s string, p []int) (r string) {
    r = s[:p[0]]
    for i:=1; i < len(p);i++ {
        r += "." + s[p[i-1]:p[i]]
    }
    return r
}

func ValidInt(s string, b,e int) bool {
    x := 0
    //fmt.Println(s,b,e)
    if e - b > 1 && s[b] == '0' {return false}
    if e - b > 3 {return false}
    for i :=b; i < e; i++ {
        x *=10
        x += int(s[i] - '0')
    }
    //fmt.Println(x,b,e,s[:e])
    return  x <= 255
}

func restoreIpAddresses(s string) (r []string) {
    var helper func(stk int, b int)
    L := len(s)
    stk := make([]int, 4)
    stk[3] = L
    helper = func(stkptr int,b int) {
        if stkptr == 3 {
            if ValidInt(s,b,L) {
                l := GetStr(s, stk)
                r = append(r,l)
            }
            return
        }
        //fmt.Println(stkptr, stk)
        for i:=b+1; i < L; i++ {
            //fmt.Println(stkptr, i)
            if ValidInt(s,b,i) {
                //fmt.Println(b,i)
                stk[stkptr]=i
                helper(stkptr+1,i)
            }
        }

    }
    helper(0,0)
    return r
}
