/*
Given an array of integers, return the maximum sum for a non-empty subarray (contiguous elements) with at most one
element deletion. In other words, you want to choose a subarray and optionally delete one element from it so that there
is still at least one element left and the sum of the remaining elements is maximum possible.

Note that the subarray needs to be non-empty after deleting one element.



Example 1:

Input: arr = [1,-2,0,3]
Output: 4
Explanation: Because we can choose [1, -2, 0, 3] and drop -2, thus the subarray [1, 0, 3] becomes the maximum value.
Example 2:

Input: arr = [1,-2,-2,3]
Output: 3
Explanation: We just choose [3] and it's the maximum sum.
Example 3:

Input: arr = [-1,-1,-1,-1]
Output: -1
Explanation: The final subarray needs to be non-empty.
You can't choose [-1] and delete -1 from it, then get an empty subarray to make the sum equals to 0.


Constraints:

1 <= arr.length <= 105
-104 <= arr[i] <= 104
*/
func Max(x,y int) int {
    if x > y { return x}
    return y
}

func maximumSum(arr []int) int {
    gmax := arr[0]
    nmax := arr[0]
    ignored := 0
    kde := 0 //kadane algorithm
    for _, x := range(arr) {

        // kde is best sum upoto i without ignoring
        // ignored + x is option if we want to continue with past ignored value
        // if kde > ignored + x, then it means time to ignore x, and start with kde
        ignored = Max(kde, ignored+x)

        kde = Max(x, x+kde)
        // pick the maximum from both
        gmax = Max(gmax, ignored)
        gmax = Max(gmax, kde)

        //if only negative numbers, just find the max in them
        nmax = Max(nmax, x)
    }
    if nmax < 0 {return nmax}
    return gmax
}
