/*
Given an array A of positive integers (not necessarily distinct), return the
lexicographically largest permutation that is smaller than A, that can be made
with one swap (A swap exchanges the positions of two numbers A[i] and A[j]).
If it cannot be done, then return the same array.

Example 1:

Input: [3,2,1]
Output: [3,1,2]
Explanation: Swapping 2 and 1.
Example 2:

Input: [1,1,5]
Output: [1,1,5]
Explanation: This is already the smallest permutation.
Example 3:

Input: [1,9,4,6,7]
Output: [1,7,4,6,9]
Explanation: Swapping 9 and 7.
Example 4:

Input: [3,1,1,3]
Output: [1,3,1,3]
Explanation: Swapping 1 and 3.


Note:

1 <= A.length <= 10000
1 <= A[i] <= 10000

Its more like a brain teaser.
Refer to 31. Next Permutation

*/
func Find1(arr []int) int {
	i := 0
	// Find the first position starting from N where A[i] > A[i+1]
	for i = len(arr) - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			return i
		}
	}
	return i
}

func Max(A []int, idx int) int {
	// find the largest numbers in A[pos1:] such that it is smaller than A[pos1] and the first one
	midx := idx + 1
	max := A[midx]

	for i := idx + 2; i < len(A); i++ {
		// Skip number that are equal to pos1
		if A[i] >= A[idx] {
			continue
		}
		if max < A[i] {
			midx = i
			max = A[i]
		}
	}
	return midx
}

func prevPermOpt1(A []int) []int {
	// Find the first position where A[i] > any a[i+1]
	pos1 := Find1(A)
	if pos1 < 0 {
		return A
	}
	pos2 := Max(A, pos1)
	//fmt.Println(pos1, A[pos1], pos2, A[pos2])
	A[pos1], A[pos2] = A[pos2], A[pos1]
	return A
}
