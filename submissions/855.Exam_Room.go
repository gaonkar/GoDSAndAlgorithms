/*
In an exam room, there are N seats in a single row, numbered 0, 1, 2, ..., N-1.

When a student enters the room, they must sit in the seat that maximizes the distance to the closest person.
If there are multiple such seats, they sit in the seat with the lowest number.
(Also, if no one is in the room, then the student sits at seat number 0.)

Return a class ExamRoom(int N) that exposes two functions:
ExamRoom.seat() returning an int representing what seat the student sat in, and
ExamRoom.leave(int p) representing that the student in seat number p now leaves the room.
It is guaranteed that any calls to ExamRoom.leave(p) have a student sitting in seat p.



Example 1:

Input: ["ExamRoom","seat","seat","seat","seat","leave","seat"], [[10],[],[],[],[],[4],[]]
Output: [null,0,9,4,2,null,5]
Explanation:
ExamRoom(10) -> null
seat() -> 0, no one is in the room, then the student sits at seat number 0.
seat() -> 9, the student sits at the last seat number 9.
seat() -> 4, the student sits at the last seat number 4.
seat() -> 2, the student sits at the last seat number 2.
leave(4) -> null
seat() -> 5, the student sits at the last seat number 5.
​​​​​​​

Note:

1 <= N <= 10^9
ExamRoom.seat() and ExamRoom.leave() will be called at most 10^4 times across all test cases.
Calls to ExamRoom.leave(p) are guaranteed to have a student currently sitting in seat number p.
*/import "fmt"

type ExamRoom struct {
	P []int
	N int
	e bool
}

func Log(enable bool, a ...interface{}) {
	if enable {
		fmt.Println(a...)
	}
}

func Constructor(N int) ExamRoom {
	var E ExamRoom
	E.P = make([]int, 0)
	E.N = N
	E.e = false
	return E
}

func (this *ExamRoom) Insert() int {
	if len(this.P) == 0 {
		this.P = append(this.P, 0)
		return 0
	}
	L := len(this.P)
	seatnum, maxd, index := 0, this.P[0], 0
	for i := 0; i < L-1; i++ {
		curd := (this.P[i+1] - this.P[i]) / 2
		if curd > maxd {
			maxd = curd
			seatnum = (this.P[i+1] + this.P[i]) / 2
			index = i + 1
		}
	}
	if this.N-1-this.P[L-1] > maxd {
		this.P = append(this.P, this.N-1)
		return this.N - 1
	}
	this.P = append(this.P[:index+1], this.P[index:]...)
	this.P[index] = seatnum
	return seatnum
}

func (this *ExamRoom) Seat() int {
	Log(this.e, "S", this.P)
	return this.Insert()
}

func (this *ExamRoom) Leave(p int) {
	lo, mi, hi := 0, 0, len(this.P)-1
	Log(this.e, "LB", p, this.P)
	for lo <= hi {
		mi = lo + (hi-lo)/2
		Log(this.e, lo, mi, hi, this.P[mi])
		if this.P[mi] == p {
			this.P = append(this.P[:mi], this.P[mi+1:]...)
			Log(this.e, "LE", this.P, mi)
			return
		} else if this.P[mi] > p {
			hi = mi - 1
		} else {
			lo = mi + 1
		}
	}
	Log(this.e, "here")
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(N);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
