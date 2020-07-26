/*
A good meal is a meal that contains exactly two different food items with a sum of deliciousness equal to a power of two.
You can pick any two different foods to make a good meal.
Given an array of integers deliciousness where deliciousness[i] is the deliciousness of the ith item of food, 
return the number of different good meals you can make from this list modulo 109 + 7.
Note that items with different indices are considered different even if they have the same deliciousness value.


Example 1:

Input: deliciousness = [1,3,5,7,9]
Output: 4
Explanation: The good meals are (1,3), (1,7), (3,5) and, (7,9).
Their respective sums are 4, 8, 8, and 16, all of which are powers of 2.
Example 2:

Input: deliciousness = [1,1,1,3,3,3,7]
Output: 15
Explanation: The good meals are (1,1) with 3 ways, (1,3) with 9 ways, and (1,7) with 3 ways.
 

Constraints:

1 <= deliciousness.length <= 10^5
0 <= deliciousness[i] <= 2^20
*/

// two sum with Hmap
func Sum2(d []int, s int) int {
    ret := 0
    M := make(map[int] int)
    for _,x :=range(d) {
        if M[s - x] > 0 {
            ret += M[s -x]
        }
        M[x]++
    }
    return ret
}

func countPairs(deliciousness []int) int {
    ret := 0
    power := 1
    maxN := 0

    //find max minor optimization
    M := make(map[int] int)
    for _,x :=range(deliciousness) {
        if maxN < x {
            maxN = x
        }
        M[x]++
    }
    for v,c1 :=range(M) {
        power = 1
        for power <= (2 * maxN) {
            diff := power -v
            c2, ok := M[diff]
            if ok  && c2 > 0 {
                if diff == v {
                    ret += (c1 * (c1-1))/2
                } else {
                    ret += (c1 * c2)
                }
                ret = ret % 1000000007
            }
            power *=2
        }
        M[v] = 0
    }
    return ret
}
