type UnionFind struct {
	label []int
	count []int
}

func CreateUF(N int) UnionFind {
	var uf UnionFind
	uf.label = make([]int, N)
	uf.count = make([]int, N)
	for i := 0; i < N; i++ {
		uf.label[i] = i
		uf.count[i] = 1
	}
	return uf
}

func (this *UnionFind) Find(lbl int) int {
	for lbl != this.label[lbl] {
		this.label[lbl] = this.label[this.label[lbl]]
		lbl = this.label[lbl]
	}
	return this.label[lbl]
}

func (this *UnionFind) Union(x, y int) {
	lblx := this.Find(x)
	lbly := this.Find(y)
	if lblx == lbly {
		return
	}
	this.label[lblx] = this.label[lbly]
	this.count[lbly] += this.count[lblx]
	this.count[lblx] = 0
}

func (this *UnionFind) Max() int {
	max := this.count[0]
	for _, x := range this.count[1:] {
		if max < x {
			max = x
		}
	}
	return max
}