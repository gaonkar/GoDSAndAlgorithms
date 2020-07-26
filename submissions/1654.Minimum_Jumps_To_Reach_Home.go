/*
A certain bug's home is on the x-axis at position x. Help them get there from position 0.

The bug jumps according to the following rules:

It can jump exactly a positions forward (to the right).
It can jump exactly b positions backward (to the left).
It cannot jump backward twice in a row.
It cannot jump to any forbidden positions.
The bug may jump forward beyond its home, but it cannot jump to positions numbered with negative integers.

Given an array of integers forbidden, where forbidden[i] means that the bug cannot jump to the position forbidden[i],
and integers a, b, and x, return the minimum number of jumps needed for the bug to reach its home. If there is no
possible sequence of jumps that lands the bug on position x, return -1.



Example 1:

Input: forbidden = [14,4,18,1,15], a = 3, b = 15, x = 9
Output: 3
Explanation: 3 jumps forward (0 -> 3 -> 6 -> 9) will get the bug home.
Example 2:

Input: forbidden = [8,3,16,6,12,20], a = 15, b = 13, x = 11
Output: -1
Example 3:

Input: forbidden = [1,6,2,14,5,17,4], a = 16, b = 9, x = 7
Output: 2
Explanation: One jump forward (0 -> 16) then one jump backward (16 -> 7) will get the bug home.


Constraints:

1 <= forbidden.length <= 1000
1 <= a, b, forbidden[i] <= 2000
0 <= x <= 2000
All the elements in forbidden are distinct.
Position x is not forbidden.
*/


/*
    The crux of the idea is the we are only allowed to jump backwards 1 time.
    Lets say we are at position i after jumping form i-a, then we can only goto
    i-b, so if we look at values from i-b...i+a, and suppose (i-b+a) was forbidden.
    what if we had choosen i-a, position to jump back to i-a-b and then jump forward
    i-a-b+a, we is same as i-b, and we are still forbidden as (i-b+a) is.

    Just be greedy!!!
*/

func minimumJumps(forbidden []int, a int, b int, x int) int {
    max := x

    q := list.New()
    F := make(map[int] bool)
    for _,x := range(forbidden) {
        F[x] = true
        if x > max {max = x}
    }
    max += 2*(a + b)
    jumps := make([]int, max)
    for i,_ := range(jumps) {
        jumps[i] = 10000
        if F[i] {jumps[i] = -1}
    }
    jumps[0] = 0
    q.PushBack(0)
    //BFS
    for q.Len() > 0 {
        pVal := q.Front()
        pos := pVal.Value.(int)
        q.Remove(pVal)

        if pos+a < max && jumps[pos+a] > jumps[pos] + 1 {
            jumps[pos+a] = jumps[pos] + 1
            q.PushBack(pos+a)
        }
        if pos-b > 0 && jumps[pos-b] > jumps[pos] + 1 {
            // we took a jump backwards
            jumps[pos-b] = jumps[pos] + 1
            // lets wrap the jump forward
            if pos+a -b  < max && jumps[pos+a-b] > jumps[pos-b] + 1 {
                jumps[pos+a-b] = jumps[pos-b] + 1
                q.PushBack(pos+a-b)
            }
        }
    }
    if jumps[x] == 10000 {return -1} // we never reached it
    return jumps[x]
}


/*
DFS has a bug
[162,118,178,152,167,100,40,74,199,186,26,73,200,127,30,124,193,84,184,36,103,149,153,9,54,154,133,95,45,198,79,157,64,122,59,71,48,177,82,35,14,176,16,108,111,6,168,31,134,164,136,72,98]
29
98
80
*/
func Min(x,y int) int {
    if x < y {return x}
    return y
}

func minimumJumps(forbidden []int, a int, b int, x int) int {
    max := x


    F := make(map[int] bool)
    for _,x := range(forbidden) {
        F[x] = true
        if x > max {max = x}
    }
    max += 2*(a+b)
    var DFS func(pos int, back bool) int
    DFS = func(pos int, back bool) int {
        if F[pos] || pos < 0 || pos > max {
            return 100000
        }
        if pos == x {return 0}
        ret := DFS(pos+a, back) + 1
        if !back {
            ret = Min(ret, DFS(pos+a-b, true))
        }
        return ret
    }
    ret := DFS(0, false)
    if ret > 10000 {return -1}
    return ret
}
