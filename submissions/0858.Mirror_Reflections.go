/*
There is a special square room with mirrors on each of the four walls.  Except for the southwest corner, there are
receptors on each of the remaining corners, numbered 0, 1, and 2.
The square room has walls of length p, and a laser ray from the southwest corner first meets the east wall at a distance q from the 0th receptor.

Return the number of the receptor that the ray meets first.  (It is guaranteed that the ray will meet a receptor eventually.)



Example 1:

Input: p = 2, q = 1
Output: 2
Explanation: The ray meets receptor 2 the first time it gets reflected back to the left wall.

Note:

1 <= p <= 1000
0 <= q <= p
*/

func Odd(x int) bool {
    if (x % 2 == 0) {return false}
    return true
}

func Even(x int) bool {
    if (x % 2 == 0) {return true}
    return false
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}


func mirrorReflection(p int, q int) int {

    // Make it relatively prime
    gcd := GCD(p,q)
    p = p / gcd
    q = q / gcd

    // the big leap here is to imagine that we have a tile of square mirrors of length p
    // arranged in column. All the internal mirrors are glass. Then u can imagine and
    // calculate the corner. So
    // receptor 0 will be at location 0,2,4,... on the right
    // receptor 1 will be at location 1,3,5,7... on the right
    // receptor 2 will be at location 2,4,6,... on the left
    if Odd(p) && Odd(q) {
        // at LCM(p,q) it will hit the corner 1
        return 1
    }
    if Odd(p) {
        // means Q is even, it would hit 0
        return 0
    }
    return 2
}
