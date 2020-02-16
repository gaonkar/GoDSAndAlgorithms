/*
Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

Note:

The length of both num1 and num2 is < 110.
Both num1 and num2 contain only digits 0-9.
Both num1 and num2 do not contain any leading zero, except the number 0 itself.
You must not use any built-in BigInteger library or convert the inputs to integer directly.

*/


func ArrToString(arr []int) string {
    str := ""
    for i:= range arr {
        str = string(arr[i] + '0') + str
    }
    return str
}

func StringToArr(s string) [] int {
    var arr []int
    for i:= len(s)-1; i >= 0; i-- {
        arr = append(arr, int(s[i]) - int('0'))
    }
    return arr
}



func MultDigit(n1 []int, d int, n int) [] int {
    sarr := make([]int, n)
    c := 0
    if d == 0 {
        return make([]int, 1)
    }

    for i := 0; i < len(n1); i++ {
        p := n1[i] * d + c
        s := p % 10
        c = p / 10
        sarr = append(sarr,s)
    }
    if (c > 0) {
        sarr = append(sarr, c)
    }
    return sarr
}

func SumArray(b []int, m []int) [] int {
    var arr []int
    i:= 0
    c:= 0
    for i < len(m) {
        p:= b[i] + m[i] + c
        s:= p % 10
        c = p / 10
        arr = append(arr, s)
        i++
    }
    for i < len(b) {
        p:= b[i] + c
        s:= p % 10
        c = p / 10
        arr = append(arr, s)
        i++
    }
    if (c > 0) {
        arr = append(arr, c)
    }
    return arr
}

func SumNumbers(narr1 []int, narr2[]int) [] int {

    N1:= len(narr1)
    N2:= len(narr2)
    if N1 == 0 {
        return narr2
    }
    if N2 == 0 {
        return narr1
    }
    if (N1 > N2) {
        return SumArray(narr1, narr2)
    }
    return SumArray(narr2,narr1)
}

func multiply(num1 string, num2 string) string {

    sarr := make([]int, 10+len(num1) + len(num2))
    narr1 := StringToArr(num1)
    narr2 := StringToArr(num2)
    //fmt.Println(narr1, narr2)
    for i:= range(num1) {
        s:= MultDigit(narr2, narr1[i], i)
        sarr = SumNumbers(sarr, s)
        //fmt.Println(sarr, s)
    }

    s:=ArrToString(sarr)
    i:=0
    for i < len(s)-1 && s[i] == '0' {
        i++
    }
    return s[i:]
}
