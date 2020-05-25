/*
 */

var r [][]int
func CHelper(c []int, t int, idx int, tmp []int)  {
    if t <= 0 { //ignore
        if t == 0 { // found a match
            var A []int
            //fmt.Println("X:",tmp)
            for _,x:= range(tmp) {
                A = append(A, x)
            }
            r = append(r, A)
        }
        //fmt.Println("<",r)
        return
    }
    for i := idx ; i < len(c); i++ {
        tmp = append(tmp, c[i])
        CHelper(c,t - c[i], i, tmp)
        tmp = tmp[:len(tmp)-1]
    }
}

func combinationSum(c []int, t int) [][]int {
    var tmp []int
    r = make([][]int, 0)
    CHelper(c,  t, 0, tmp)
    //fmt.Println(r)
    return r
}
