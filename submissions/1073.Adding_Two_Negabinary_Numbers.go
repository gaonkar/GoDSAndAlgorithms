/*
Given two numbers arr1 and arr2 in base -2, return the result of adding them together.
Each number is given in array format:  as an array of 0s and 1s, from most significant bit to least significant bit.
For example, arr = [1,1,0,1] represents the number (-2)^3 + (-2)^2 + (-2)^0 = -3.  A number arr in array, format is also
guaranteed to have no leading zeros: either arr == [0] or arr[0] == 1.
Return the result of adding arr1 and arr2 in the same format: as an array of 0s and 1s with no leading zeros.



Example 1:

Input: arr1 = [1,1,1,1,1], arr2 = [1,0,1]
Output: [1,0,0,0,0]
Explanation: arr1 represents 11, arr2 represents 5, the output represents 16.
Example 2:

Input: arr1 = [0], arr2 = [0]
Output: [0]
Example 3:

Input: arr1 = [0], arr2 = [1]
Output: [1]


Constraints:

1 <= arr1.length, arr2.length <= 1000
arr1[i] and arr2[i] are 0 or 1
arr1 and arr2 have no leading zeros
*/

func Reverse(r []int) {
    n := len(r)-1
    for i := 0; i < n;i++ {
        r[i], r[n] = r[n], r[i]
        n--
    }
}

func Compose(x int) (v,c int) {
    if x >= 2 {
        // we want to set 2 or 3 at position i, the carry has to be -1 (-2)^i + (-2)^i = -1 * (-2)^(i+1)
        // and value is x % 2
        return (x%2), -1
    }
    if x == -1 {
        // if -1 to be set a position i, we convert that to 1 at position i and 1 at carry
        // (-2)^i + (-2)^(i+1) = -1 * (-2)^(i)
        return 1, 1
    }
    // it is 1 to set a position 1, carry is 0
    return x, 0
}

func Val(A []int, i int) int {
    if i >= len(A) {return 0}
    return A[i]
}

func addNegabinary(arr1 []int, arr2 []int) []int {
    Reverse(arr1)
    Reverse(arr2)
    l1, l2 := 0, 0
    ret := make([]int, 0)
    sum, carry, i := 0, 0, 0
    for l1 < len(arr1) || l2 < len(arr2) || carry != 0 {
        sum = Val(arr1, l1) + Val(arr2, l2) + carry
        i, carry = Compose(sum)
        ret = append(ret, i)
        l1++
        l2++
    }
    Reverse(ret)
    l1 = 0
    for l1 < len(ret)-1 && ret[l1] == 0 {l1++}
    return ret[l1:]
}
