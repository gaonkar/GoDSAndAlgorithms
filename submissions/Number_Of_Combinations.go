package main

import "fmt"

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

// NCR = N!/(N-R)!R!
func NumComb(N, R int) int {
	Nr, Dr := 1, 1
	N_R := N - R
	for N_R > 0 || R > 0 {
		Nr = Nr * N
		if R > N_R {
			Dr = Dr * R
			R--
		} else {
			Dr = Dr * N_R
			N_R--
		}
		N--
		fmt.Println(Nr, Dr)
		gcd := GCD(Nr, Dr)
		Nr = Nr / gcd
		Dr = Dr / gcd
	}
	return Nr / Dr
}

func main() {
	fmt.Println(NumComb(100, 24))
}
