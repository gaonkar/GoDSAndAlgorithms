/*
There are 2N people a company is planning to interview. The cost of flying the i-th person to city A is costs[i][0], and the cost of flying the i-th person to city B is costs[i][1].

Return the minimum cost to fly every person to a city such that exactly N people arrive in each city.



Example 1:

Input: [[10,20],[30,200],[400,50],[30,20]]
Output: 110
Explanation:
The first person goes to city A for a cost of 10.
The second person goes to city A for a cost of 30.
The third person goes to city B for a cost of 50.
The fourth person goes to city B for a cost of 20.

The total minimum cost is 10 + 30 + 50 + 20 = 110 to have half the people interviewing in each city.


Note:

1 <= costs.length <= 100
It is guaranteed that costs.length is even.
1 <= costs[i][0], costs[i][1] <= 1000


From user SuccessInVain,
one item can go to either city A or city B. Lets consider both options and write recursive equations:

int AChosen = costs[P][0] + minCost( P-1, A-1, B, costs );
int BChosen = costs[P][1] + minCost( P-1, A, B-1, costs );
where P is total number of persons, A, B are remaining cities. For the next recursion, if you chose A, A reduces by 1 and if you chose B, B reduces by 1 and the total P reduces by 1 too.

So the dp equation is :

dp[P][A][B] = Math.min( costs[P][0] + dp[P-1][A-1][B], costs[P][1] + dp[P-1][A][B-1] )
You can see that it uses three variables: total persons remaining ( P ), remaining for city ( A ), remaining for city ( B ).
and you can "trim" it considering the fact that P = A + B at all times !
so
dp[A][B] = Math.min( costs[A+B][0] + dp[A-1][B], costs[A+B][1] + dp[A][B-1])

*/

func Min(x,y int) int {
    if x < y {return x}
    return y
}

func TCSC_DP(costs [][]int) int {
    L := len(costs)/2
    N := make([][]int, L+1)
    for i:=0; i < L+1; i++ {
        N[i] = make([]int, L+1)
    }
    for i:=1; i <= L; i++ {
        N[i][0] = N[i-1][0] + costs[i-1][0] // cost of having everybody goto city A
        N[0][i] = N[0][i-1] + costs[i-1][1] // cost of having everybody goto city B
    }
    for i:=1; i <= L; i++ {
        for j:=1; j <= L; j++ {
            N[i][j] = Min(N[i][j-1]+costs[i+j-1][1], N[i-1][j] + costs[i+j-1][0])
        }
    }
    return N[L][L]
}

func TCSC_MySol(costs [][]int) (r int) {
    //sort them by costs send the 1st half to A and rest to B
    sort.SliceStable(costs, func(i, j int) bool {
        return (costs[i][0] - costs[i][1]) <= (costs[j][0] - costs[j][1])
    })
    N := len(costs)
    r = 0
    //fmt.Println(costs)
    for i:=0; i < N/2; i++ {
        r = costs[i][0] + costs[N-i-1][1] + r
    }
    return r
}

func twoCitySchedCost(costs [][]int) int {

    //return TCSC_DP(costs)
    return TCSC_MySol(costs)
}
