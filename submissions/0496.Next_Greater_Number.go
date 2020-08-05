/*
You are given two arrays (without duplicates) nums1 and nums2 where nums1’s elements are subset of nums2.
Find all the next greater numbers for nums1's elements in the corresponding places of nums2.

The Next Greater Number of a number x in nums1 is the first greater number to its right in nums2.
If it does not exist, output -1 for this number.

Example 1:

Input: nums1 = [4,1,2], nums2 = [1,3,4,2].
Output: [-1,3,-1]
Explanation:
    For number 4 in the first array, you cannot find the next greater number for it in the second array, so output -1.
    For number 1 in the first array, the next greater number for it in the second array is 3.
    For number 2 in the first array, there is no next greater number for it in the second array, so output -1.
Example 2:

Input: nums1 = [2,4], nums2 = [1,2,3,4].
Output: [3,-1]
Explanation:
    For number 2 in the first array, the next greater number for it in the second array is 3.
    For number 4 in the first array, there is no next greater number for it in the second array, so output -1.
Note:

All elements in nums1 and nums2 are unique.
The length of both nums1 and nums2 would not exceed 1000.
*/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	M := make(map[int]int)
	L := len(nums2)
	if L == 0 {
		return nil
	}
	S := make([]int, 0)
	for _, x := range nums2 {
		for len(S) > 0 && x > S[len(S)-1] {
			M[S[len(S)-1]] = x // set the next greater element, and then pop the stack
			S = S[:len(S)-1]
		}
		S = append(S, x)
	}
	for i := 0; i < len(nums1); i++ {
		v, ok := M[nums1[i]]
		nums1[i] = -1
		if ok {
			nums1[i] = v
		}
	}
	return nums1
} // IDEA:

/*
503 Same problem with a twist
Given a circular array (the next element of the last element is the first element of the array), print the Next Greater Number for every element. The Next Greater Number of a number x is the first greater number to its traversing-order next in the array, which means you could search circularly to find its next greater number. If it doesn't exist, output -1 for this number.

Example 1:
Input: [1,2,1]
Output: [2,-1,2]
Explanation: The first 1's next greater number is 2;
The number 2 can't find next greater number;
The second 1's next greater number needs to search circularly, which is also 2.

There can be duplicate elements. So use the stack to store the index instead of the value.

*/
func nextGreaterElements(nums2 []int) []int {
	M := make(map[int]int)
	L := len(nums2)
	if L == 0 {
		return nil
	}
	S := make([]int, 0)
	for i, _ := range nums2 {
		M[i] = -1
	}
	for i := 0; i < 2*L; i++ {
		x := nums2[i%L]
		for len(S) > 0 && x > nums2[S[len(S)-1]] {
			M[S[len(S)-1]] = x
			S = S[:len(S)-1]
		}
		S = append(S, i%L)
	}

	for i := 0; i < len(nums2); i++ {
		nums2[i] = M[i]
	}
	return nums2
}
