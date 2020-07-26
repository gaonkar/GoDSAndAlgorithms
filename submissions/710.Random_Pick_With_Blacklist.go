/*
Given a blacklist B containing unique integers from [0, N), write a function to return a uniform random integer from [0, N) which is NOT in B.

Optimize it such that it minimizes the call to system’s Math.random().

Note:

1 <= N <= 1000000000
0 <= B.length < min(100000, N)
[0, N) does NOT include N. See interval notation.
Example 1:

Input:
["Solution","pick","pick","pick"]
[[1,[]],[],[],[]]
Output: [null,0,0,0]
Example 2:

Input:
["Solution","pick","pick","pick"]
[[2,[]],[],[],[]]
Output: [null,1,1,1]
Example 3:

Input:
["Solution","pick","pick","pick"]
[[3,[1]],[],[],[]]
Output: [null,0,0,2]
Example 4:

Input:
["Solution","pick","pick","pick"]
[[4,[2]],[],[],[]]
Output: [null,1,3,1]

*/
type WLR struct {
    s, lo, hi int
}
type Solution struct {
    N int
    b *[]int
    w []WLR
}


func Constructor(N int, blacklist []int) Solution {
    var S Solution
    S.N = N - len(blacklist)
    S.b = &blacklist
    L := len(blacklist)
    sort.Ints(blacklist)
    S.w = make([]WLR, 0)
    if L == 0 {
        S.w = append(S.w, WLR{N, 0,N-1})
        return S
    }
    sum := 0
    if blacklist[0] != 0 {
        sum = blacklist[0]
        S.w = append(S.w, WLR{sum, 0,blacklist[0]-1})

    }
    for i:=1; i < len(blacklist); i++ {
        if blacklist[i] == blacklist[i-1] + 1 {continue}
        sum += blacklist[i] - blacklist[i-1]-1
        S.w = append(S.w, WLR{sum, blacklist[i-1]+1,blacklist[i]-1})

    }
    if blacklist[L-1] < N-1 {
        sum += N - blacklist[L-1]-1
        S.w = append(S.w, WLR{sum, blacklist[L-1]+1, N-1})
    }
    fmt.Println(S.w)
    return S
}

func BinarySearch(nums []WLR, target int) int {
    low := 0
    hi := len(nums)
    for low < hi {
        mid := (hi - low) /2 + low
        if nums[mid].s == target {return mid}
        if nums[mid].s > target {
            hi = mid
        } else {
            low = mid + 1
        }

    }
    return low // insert location
}



func (this *Solution) Pick() int {
    N := rand.Intn(this.N)
    /*
    //linear solution 
    for _,x := range(*this.b) {
        if x <= N {
            N++
        } else {
            break
        }
    }
    */
    idx := BinarySearch(this.w, N+1)
    //fmt.Println(idx, N, this.N)
    N = N % (this.w[idx].hi - this.w[idx].lo + 1)
    return this.w[idx].lo + N
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(N, blacklist);
 * param_1 := obj.Pick();
 */
