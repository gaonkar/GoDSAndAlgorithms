/*
Implement a trie with insert, search, and startsWith methods.
*/
type Trie struct {
    Val byte
    last bool
    C map[byte] *Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
    var T Trie
    T.Val = 0
    T.last = false
    T.C = make(map[byte] *Trie)
    return T
}

func Cons(x byte, l bool) Trie {
    var T Trie
    T.Val = x
    T.last = l
    T.C = make(map[byte] *Trie)
    return T
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    if len(word) == 0 {return}
    m := this.C
    for i := range(word) {
        T,ok := m[word[i]]
        if !ok {
            l := i == (len(word) - 1)
            X := Cons(word[i], l)
            m[word[i]] = &X
            T = &X
            //fmt.Println(word[i], T)
        } else if i == len(word) - 1 {
            T.last = true
        }
        m = m[word[i]].C
    }
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    if len(word) == 0 {return true}
    m := this.C
    for i := range(word) {
        T,ok := m[word[i]]
        if !ok {
            return false
        }
        m = T.C
        if i == len(word) - 1 && T.last == true {
            return true
        }
    }
    return len(m) == 0
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(word string) bool {
    if len(word) == 0 {return true}
    m := this.C
    T := this
    ok := false
    for i := range(word) {
        //fmt.Println(word[i], T)
        T,ok = m[word[i]]
        if !ok {
            return false
        }
        m = T.C
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
