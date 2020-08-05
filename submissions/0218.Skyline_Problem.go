/*
A city's skyline is the outer contour of the silhouette formed by all the buildings in that city when viewed from a distance.
Now suppose you are given the locations and height of all the buildings as shown on a cityscape photo (Figure A), write a
program to output the skyline formed by these buildings collectively (Figure B).

Buildings  Skyline Contour
The geometric information of each building is represented by a triplet of integers [Li, Ri, Hi], where Li and Ri are
the x coordinates of the left and right edge of the ith building, respectively, and Hi is its height. It is guaranteed
that 0 ≤ Li, Ri ≤ INT_MAX, 0 < Hi ≤ INT_MAX, and Ri - Li > 0. You may assume all buildings are perfect rectangles
grounded on an absolutely flat surface at height 0.

For instance, the dimensions of all buildings in Figure A are recorded as: [ [2 9 10], [3 7 15], [5 12 12], [15 20 10], [19 24 8] ] .

The output is a list of "key points" (red dots in Figure B) in the format of [ [x1,y1], [x2, y2], [x3, y3], ... ]
that uniquely defines a skyline. A key point is the left endpoint of a horizontal line segment. Note that the last key
point, where the rightmost building ends, is merely used to mark the termination of the skyline, and always has zero height.
Also, the ground in between any two adjacent buildings should be considered part of the skyline contour.

For instance, the skyline in Figure B should be represented as:[ [2 10], [3 15], [7 12], [12 0], [15 10], [20 8], [24, 0] ].

Notes:

The number of buildings in any input list is guaranteed to be in the range [0, 10000].
The input list is already sorted in ascending order by the left x position Li.
The output list must be sorted by the x position.
There must be no consecutive horizontal lines of equal height in the output skyline. For instance,
[...[2 3], [4 5], [7 5], [11 5], [12 7]...] is not acceptable; the three lines of height 5 should be merged
into one in the final output as such: [...[2 3], [4 5], [12 7], ...]

*/

// An MaxHeap is a min-heap of ints.
type MaxHeap struct {
    hh [][]int
    M map[int] int //hashtable to track the position of the building in the heap, makes delete easy
}

func (h MaxHeap) Len() int           { return len(h.hh) }
func (h MaxHeap) Less(i, j int) bool { return h.hh[i][0] > h.hh[j][0] }
func (h MaxHeap) Swap(i, j int)      {
    x := h.hh[i]
    y := h.hh[j]
    h.M[x[1]] = j
    h.M[y[1]] = i
    h.hh[i], h.hh[j] = h.hh[j], h.hh[i]
}

func (h *MaxHeap) Push(x interface{}) {

    arr := x.([]int)
    h.hh = append(h.hh, arr)
    h.M[arr[1]] = len(h.hh) - 1
}

func (h *MaxHeap) Pop() interface{} {
    old := h.hh
    n := len(old)
    x := old[n-1]
    h.hh = old[0 : n-1]
    delete(h.M, x[1])
    return x
}

func getSkyline(buildings [][]int) [][]int {
    if len(buildings) == 0 {return nil}
    V := make([][]int, 0)
    H :=&MaxHeap{}
    heap.Init(H)
    H.M = make(map[int] int)

    for i:=0; i < len(buildings);i++ {
        V = append(V, []int{buildings[i][0], i})
        V = append(V, []int{buildings[i][1], i})
    }
    sort.SliceStable(V, func(i, j int) bool {
    return V[i][0] < V[j][0]
    })

    ret := make([][]int, 0)
    y := 0
    x := 0
    fmt.Println(V)
    for i:=0; i < len(V); i++ {
        bidx := V[i][1]
        B := buildings[bidx]
        x = V[i][0]
        if B[0] == x { // initial position
            heap.Push(H, []int{B[2],bidx})
        } else {
            pos := H.M[bidx]
            heap.Remove(H, pos)
        }
        if H.Len() == 0 {
            y = 0
        } else {
            y = H.hh[0][0]
        }
        if len(ret) == 0 || y != ret[len(ret)-1][1] {
            ret = append(ret,[]int{x,y})
        }
    }
    com := make([][]int, 0)
    i := 0
    //fmt.Println(ret)
    //merge common x
    for i < len(ret) {
        j := i + 1
        x = ret[i][0]
        y = ret[i][1]
        for j < len(ret) && x == ret[j][0] {
            y = ret[j][1]
            j++
        }
        com = append(com, []int{x,y})
        i = j
    }


    ret = com
    //fmt.Println(ret)
    com = make([][]int, 0)
    i = 0
    //merge common y
    for i < len(ret) {
        j := i + 1
        x = ret[i][0]
        y = ret[i][1]
        for j < len(ret) && y == ret[j][1] {
            j++
        }
        com = append(com, []int{x,y})
        i = j
    }

    return com
}
