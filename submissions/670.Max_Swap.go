/*
Given a non-negative integer, you could swap two digits at most once to get the
maximum valued number. Return the maximum valued number you could get.

Example 1:
Input: 2736
Output: 7236
Explanation: Swap the number 2 and the number 7.
Example 2:
Input: 9973
Output: 9973
Explanation: No swap.
Note:
The given number is in the range [0, 108]
*/

func Reverse(r []int) {
    n := len(r)-1
    for i := 0; i < n;i++ {
        r[i], r[n] = r[n], r[i]
        n--
    }
}
func BreakInt(x int) (r []int) {
    for x > 0 {
        r = append(r, x % 10)
        x = x / 10
    }
    Reverse(r)
    return r
}

func MakeInt(r []int) (x int) {
    for _, n:= range(r) {
        x = x * 10 + n
    }
    return x
}

func Find1(arr []int) int {
    i := 0
    for i < len(arr)-1{
        for j := i+1; j < len(arr);j++ {
            if arr[i] < arr[j] {return i}
        }
        i++
    }
    return -1
}

func findLastMax(arr []int, idx int) int {
    j := len(arr) - 1
    mpos := j
    mval := arr[j]
    j--
    for j > idx {
        if mval < arr[j] {
            mpos = j
            mval = arr[j]
        }
        j--
    }
    return mpos
}

func maximumSwap(num int) int {
    arr := BreakInt(num)
    fmt.Println(num, arr)
    // find the first position that can be swapped
    pos1 := Find1(arr)
    if pos1 < 0 {return num}
    pos2 := findLastMax(arr, pos1)
    //fmt.Println(pos1, arr, pos2)
    arr[pos1], arr[pos2] = arr[pos2], arr[pos1]
    return MakeInt(arr)
}
