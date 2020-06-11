/*
Given n non-negative integers representing the histogram's bar height where the width of each bar is 1, find the area of largest rectangle in the histogram.
*/

func Min(x,y int) int {
    if x < y {return x}
    return y
}

func largestRectangleAreaNSquare(H []int) int {
    M := 0
    for i:=0; i < len(H); i++ {
        C := H[i]
        min := H[i]
        for j := i+1; j < len(H); j++ {
            min = Min(min, H[j])
            N := min * (j - i + 1)
            if N > C { C = N}
        }
        if M < C {M = C}
    }
    return M
}

/*
Think of the stack as tracker that keeps track of the last valid left location where the current bar is minimum. So for any element in the stack at location i, location i-1 to i is the range where it is the minimum.

 */
func largestRectangleAreaStack(H []int) int {
    M := 0
    stack := make([]int, 1)
    S := 1
    CA := 0
    i := 0
    stack[0] = -1
    for i < len(H) {
        if S == 1 || H[stack[S-1]] < H[i] {
            stack = append(stack, i)
            i++
            S++
            continue
        }
        //fmt.Println(stack, S)
        for S > 1 && H[stack[S-1]] >= H[i] {
            p := stack[S-1]
            S--
            CA = (i - 1 - stack[S-1]) * H[p]
            //fmt.Println("CA:", stack, S, stack[S-1],"p", p,  H[p], i, CA)
            if CA > M {M = CA}

        }
        //fmt.Println(stack, S, M)
        stack = stack[:S]
    }
    for S > 1 {
        p := stack[S-1]
        S--
        CA = (i - 1 - stack[S-1]) * H[p]
        if CA > M {M = CA}

    }
    //fmt.Println(stack, S, M)
    return M
}

/*
 3 pass algo
*/
func largestRectangleArea(H []int) int {
    L := make([]int, len(H))
    R := make([]int, len(H))
    B := 0
    L[0] = 0
    for i:=1; i < len(H); i++ {
        B = i-1
        fmt.Println(i, B, L[B], H[B], H[i])
        for B >= 0 && H[B] >= H[i] {
            B = L[B]-1
        }
        L[i] = B+1

    }
    B = 0
    R[len(H)-1] = len(H)-1
    for i:=len(H)-2; i >= 0; i-- {
        B = i+1
        for B < len(H) && H[i] <= H[B] {
            B = R[B]+1
        }
        R[i] = B-1
    }
    fmt.Println(L, R)
    return 0
}
