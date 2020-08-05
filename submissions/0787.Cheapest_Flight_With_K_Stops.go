/*
There are n cities connected by m flights. Each flight starts from city u and arrives at v with a price w.

Now given all the cities and flights, together with starting city src and the destination dst, your task is to
find the cheapest price from src to dst with up to k stops. If there is no such route, output -1.

Example 1:
Input:
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 1
Output: 200

The cheapest price from city 0 to city 2 with at most 1 stop costs 200, as marked red in the picture.
Example 2:
Input:
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 0
Output: 500
The cheapest price from city 0 to city 2 with at most 0 stop costs 500, as marked blue in the picture.
https://en.wikipedia.org/wiki/Bellman–Ford_algorithm
*/

func min(x,y uint) uint {
    if x < y {return x}
    return y
}
//BellMan Ford
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
    dp := make([][]uint, K + 2)
    for i,_ := range(dp) {
        dp[i] = make([]uint, n+1)
        for j:=0; j < n+1; j++ {
            dp[i][j] = ^uint(0)
        }
    }
    for  i := 0; i <= K + 1; i++ { dp[i][src] = 0}
    for i := 1; i <= K + 1; i++ {
        for _,f := range(flights) {
            if(dp[i-1][f[0]] != ^uint(0)) {
                dp[i][f[1]] = min(dp[i][f[1]], dp[i-1][f[0]] + uint(f[2]));
            }
        }
    }
    return int(dp[K+1][dst])
}

