/*
Design a Skiplist without using any built-in libraries.

A Skiplist is a data structure that takes O(log(n)) time to add, erase and search. Comparing with treap and red-black
tree which has the same function and performance, the code length of Skiplist can be comparatively short and the idea
behind Skiplists are just simple linked lists.

For example: we have a Skiplist containing [30,40,50,60,70,90] and we want to add 80 and 45 into it. The Skiplist works this way:


Artyom Kalinin [CC BY-SA 3.0], via Wikimedia Commons

You can see there are many layers in the Skiplist. Each layer is a sorted linked list. With the help of the top layers,
add , erase and search can be faster than O(n). It can be proven that the average time complexity for each operation is
O(log(n)) and space complexity is O(n).

To be specific, your design should include these functions:

bool search(int target) : Return whether the target exists in the Skiplist or not.
void add(int num): Insert a value into the SkipList.
bool erase(int num): Remove a value in the Skiplist. If num does not exist in the Skiplist, do nothing and return false.
If there exists multiple num values, removing any one of them is fine.
See more about Skiplist : https://en.wikipedia.org/wiki/Skip_list

Note that duplicates may exist in the Skiplist, your code needs to handle this situation.



Example:

Skiplist skiplist = new Skiplist();

skiplist.add(1);
skiplist.add(2);
skiplist.add(3);
skiplist.search(0);   // return false.
skiplist.add(4);
skiplist.search(1);   // return true.
skiplist.erase(0);    // return false, 0 is not in skiplist.
skiplist.erase(1);    // return true.
skiplist.search(1);   // return false, 1 has already been erased.


Constraints:

0 <= num, target <= 20000
At most 50000 calls will be made to search, add, and erase.
*/

type Node struct {
    val int
    next []*Node // log N pointers
}

type Skiplist struct {
    n *Node// log N pointers
    r *rand.Rand
}

func Constructor() Skiplist {
    var s Skiplist
    s.r = rand.New(rand.NewSource(time.Now().UnixNano()))
    s.n = &Node{val:0} // root node
    return s
}

func (this *Skiplist) Locate(target int) []*Node {
    l := len(this.n.next) - 1
    prev := make([]*Node, l+1)
    current := this.n
    for current != nil {
        next := current.next[l]
        if next == nil || next.val >= target {
            prev[l] = current
            if l == 0 {
                break
            } else {
                l--
            }
        } else {
            prev[l] = current
            current = next
        }
    }
    return prev
}

func (this *Skiplist) Search(target int) bool {
    l := len(this.n.next) - 1
    current := this.n
    for current != nil {
        next := current.next[l]
        if next == nil || next.val > target {
            if l == 0 {
                return false
            } else {
                l--
            }
        } else if next.val == target {
            return true
        } else {
            current = next
        }
    }

    return false
}


func (this *Skiplist) Add(num int) {
    l := len(this.n.next)
    var prev []*Node
    if l > 0 {
        prev = this.Locate(num)
    }

    n := 1
    for this.r.Intn(2) != 0 {
        n++
    }

    node := &Node{val:num}
    node.next = make([]*Node, n)
    for ; n > l; n-- {
        this.n.next = append(this.n.next, node)
    }

    for i := 0; i < n; i++ {
        node.next[i] = prev[i].next[i]
        prev[i].next[i] = node
    }
}

func (this *Skiplist) Erase(num int) bool {
    if len(this.n.next) == 0 {
        return false
    }
    prev := this.Locate(num)

    node := prev[0].next[0]
    if node == nil || node.val != num {
        return false
    }
    for i := 0; i < len(node.next); i++ {
        prev[i].next[i] = node.next[i]
    }
    return true
}
