

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

*/

import (
    "fmt"
    "os"
)
func twoSum(nums []int, target int) []int {
    var m map[int]int
    var idx map[int]int
    var ret = make([]int, 2)
    m = make(map[int]int)
    idx = make(map[int]int)
    for i, n:= range nums {
        var diff = target - n
        //fmt.Println(i, n, diff, m[diff])
        if m[diff]>0 {
            ret[0] = idx[diff]
            ret[1] = i
            fmt.Println(ret)
            return ret
        }
        m[n] = 1
        idx[n] = i
    }
    return make([]int, 2)
}

