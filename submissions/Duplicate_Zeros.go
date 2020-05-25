/*
Given a fixed length array arr of integers, duplicate each occurrence of zero, shifting the remaining elements to the right.

Note that elements beyond the length of the original array are not written.

Do the above modifications to the input array in place, do not return anything from your function.



Example 1:

Input: [1,0,2,3,0,4,5,0]
Output: null
Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]
Example 2:

Input: [1,2,3]
Output: null
Explanation: After calling your function, the input array is modified to: [1,2,3]


Note:

1 <= arr.length <= 10000
0 <= arr[i] <= 9



*/


func duplicateZeros(arr []int)  {
    Z := 0
    L := len(arr)
    i :=0
    p :=0
    F :=-1

    // Count the number of original 0 that will remain as they occupy 2 spaces each
    // Find the first 0 position, as we need to move only that data
    for i=0; p < L; i++{
        if arr[i] == 0 {
            if F == -1 { F = i}
            Z++
            p+=2
        } else {
            p++
        }
    }
    i = i - 1
    p--
    // there are no 0s
    if (F < 0) {return }
    //fmt.Println("i=", i,"p=", p, F,"len=",L, "a[i]=",arr[i],  arr)

    for p >= F && i >= F{
        //fmt.Println(i, p, F, arr[i:])
        if arr[i] == 0 {
            if p < L { //the element i could be out of bounds
                arr[p] = 0
            }
            p--
        }
        //fmt.Println("Z", i, p, F, arr[i:])
        if p < F {break}
        if (p < L){
            arr[p] = arr[i]
        }
        i--
        p--
        //fmt.Println("i=", i,"p=", p, F,"len=",L, arr)
    }
    //fmt.Println(i, Z,L, arr)

}
