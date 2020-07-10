/*
Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.

Note:

The length of both num1 and num2 is < 5100.
Both num1 and num2 contains only digits 0-9.
Both num1 and num2 does not contain any leading zero.
You must not use any built-in BigInteger library or convert the inputs to integer directly.
*/
func Reverse(r []byte) {
    n := len(r) - 1
    for i:=0; i < len(r)/2;i++ {
        r[i], r[n] = r[n],r[i]
        n--
    }
}

func Add (x, y byte, R int) (s byte, r int) {
    a := int(x - '0')
    b := int(y - '0')
    c := a + b + R
    r = c /10
    s = byte('0' + c % 10)
    return s, r
}

func GetByte(A []byte, i int) byte {
    if i >= len(A) {return '0'}
    return A[i]
}

func addStrings(num1 string, num2 string) string {
    N1 := []byte(num1)
    N2 := []byte(num2)
    Reverse(N1)
    Reverse(N2)
    S := make([]byte, 0)
    s := byte('0')
    n1, n2, r := 0, 0, 0
    for n1 < len(N1) ||  n2 < len(N2) {
        a := GetByte(N1, n1)
        b := GetByte(N2, n2)
        s, r = Add(a,b,r)
        S = append(S, s)
        n1++
        n2++
    }
    if r > 0 {
        s, r = Add('0','0',r)
        S = append(S, s)
    }
    Reverse(S)
    return string(S)
}
