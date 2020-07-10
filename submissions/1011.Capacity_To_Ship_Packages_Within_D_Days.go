/*
A conveyor belt has packages that must be shipped from one port to another within D days.

The i-th package on the conveyor belt has a weight of weights[i].  Each day,
we load the ship with packages on the conveyor belt (in the order given by weights).
We may not load more weight than the maximum weight capacity of the ship.

Return the least weight capacity of the ship that will result in all the packages
on the conveyor belt being shipped within D days.



Example 1:

Input: weights = [1,2,3,4,5,6,7,8,9,10], D = 5
Output: 15
Explanation:
A ship capacity of 15 is the minimum to ship all the packages in 5 days like this:
1st day: 1, 2, 3, 4, 5
2nd day: 6, 7
3rd day: 8
4th day: 9
5th day: 10

Note that the cargo must be shipped in the order given, so using a ship of
capacity 14 and splitting the packages into parts like (2, 3, 4, 5), (1, 6, 7), (8), (9), (10) is not allowed.
Example 2:

Input: weights = [3,2,2,4,1,4], D = 3
Output: 6
Explanation:
A ship capacity of 6 is the minimum to ship all the packages in 3 days like this:
1st day: 3, 2
2nd day: 2, 4
3rd day: 1, 4
Example 3:

Input: weights = [1,2,3,1,1], D = 4
Output: 3
Explanation:
1st day: 1
2nd day: 2
3rd day: 3
4th day: 1, 1


Constraints:

1 <= D <= weights.length <= 50000
1 <= weights[i] <= 500

        ## Similar to Leetcode: 410. Split Array Largest Sum ##
        ## Similar to Leetcode: 875. Koko Eating Bananas ##
        ## Similar to Leetcode: 1482. Minimum Number of Days to Make m Bouquets ##
*/
func Days(W []int, T int) int {
	i := 1
	t := 0
	for _, w := range W {
		t += w
		if t > T {
			t = w
			i++
		}
	}

	return i
}
func shipWithinDays(weights []int, D int) int {

	lo := weights[0]
	hi := 0
	for _, w := range weights {
		if lo < w {
			lo = w
		}
		hi += w
	}
	//fmt.Println(lo, hi, Days(weights, 8))
	mi := 0
	for lo < hi {
		mi = lo + (hi-lo)/2
		num := Days(weights, mi)
		//fmt.Println(num, lo, mi, hi, D)
		if num <= D {
			hi = mi
		} else {
			lo = mi + 1
		}
	}
	//fmt.Println(lo,mi,hi)
	return lo
}

/*
875. Koko Eating Bananas

Koko loves to eat bananas.  There are N piles of bananas, the i-th pile has
piles[i] bananas.  The guards have gone and will come back in H hours.

Koko can decide her bananas-per-hour eating speed of K.  Each hour, she chooses
some pile of bananas, and eats K bananas from that pile.  If the pile has less
than K bananas, she eats all of them instead, and won't eat any more bananas during this hour.

Koko likes to eat slowly, but still wants to finish eating all the bananas
before the guards come back.

Return the minimum integer K such that she can eat all the bananas within H hours.
*/

func NumSplits(W []int, T int) int {
	i := 0
	for j := 0; j < len(W); j++ {
		t := W[j]
		i = i + t/T
		if t%T != 0 {
			i++
		}
	}
	return i
}

func splitArray(nums []int, D int) int {
	lo := 1
	hi := 0
	for _, w := range nums {
		hi += w
	}
	mi := 3
	//fmt.Println(NumSplits(nums, mi))
	//return 0
	for lo < hi {
		mi = lo + (hi-lo)/2
		num := NumSplits(nums, mi)
		//fmt.Println(num, lo, mi, hi, D)
		if num <= D {
			hi = mi
		} else {
			lo = mi + 1
		}
	}

	//fmt.Println(lo, mi, hi)
	return lo
}

func minEatingSpeed(piles []int, H int) int {
	return splitArray(piles, H)
}
