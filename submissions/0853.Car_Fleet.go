/*
N cars are going to the same destination along a one lane road.  The destination is target miles away.
Each car i has a constant speed speed[i] (in miles per hour), and initial position position[i] miles towards the target along the road.
A car can never pass another car ahead of it, but it can catch up to it, and drive bumper to bumper at the same speed.
The distance between these two cars is ignored - they are assumed to have the same position.
A car fleet is some non-empty set of cars driving at the same position and same speed.  Note that a single car is also a car fleet.
If a car catches up to a car fleet right at the destination point, it will still be considered as one car fleet.
How many car fleets will arrive at the destination?



Example 1:

Input: target = 12, position = [10,8,6, 0,5,3], speed = [2,4,6, 1,1,3]
Output: 3
Explanation:
The cars starting at 10 and 8 become a fleet, meeting each other at 12.
The car starting at 0 doesn't catch up to any other car, so it is a fleet by itself.
The cars starting at 5 and 3 become a fleet, meeting each other at 6.
Note that no other cars meet these fleets before the destination, so the answer is 3.

Note:

0 <= N <= 10 ^ 4
0 < target <= 10 ^ 6
0 < speed[i] <= 10 ^ 6
0 <= position[i] < target
All initial positions are different.
*/

/*
 Solution: Using Map
 1. For each car, find how long will it take to reach the target and put that in a Map
 2. sort the position such that farthest is first in the list.
    map[-10:1 -8:1 -6:1 -5:7 -3:3 0:12] for the given example
    So if you look at this map, any cars that have same speed will reach as 1 fleet.
    So if there are increasing time duration, each reach at independent time. So 1 will be followed by 7.
    But those car stuck behind -5 (like -3) will join to become 1 fleet.
    then car at position 0 will reach.

    Time : O(N Log N)
    Space: O(N)
*/

func carFleet(target int, position []int, speed []int) int {
    M := make(map[int]float64)
    for i,_ := range(position) {
        M[-position[i]] = float64((target - position[i])) / float64(speed[i])
        position[i] = -position[i]
    }
    sort.Ints(position)
    ret := 0
    cur := 0.0
    //fmt.Println(M)
    for _,p := range(position) {
        x := M[p]
        if cur < x {
            cur = x
            ret++
        }
    }
    return ret
}
