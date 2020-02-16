/*
Given a list of non negative integers, arrange them such that they form the largest number.

Example 1:

Input: [10,2]
Output: "210"
Example 2:

Input: [3,30,34,5,9]
Output: "9534330"
Note: The result may be very large, so you need to return a string instead of an integer.
*/


import "sort"
import "fmt"
// custom interface for sort

type CString []string
func (b CString) Len() int    { return len(b) }
func (b CString) Less(i, j int) bool { return b[i]+b[j] < b[j]+b[i] } // concat ;)
func (b CString) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func largestNumber(nums []int) string {
    var ns CString
    var allzeros bool = true
    out:= ""
    for _, num := range nums {
		ns = append(ns, strconv.Itoa(num))
        if num != 0 {
            allzeros = false
        }
	}
    if allzeros { return "0"}
    sort.Sort(ns)
    //fmt.Println(ns)

    for j:= len(ns)-1; j>=0; j-- {
        out = out + ns[j]
    }

    return out
}
