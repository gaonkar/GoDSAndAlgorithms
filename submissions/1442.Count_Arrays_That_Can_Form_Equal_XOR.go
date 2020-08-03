/*
Given an array of integers arr.

We want to select three indices i, j and k where (0 <= i < j <= k < arr.length).

Let's define a and b as follows:

a = arr[i] ^ arr[i + 1] ^ ... ^ arr[j - 1]
b = arr[j] ^ arr[j + 1] ^ ... ^ arr[k]
Note that ^ denotes the bitwise-xor operation.

Return the number of triplets (i, j and k) Where a == b.

  Note that a == b, then a^b = 0, so we need to find a range i to k where k - i > 2 and a[i]^...a[k] == 0
  Now that you have generated prefix array, if prefix[i] == prefix[k], then then that range is good count of such ranges

  I think a lot of effort was spent in understanding the question than implementing the solution.
  The question confused the problem with triplets because
  if we have a set of numbers all > 0, and iff l^m^n^o = 0, then l == m^n^o, l^m == n^o, so on, so forth. So this
  question of finding these ijk pairs just confuses the problem. It could have been simplified as find range i..j Where
  a[i]^..^a[j] == 0.
*/

func countTriplets(arr []int) (r int) {
	L := len(arr)
	if L < 2 {
		return 0
	}
	prefix := make([]int, L+1)
	M := make(map[int][]int)
	M[0] = []int{0}
	for i := 1; i < L+1; i++ {
		prefix[i] = arr[i-1] ^ prefix[i-1]
		lst, ok := M[prefix[i]]
		if ok {
			for _, x := range lst {
				lr := (i - 1) - x //count on
				//fmt.Println(x, i-1, lr)
				r += lr
			}
			M[prefix[i]] = append(M[prefix[i]], i)
		} else {
			M[prefix[i]] = []int{i}
		}
	}
	//fmt.Println(prefix, M)

	return r
}
