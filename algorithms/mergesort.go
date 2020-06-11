package main

import (
	"fmt"
)

type A struct {
    V int       // value
    I int       // index
}


func merge(arr, buf []A, l, m, r int, inv []int) (ret int){
    copy(buf[l:r],arr[l:r])
    k := l
    i := l
    j := m
    ret = 0
    //The move tracks how many elements from arr[m:r] moves to arr[l:m].
    move := 0
    for i < m && j < r {
        if buf[i].V <= buf[j].V {
            arr[k] = buf[i]
            i++
            // the current move for this element in this iteration is fixed, so record it
            inv[arr[k].I] += move
        } else {
            arr[k] = buf[j]
            ret += (j - k)
            // element from arr[m:r] is moving so for all elements remaining in i:m, we need to add move
            move++
            j++
        }
        k++
    }
    for i < m {
        arr[k] = buf[i]
        // means we have already moved all elements of arr[m:r] so add that these are numbers smaller for all the
        // elements being copied from arr[l:m]. you can assert move == (l-m)
        inv[arr[k].I] +=move
        i++
        k++
    }
    for j < r {
        arr[k] = buf[j]
        j++
        k++
    }
    return ret
}


func InvCount(arr, ver []int) int {
    inv_count := 0
    for i := 0; i < len(arr)-1; i++ {
        for j := i + 1; j < len(arr); j++ {
            if arr[i] > arr[j] {
                inv_count++
                ver[i]++
	    }
        }
    }
    return inv_count
}


/*
    buf -  temp buffer for merging
    l   -   left
    r   -   right
*/
func mergesort(arr, buf []A, l int, r int, inv []int) int{
    if r - l <= 1 {return 0}
    m := l + (r - l)/2
    ret := mergesort(arr, buf, l, m, inv)
    ret += mergesort(arr, buf, m, r, inv)
    ret += merge(arr, buf, l, m, r, inv)
    return ret
}
func main() {
	arr := []int{1,4,2,3, 9,7,5,4,3,2,3,4,6,7,8,5,4,3,2}
        L := len(arr)
        inv := make([]int, L)
        ver := make([]int, L)
        av := make([]A,0)
        for i:=0; i < L; i++ {
            av = append(av, A{arr[i], i})
        }
        buf := make([]A, L)
        fmt.Println("Orig:", arr)
        fmt.Println("Inv count:", InvCount(arr, ver), mergesort(av, buf, 0, L, inv))
        fmt.Println("Inv Arr", inv, ver)
        fmt.Println("Sorted:", av)
}
