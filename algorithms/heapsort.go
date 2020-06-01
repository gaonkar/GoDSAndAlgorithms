package main

import (
    "fmt"
)

func heapify(arr []int, N int) (r bool){
    r = false
    for i:=1; i < N; i++ {
        c := i
        p := (c - 1)/2  //parent
        for arr[c] < arr[p] {
            arr[c],arr[p] = arr[p], arr[c]
            c = p
            p = (c-1)/2
        r = true
        }
    }
    return r
}

func heapsort(arr []int) {
    heapify(arr, len(arr))
    for i := len(arr)-1; i > 0 ; i-- {
	// move the heap min to its proper position and then re-heapify 0 to i - 1
        arr[0], arr[i] = arr[i], arr[0]
        heapify(arr, i)
    }
}

func main() {
    arr := []int{9,8,7,6,5,4,3,2,1}
    heapsort(arr);
    fmt.Println(arr)
}

