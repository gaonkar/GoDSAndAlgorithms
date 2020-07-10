/*
Given two words (beginWord and endWord), and a dictionary's word list, find the length
of shortest transformation sequence from beginWord to endWord, such that:

Only one letter can be changed at a time.
Each transformed word must exist in the word list.
Note:

Return 0 if there is no such transformation sequence.
All words have the same length.
All words contain only lowercase alphabetic characters.
You may assume no duplicates in the word list.
You may assume beginWord and endWord are non-empty and are not the same.
Example 1:

Input:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

Output: 5

Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog",
return its length 5.
Example 2:

Input:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

Output: 0

Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.
*/

func Distance(s, t string) bool {
    x := 0
    for i:=0; i < len(s); i++ {
        if s[i] != t[i] {x++}
    }
    return x == 1
}

func AddMap(M *map[string][]string, s, d string) {
    _,ok := (*M)[s]
    if ok {
        (*M)[s]= append((*M)[s], d)
    } else {
        (*M)[s] = []string{d}
    }
}

func Min(x, y int) int {
    if x == 0 {return y} // 0 is infinity
    if x < y {return x}
    return y
}

func ladderLength(beginWord string, endWord string, wordlist []string) int {
    M :=make(map[string] []string)
    V := make(map[string] bool)
    D := make(map[string] int)
    // Generate Adjacency Matrix
    for i:=0; i < len(wordlist); i++ {
        if Distance(beginWord, wordlist[i]) {
            AddMap(&M, beginWord, wordlist[i])
        }
        for j := i+1 ; j < len(wordlist); j++ {
            if Distance(wordlist[i], wordlist[j]) {
                AddMap(&M, wordlist[i], wordlist[j])
                AddMap(&M, wordlist[j], wordlist[i])
            }
        }
    }
    Q:= []string{beginWord}
    D[beginWord] = 1
    // find the shortest path
    for len(Q) > 0 {
        n := Q[0]
        //fmt.Println(n)
        Q = Q[1:]
        if V[n] {
            continue
        }
        V[n] = true
        neighbors := M[n]
        for _,x := range(neighbors) {
            D[x] = Min(D[x], D[n] + 1)
            Q = append(Q, x)
        }
    }
    return D[endWord]
}
