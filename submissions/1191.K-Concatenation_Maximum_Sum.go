/*
Given an integer array arr and an integer k, modify the array by repeating it k times.

For example, if arr = [1, 2] and k = 3 then the modified array will be [1, 2, 1, 2, 1, 2].

Return the maximum sub-array sum in the modified array. Note that the length of the sub-array can be 0 and its sum in that case is 0.

As the answer can be very large, return the answer modulo 10^9 + 7.



Example 1:

Input: arr = [1,2], k = 3
Output: 9
Example 2:

Input: arr = [1,-2,1], k = 5
Output: 2
Example 3:

Input: arr = [-1,-2], k = 7
Output: 0


Constraints:

1 <= arr.length <= 10^5
1 <= k <= 10^5
-10^4 <= arr[i] <= 10^4
*/
func Max(x,y int) int {
    if x > y { return x}
    return y
}
var MOD int

func Kadane(arr []int, lmax int) (int,int) {
    gmax :=  arr[0]
    for _, x := range(arr) {
        lmax = Max(x, (lmax + x))
        gmax = Max(lmax, gmax)
    }
    return gmax%MOD, lmax
}

func Sum(arr []int) int {
    ret := 0
    for _, x := range(arr) {
        ret += x
    }
    return ret
}

func kConcatenationMaxSum(arr []int, k int) int {
    MOD = 1000000009
    if len(arr) == 0 {return 0}
    kdn, lmax := Kadane(arr, 0)
    tot := Sum(arr)
    if k == 1 {
        return Max(tot, kdn)
    }
    //run it as if the array are now concatenated
    kdn, _ = Kadane(arr, lmax)
    if k == 2 {
        return Max(tot, kdn)
    }
    ret := Max(0, kdn)
    if tot < 0 {return ret}
    //each time we go over a cycle, we can add up the total
    ret =  (tot * (k-2)  + kdn) %  MOD
    return ret
}
