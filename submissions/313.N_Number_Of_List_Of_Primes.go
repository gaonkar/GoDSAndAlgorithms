/*
Ugly Numbers
Write a program to find the nth super ugly number.

Super ugly numbers are positive numbers whose all prime factors are in the given prime list primes of size k.

Example:

Input: n = 12, primes = [2,7,13,19]
Output: 32
Explanation: [1,2,4,7,8,13,14,16,19,26,28,32] is the sequence of the first 12
             super ugly numbers given primes = [2,7,13,19] of size 4.
Note:

1 is a super ugly number for any given primes.
The given numbers in primes are in ascending order.
0 < k ≤ 100, 0 < n ≤ 106, 0 < primes[i] < 1000.
The nth super ugly number is guaranteed to fit in a 32-bit signed integer.
*/
func Min(x, y int) int {
    if x < y {return x}
    return y
}

func nthUglyNumber(n int, primes []int) int {
    if n == 0 {return 0}
    upm:= make([]int, len(primes))
    ugly := make([]int, n+1)
    ugly[0] = 1
    i := 1
    for i < n {
        ugly[i]=ugly[upm[0]] * primes[0]
        for j:=0; j < len(primes);j++ {
            ugly[i] = Min(ugly[i], ugly[upm[j]]*primes[j])
        }
        for j:=0; j < len(primes);j++ {
            if ugly[i] == ugly[upm[j]]*primes[j] {
                upm[j]++
            }
        }
        i++
    }
    //fmt.Println(ugly)
    return ugly[n-1]
}

func nthSuperUglyNumber(n int, primes []int) int {
    return nthUglyNumber(n,primes)
}
