/*
845. Longest Mountain in Array
You may recall that an array arr is a mountain array if and only if:

arr.length >= 3
There exists some index i (0-indexed) with 0 < i < arr.length - 1 such that:
arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
Given an integer array arr, return the length of the longest subarray, which is a mountain. Return 0 if there is no mountain subarray.



Example 1:

Input: arr = [2,1,4,7,3,2,5]
Output: 5
Explanation: The largest mountain is [1,4,7,3,2] which has length 5.
*/

func longestMountain(arr []int) int {
    i, ret := 1, 0
    L := len(arr)
    for i < L {
        for i < L && arr[i-1] == arr[i] {i++}
        up := 0
        //climb up hill
        for i < L && arr[i-1] < arr[i] {
            i++
            up++
        }
        dn := 0
        // climb down hill
        for i < L && arr[i-1] > arr[i] {
            i++
            dn++
        }
        // better have seen up and down hill
        if up == 0 || dn == 0 {continue}
        if up + dn + 1 > ret {ret = up + dn + 1}
    }
    return ret
}
