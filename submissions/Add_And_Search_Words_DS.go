/*
Add and Search Word - Data structure design
Design a data structure that supports the following two operations:

void addWord(word)
bool search(word)
search(word) can search a literal word or a regular expression string containing only letters a-z or .. A . means it can represent any one letter.

Example:

addWord("bad")
addWord("dad")
addWord("mad")
search("pad") -> false
search("bad") -> true
search(".ad") -> true
search("b..") -> true
Note:
You may assume that all words are consist of lowercase letters a-z

*/
type WordDictionary struct {
    Val byte
    last bool
    C map[byte] *WordDictionary
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
    var T WordDictionary
    T.Val = 0
    T.last = false
    T.C = make(map[byte] *WordDictionary)
    return T
}
func Cons(x byte, l bool) WordDictionary {
    var T WordDictionary
    T.Val = x
    T.last = l
    T.C = make(map[byte] *WordDictionary)
    return T
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
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


/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
    //if len(word) == 0 {return true}
    m := this.C
    T := this
    ok := false
    for i := range(word) {
        if word[i] == '.' {
            for _,x:= range(T.C) {
                ok = x.Search(word[i+1:])
                if ok { return ok}
            }
            return false
        }
        T,ok = m[word[i]]
        if !ok {
            return false
        }
        m = T.C
        if i == len(word) - 1 && T.last == true {
            return true
        }
    }
    return len(m) == 0 || this.last
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
