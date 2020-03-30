/*
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0).
Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.

Input: [1,8,6,2,5,4,8,3,7]
Output: 49

*/
func Max (x int, y int) int {
    if (x > y) {
        return x
    }
    return y
}
func Min (x int, y int) int {
    if (x < y) {
        return x
    }
    return y
}
func maxArea(height []int) int {
    var l int;
    var m int;
    var MA int = 0;
    var A int;
    l = 0
    m = len(height) - 1
    for l < m {
        A = (m - l) * Min(height[l], height[m])
        MA = Max(MA, A)
        if (height[l] < height[m]) {
            l = l + 1
        } else {
            m = m - 1
        }
    }
    return MA
}
