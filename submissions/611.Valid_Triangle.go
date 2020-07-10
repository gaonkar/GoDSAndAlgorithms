/*
Given an array consists of non-negative integers, your task is to count the number of triplets chosen from the array that can make triangles if we take them as side lengths of a triangle.
Example 1:

Input: [2,2,3,4]
Output: 3
Explanation:
Valid combinations are:
2,3,4 (using the first 2)
2,3,4 (using the second 2)
2,2,3
Note:

The length of the given array won't exceed 1000.
The integers in the given array are in the range of [0, 1000].
*/

/*
 Once we sort them, say [2 3 4 6 7 8]
 we position i at 8, now we need to find all l and h such that nums[l] + nums[h] > 8
 Each time we satisfy this condition, note that for all values from l to h-1 will satisfy this, so add
 r = r + (h-l)
 */

func triangleNumber(nums []int) (r int) {
    L := len(nums)
    sort.Ints(nums)
    //fmt.Println(nums)
    for i:=L-1; i > 1; i-- {
        l := 0
        h := i-1
        for l < h {
            if nums[l] + nums[h] > nums[i] {
                r = r + h - l
                h--
            } else {
                l++
            }
        }
        //fmt.Println(i, r)
    }
    return r
}
