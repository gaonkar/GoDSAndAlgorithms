/*
Given an integer array instructions, you are asked to create a sorted array from the elements in instructions.
You start with an empty container nums. For each element from left to right in instructions, insert it into nums.
The cost of each insertion is the minimum of the following:

The number of elements currently in nums that are strictly less than instructions[i].
The number of elements currently in nums that are strictly greater than instructions[i].
For example, if inserting element 3 into nums = [1,2,3,5], the cost of insertion is min(2, 1)
(elements 1 and 2 are less than 3, element 5 is greater than 3) and nums will become [1,2,3,3,5].

Return the total cost to insert all elements from instructions into nums. Since the answer may be large, return it modulo 109 + 7



Example 1:

Input: instructions = [1,5,6,2]
Output: 1
Explanation: Begin with nums = [].
Insert 1 with cost min(0, 0) = 0, now nums = [1].
Insert 5 with cost min(1, 0) = 0, now nums = [1,5].
Insert 6 with cost min(2, 0) = 0, now nums = [1,5,6].
Insert 2 with cost min(1, 2) = 1, now nums = [1,2,5,6].
The total cost is 0 + 0 + 0 + 1 = 1.
Example 2:

Input: instructions = [1,2,3,6,5,4]
Output: 3
Explanation: Begin with nums = [].
Insert 1 with cost min(0, 0) = 0, now nums = [1].
Insert 2 with cost min(1, 0) = 0, now nums = [1,2].
Insert 3 with cost min(2, 0) = 0, now nums = [1,2,3].
Insert 6 with cost min(3, 0) = 0, now nums = [1,2,3,6].
Insert 5 with cost min(3, 1) = 1, now nums = [1,2,3,5,6].
Insert 4 with cost min(3, 2) = 2, now nums = [1,2,3,4,5,6].
The total cost is 0 + 0 + 0 + 0 + 1 + 2 = 3.
Example 3:

Input: instructions = [1,3,3,3,2,4,2,1,2]
Output: 4
Explanation: Begin with nums = [].
Insert 1 with cost min(0, 0) = 0, now nums = [1].
Insert 3 with cost min(1, 0) = 0, now nums = [1,3].
Insert 3 with cost min(1, 0) = 0, now nums = [1,3,3].
Insert 3 with cost min(1, 0) = 0, now nums = [1,3,3,3].
Insert 2 with cost min(1, 3) = 1, now nums = [1,2,3,3,3].
Insert 4 with cost min(5, 0) = 0, now nums = [1,2,3,3,3,4].
​​​​​​​Insert 2 with cost min(1, 4) = 1, now nums = [1,2,2,3,3,3,4].
​​​​​​​Insert 1 with cost min(0, 6) = 0, now nums = [1,1,2,2,3,3,3,4].
​​​​​​​Insert 2 with cost min(2, 4) = 2, now nums = [1,1,2,2,2,3,3,3,4].
The total cost is 0 + 0 + 0 + 0 + 1 + 0 + 1 + 0 + 2 = 4.


Constraints:

1 <= instructions.length <= 105
1 <= instructions[i] <= 105
*/

//TLE O(N^2)

func BinarySearch(nums []int, L, target int) int {
    low := 0
    hi := L-1
    for low <= hi {
        mid := (hi - low) /2 + low ;
        //fmt.Println(low, mid, hi)
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            hi = mid - 1
        } else {
            low = mid + 1
        }

    }
    return low // insert location
}

func Min(x,y int) int {
    if x < y {return x}
    return y
}

func Move(A []int, idx, L int) {
    for i := L; i > idx; i-- {
        A[i] = A[i-1]
    }
}

func createSortedArray(instructions []int) int {
    nums := make([]int, len(instructions)+1)
    Lnum := 0
    ret := 0
    for _,x := range(instructions) {
        idx := BinarySearch(nums, Lnum, x) // find the index
        l,h := idx, idx
        for l > 0 && nums[l-1] == x {l--} // handle the duplicates
        for h <= Lnum && nums[h] == x {h++} // handle the duplicates
        //fmt.Println(x, l, idx, h, Lnum - h, nums, ret)
        Move(nums, idx, Lnum) // O(N) insert
        nums[idx] = x
        ret += Min(l, Lnum-h)

        Lnum++
    }
    return ret
}


// https://en.wikipedia.org/wiki/Fenwick_tree#Implementation to maintain the prefix index
// O(NlogN) time but this is O(Maxarr) space too lazy to implement btree 

// find the least set bit
func LSB(i int) int {
    return (i & -i)
}

type  BITree struct{
    A []int
}

func (bit BITree) Update(x int) {
    for x < len(bit.A) {
        bit.A[x]++
        x += LSB(x)
    }
}

func (bit BITree) PSum(x int) int {
    ret := 0
    for x > 0 {
        ret += bit.A[x]
        x -= LSB(x)
    }
    return ret
}

func MaxArr(A []int) int {
    ret := 0
    for _, x:= range(A) {
        if ret < x {ret = x}
    }
    return ret
}

func CreateBIT(A []int) BITree {
    var bit BITree
    bit.A = make([]int, MaxArr(A) + 1)
    return bit
}

func Min(x,y int) int {
    if x < y {return x}
    return y
}




func createSortedArray(instructions []int) int {
    // https://en.wikipedia.org/wiki/Fenwick_tree#Implementation
    bit := CreateBIT(instructions)
    ret := 0
    for i, x:= range(instructions) {
        // [..Psum(x)...i]
        ret += Min(bit.PSum(x-1), i - bit.PSum(x))
        bit.Update(x)
    }

    return ret % (1000000007) //requires modulo but the int is 64 bit
}
