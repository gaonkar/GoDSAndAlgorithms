/*
N couples sit in 2N seats arranged in a row and want to hold hands. We want to know the minimum number of swaps so that
every couple is sitting side by side. A swap consists of choosing any two people, then they stand up and switch seats.
The people and seats are represented by an integer from 0 to 2N-1, the couples are numbered in order, the first couple
being (0, 1), the second couple being (2, 3), and so on with the last couple being (2N-2, 2N-1).

The couples' initial seating is given by row[i] being the value of the person who is initially sitting in the i-th seat.

Example 1:

Input: row = [0, 2, 1, 3]
Output: 1
Explanation: We only need to swap the second (row[1]) and third (row[2]) person.
Example 2:

Input: row = [3, 2, 0, 1]
Output: 0
Explanation: All couples are already seated side by side.
Note:

len(row) is even and in the range of [4, 60].
row is guaranteed to be a permutation of 0...len(row)-1.
*/

/*
    [0,1,2,3,4,5,6,7,8,9] //position
    [5,0,1,3,5,2,9,7,8,6] //folks
    This problem can be reduced to connected components.
    1. Suppose there are N couples, (0,1)..(2.3).. (N-2, N-1)
    2. Each of them can be treated as a Node in a graph.
    3. Now, the positions in the row can be treated like tuples because every couple will start
       an even position 0,2,4,6,....
    4. If we pick a position, say 4, then the people from that position 4,5. (folks 5 and 2) will
       be involved in a swap. So pair [2,3] and [4,5] are going to be in that cliq.
    5. We need to find this connected components, and if there are K nodes in this component, we
       need k-1 swaps.
    6. Find all disjoint connected components. So here we will use Union Find
    Why can we use Union Find?
    Note that couples name and location (index) are the same.  For example, we can have
    [2,3,4,5,6,7,0,1] where pair 0,1 are in location(6,7). We can pretend that location is 0,1

*/

func UFind(e int, P []int) int {
    if e == P[e] {return e}
    return UFind(P[e], P)
}

func minSwapsCouples(row []int) int {
    P := make([]int, len(row))

    for i,_:= range(row) {
        P[i] = i - i % 2 // mark each pair as disjoint
    }
    count := 0
    for i:=0; i < len(row); i+=2 {
        fmt.Println(P)
        a, b := UFind(row[i], P), UFind(row[i+1], P)
        //find the parents of adjacent folks. if they are not same, increment as we know
        // we need a swap
        if a != b {count++}
        P[b] = a // add them to same group

    }
    return count
}
