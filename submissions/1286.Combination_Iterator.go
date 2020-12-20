/*
Design an Iterator class, which has:

A constructor that takes a string characters of sorted distinct lowercase English letters and a number combinationLength as arguments.
A function next() that returns the next combination of length combinationLength in lexicographical order.
A function hasNext() that returns True if and only if there exists a next combination.


Example:

CombinationIterator iterator = new CombinationIterator("abc", 2); // creates the iterator.

iterator.next(); // returns "ab"
iterator.hasNext(); // returns true
iterator.next(); // returns "ac"
iterator.hasNext(); // returns true
iterator.next(); // returns "bc"
iterator.hasNext(); // returns false


Constraints:

1 <= combinationLength <= characters.length <= 15
There will be at most 10^4 function calls per test.
It's guaranteed that all calls of the function next are valid.

The cool idea is that use permuation of 0,1 to index the next combination

*/
type CombinationIterator struct {
	N, R, NCR, i int
	arr          []byte
	str          string
}

func Reverse(r []byte, i, n int) {
	//fmt.Println(r,i,n)
	for i < n {
		r[i], r[n] = r[n], r[i]
		n--
		i++
	}
}

//https://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order
func nextPermutation(nums []byte) {
	i := 0
	L := len(nums)
	k := -1
	// find the largest index k such that a[k] < a[k+1]
	for i < L-1 {
		if nums[i] < nums[i+1] {
			k = i
		}
		i++
	}
	if k == -1 {
		Reverse(nums, 0, L-1)
		return
	}
	// find the number that is greater than nums[k] from k+2 to N
	l := L - 1
	for l > k+1 && nums[l] <= nums[k] {
		l--
	}
	//fmt.Println(nums, k,nums[k], l, nums[l])
	// swap them
	nums[k], nums[l] = nums[l], nums[k]
	Reverse(nums, k+1, L-1)
}
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
		gcd := GCD(Nr, Dr)

		if gcd > 1 {
			Nr = Nr / gcd
			Dr = Dr / gcd
		}
	}
	return Nr / Dr
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	var C CombinationIterator
	C.N = len(characters)
	C.R = combinationLength
	C.arr = make([]byte, C.N)
	for i := 0; i < C.N-C.R; i++ {
		C.arr[i] = 1
	}
	C.str = characters
	C.NCR = NumComb(C.N, C.R)
	//fmt.Println(C.arr)
	return C
}

func (this *CombinationIterator) Next() string {
	ret := make([]byte, this.R)
	j := 0
	nextPermutation(this.arr)
	for i := 0; i < this.N; i++ {
		if this.arr[i] == 0 {
			ret[j] = this.str[i]
			j++
		}
	}
	this.i++
	return string(ret)
}

func (this *CombinationIterator) HasNext() bool {
	if this.i < this.NCR {
		return true
	}
	return false
}

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
