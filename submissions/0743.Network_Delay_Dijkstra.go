/*

There are N network nodes, labelled 1 to N.

Given times, a list of travel times as directed edges times[i] = (u, v, w),
where u is the source node, v is the target node, and w is the time it takes
for a signal to travel from source to target.

Now, we send a signal from a certain node K. How long will it take for all
nodes to receive the signal? If it is impossible, return -1.



Example 1:



Input: times = [[2,1,1],[2,3,1],[3,4,1]], N = 4, K = 2
Output: 2


Note:

N will be in the range [1, 100].
K will be in the range [1, N].
The length of times will be in the range [1, 6000].
All edges times[i] = (u, v, w) will have 1 <= u, v <= N and 0 <= w <= 100.
*/

func Min(x, y int) int {
	if x < 0 || y < x {
		return y
	}
	return x
}

func Max(x, y int) int {
	if y < x {
		return x
	}
	return y
}

// Implement a heap to reduce it to log(n) operations
func MinNode(D []int, v []bool) int {
	min := -1
	idx := 0
	for i := 1; i < len(D); i++ {
		if v[i] || D[i] < 0 {
			continue
		}
		if min == -1 || min > D[i] {
			min = D[i]
			idx = i
		}
	}
	return idx
}
func networkDelayTime(times [][]int, N int, K int) (r int) {
	E := make(map[int]int)
	for _, x := range times {
		E[x[0]*N+x[1]] = x[2]
	}
	D := make([]int, N+1)
	v := make([]bool, N+1)
	//fmt.Println(A, E)
	Q := make([]int, 0)
	Q = append(Q, K)
	for i := 0; i < len(D); i++ {
		D[i] = -1
	}
	D[K] = 0
	R := N

	for R > 0 {
		n := MinNode(D, v)
		v[n] = true
		for i := 1; i < len(D); i++ { // use a not visited hash map and interate on that
			w, ok := E[n*N+i]
			if i == n || v[i] || !ok {
				continue
			}
			D[i] = Min(D[i], D[n]+w)
		}
		R--
	}
	r = -1
	for i := 1; i < len(D); i++ {
		if D[i] == -1 {
			return -1
		}
		if r < D[i] && i != K {
			r = D[i]
		}
	}
	return r
}
