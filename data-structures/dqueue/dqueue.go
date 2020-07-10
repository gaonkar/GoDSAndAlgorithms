package deque

const initSize=20

type Deque struct {
    lIdx    int     // lower index
    rIdx    int     // upper index
    arr     []interface{}
}

func DequeNew() *Deque {
    ret := new(Deque)
    res.arr = make([]interface{}, initSize)
    res.lIdx = initSize/2
    res.rIdx = initSize/2
    return ret
}

func (d *Deque)(PushLeft
