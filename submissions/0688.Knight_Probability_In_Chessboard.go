/*
688
On an NxN chessboard, a knight starts at the r-th row and c-th column and attempts to make exactly K moves.
The rows and columns are 0 indexed, so the top-left square is (0, 0), and the bottom-right square is (N-1, N-1).

A chess knight has 8 possible moves it can make, as illustrated below. Each move is two squares in a cardinal direction,
then one square in an orthogonal direction.

Each time the knight is to move, it chooses one of eight possible moves uniformly at random (even if the piece
would go off the chessboard) and moves there.

The knight continues moving until it has made exactly K moves or has moved off the chessboard.
Return the probability that the knight remains on the board after it has stopped moving.

Similar to Out of Bounds problem

*/
func Valid(i, j, N int) bool {
	if 0 <= i && i < N && 0 <= j && j < N {
		return true
	}
	return false
}

func findPaths(N, K, i, j int) float64 {
	M := map[int]map[int]map[int]float64{}
	D := [][]int{{-2, -1}, {-1, -2}, {1, -2}, {2, -1}, {2, 1}, {1, 2}, {-1, 2}, {-2, 1}}
	var PathMemo func(N, i, j int) float64
	PathMemo = func(K, i, j int) float64 {
		if K == 0 {
			return 1
		}
		if val, ok := M[i][j][K]; ok {
			return val
		}
		r := float64(0)
		for _, d := range D {
			x := i + d[0]
			y := j + d[1]
			if Valid(x, y, N) {
				r += 0.125 * PathMemo(K-1, x, y)
			}
		}
		_, ok := M[i]
		if !ok {
			M[i] = map[int]map[int]float64{}
		}
		_, ok = M[i][j]
		if !ok {
			M[i][j] = map[int]float64{}
		}
		M[i][j][K] = r
		return r
	}
	return PathMemo(K, i, j)
}

func knightProbability(N int, K int, r int, c int) float64 {
	return findPaths(N, K, r, c)
}

// 576 Out of Bounds

func Valid(i, j, m, n int) bool { // check if i,j is valid
	if 0 <= i && i < m && 0 <= j && j < n {
		return true
	}
	return false
}
func findPaths(m int, n int, N int, i int, j int) int {
	M := map[int]map[int]map[int]int{}             //memoization
	D := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} //direction
	var PathMemo func(N, i, j int) int
	MOD := 1000*1000*1000 + 7
	//fmt.Println(MOD)
	PathMemo = func(N, i, j int) int {
		if N == 0 {
			if Valid(i, j, m, n) {
				return 0
			}
			return 1
		}
		if val, ok := M[i][j][N]; ok {
			return val
		}
		r := 0
		for _, d := range D {
			x := i + d[0]
			y := j + d[1]
			if Valid(x, y, m, n) {
				r = (PathMemo(N-1, x, y) + r) % MOD
			} else {
				r = (r + 1) % MOD
			}
		}
		_, ok := M[i]
		if !ok {
			M[i] = map[int]map[int]int{}
		}
		_, ok = M[i][j]
		if !ok {
			M[i][j] = map[int]int{}
		}
		M[i][j][N] = r
		return r
	}
	return PathMemo(N, i, j)
}
