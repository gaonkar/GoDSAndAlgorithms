/*
There are two sorted arrays nums1 and nums2 of size m and n respectively.
Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
You may assume nums1 and nums2 cannot be both empty.

02/13/2020
Runtime: 12 ms, faster than 88.75% of Go online submissions for Median of Two Sorted Arrays.
Memory Usage: 5.6 MB, less than 65.00% of Go online submissions for Median of Two Sorted Arrays.

*/
import "fmt"


func Median(arr []int) float64 {
    ln := len(arr)
    if ln == 0 {
        return 0
    }
    if ln% 2 == 0 {
        return float64(arr[ln/2] + arr[ln/2-1])/2.0
    }
    return float64(arr[ln/2])
}
func Min(x, y int) int{
    if x > y {
        return y
    }
    return x
}

func Max(x, y int) int{
    if x < y {
        return y
    }
    return x
}

func fMed(barr []int, sarr []int) float64 {
    var  b, a int = 0, 0
    //fmt.Println("b:", barr)
    //fmt.Println("s:", sarr)

    t:= len(sarr) + len(barr)
    sh:=len(sarr)-1
    bh := (len(barr) - len(sarr))/2
    //fmt.Println("bh:",bh,"sh:",sh, barr[bh], sarr[sh])
    for (sh >= 0 && barr[bh] < sarr[sh]) {
        bh++
        sh--
    }
    //fmt.Println("bh:",bh,"sh:",sh)
    if (t%2 == 1) {
        if sh < len(sarr) -1 {
            return float64(Min(sarr[sh+1],barr[bh]))
        } else {
            return float64(barr[bh])
        }
    } else {
        if sh < len(sarr) - 1 && bh < len(barr) {
            b = Min(sarr[sh+1], barr[bh])
        } else if bh == len(barr) {
            b = sarr[sh + 1]
        } else {
            b = barr[bh]
        }
        if sh >= 0  && bh > 0{
            a = Max(sarr[sh],barr[bh-1])
        } else if bh == 0{
            a = sarr[sh]
        } else {
            a = barr[bh-1]
        }
        return float64((a+b))/2
    }
    return 0.0;
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums1) == 0 {
        return Median(nums2)
    } else if len(nums2) == 0 {
        return Median(nums1)
    } else if len(nums1) > len(nums2) {
        return fMed(nums1,nums2)
    }
    return fMed(nums2, nums1)
}

