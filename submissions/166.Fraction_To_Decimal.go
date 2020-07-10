/*
iven two integers representing the numerator and denominator of a fraction, return the fraction in string format.

If the fractional part is repeating, enclose the repeating part in parentheses.

Example 1:

Input: numerator = 1, denominator = 2
Output: "0.5"
Example 2:

Input: numerator = 2, denominator = 1
Output: "2"
Example 3:

Input: numerator = 2, denominator = 3
Output: "0.(6)"
*/
func Abs(x int) (int, int) {
    if x < 0 {
        return -1, -x
    }
    return 1, x
}

func fractionToDecimal(numerator int, denominator int) (r string) {

    n, N :=Abs(numerator)
    d, D :=Abs(denominator)
    M := make(map[int] int)
    if n * d < 0 {
        r = "-"
    } else {
        r = ""
    }
    if N == 0 {return "0"}
    r += strconv.Itoa(N/D)
    if (N%D==0){return r}
    r += "."
    dec := ""
    div := N % D
    i := 0
    for true {
        //fmt.Println(div,D, dec, M)

        div = div * 10
        if M[div] > 0 {
            break
        }
        if div % D == 0 {
            dec += strconv.Itoa(div/D)
            return r + dec
        }
        dec += strconv.Itoa(div/D)
        M[div]=i+1
        div = div % D
        i++

    }
    //fmt.Println(M[div], div, M)
    p := M[div] - 1
    dec = dec[:p] + "(" + dec[p:] + ")"
    r = r + dec
    return r
}
