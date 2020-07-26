/*
Implement FreqStack, a class which simulates the operation of a stack-like data structure.

FreqStack has two functions:

push(int x), which pushes an integer x onto the stack.
pop(), which removes and returns the most frequent element in the stack.
If there is a tie for most frequent element, the element closest to the top of the stack is removed and returned.


Example 1:

Input:
["FreqStack","push","push","push","push","push","push","pop","pop","pop","pop"],
[[],[5],[7],[5],[7],[4],[5],[],[],[],[]]
Output: [null,null,null,null,null,null,null,5,7,5,4]
Explanation:
After making six .push operations, the stack is [5,7,5,7,4,5] from bottom to top.  Then:

pop() -> returns 5, as 5 is the most frequent.
The stack becomes [5,7,5,7,4].

pop() -> returns 7, as 5 and 7 is the most frequent, but 7 is closest to the top.
The stack becomes [5,7,5,4].

pop() -> returns 5.
The stack becomes [5,7,4].

pop() -> returns 4.
The stack becomes [5,7].


Note:

Calls to FreqStack.push(int x) will be such that 0 <= x <= 10^9.
It is guaranteed that FreqStack.pop() won't be called if the stack has zero elements.
The total number of FreqStack.push calls will not exceed 10000 in a single test case.
The total number of FreqStack.pop calls will not exceed 10000 in a single test case.
The total number of FreqStack.push and FreqStack.pop calls will not exceed 150000 across all test cases.
*/
type FreqStack struct {
    max int
    freq [][]int
    fmap map[int] int
}


func Constructor() FreqStack {
    var F FreqStack
    F.max = 0
    F.freq = make([][]int, 0)
    F.fmap = make(map[int] int)
    return F
}

func Max(x,y int) int {
    if x < y {return y}
    return x
}


func (this *FreqStack) Push(x int)  {
    v, _ := this.fmap[x]
    if v == this.max {
        this.max = v + 1
        this.freq = append(this.freq, []int{})
    }
    this.fmap[x] = v + 1
    this.freq[v] = append(this.freq[v], x)
}


func (this *FreqStack) Pop() int {
    idx := this.max - 1
    ln := len(this.freq[idx])
    ret := this.freq[idx][ln-1]
    if ln == 1 {
        this.max--
        this.freq = this.freq[:this.max]
    } else {
        this.freq[idx] = this.freq[idx][:ln-1]
    }
    this.fmap[ret]--
    return ret
}
