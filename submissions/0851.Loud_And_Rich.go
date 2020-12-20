/*
In a group of N people (labelled 0, 1, 2, ..., N-1), each person has different amounts of money, and different levels of quietness.
For convenience, we'll call the person with label x, simply "person x".
We'll say that richer[i] = [x, y] if person x definitely has more money than person y.  Note that richer may only be a
subset of valid observations. Also, we'll say quiet[x] = q if person x has quietness q.
Now, return answer, where answer[x] = y if y is the least quiet person (that is, the person y with the smallest value of
quiet[y]), among all people who definitely have equal to or more money than person x.

Example 1:

Input: richer = [[1,0],[2,1],[3,1],[3,7],[4,3],[5,3],[6,3]], quiet = [3,2,5,4,6,1,7,0]
Output: [5,5,2,5,4,5,6,7]
Explanation:
answer[0] = 5.
Person 5 has more money than 3, which has more money than 1, which has more money than 0.
The only person who is quieter (has lower quiet[x]) is person 7, but
it isn't clear if they have more money than person 0.

answer[7] = 7.
Among all people that definitely have equal to or more money than person 7
(which could be persons 3, 4, 5, 6, or 7), the person who is the quietest (has lower quiet[x])
is person 7.

The other answers can be filled out with similar reasoning.
*/

func loudAndRich(richer [][]int, quiet []int) []int {
    pdict := make(map[int][]int)
    rdict := make(map[int][]int)
    ret := make([]int, len(quiet))
    // create a poor dictionary of all the folks poor than somebody -->pdict
    // create a rich dictionary of all folks richer than somebody -->rdict
    for _,r  := range(richer) {
        _, ok := pdict[r[1]]
        if ok {
            pdict[r[1]] = append(pdict[r[1]], r[0])
        } else {
            pdict[r[1]] = make([]int, 1)
            pdict[r[1]][0] = r[0]
        }
        _, ok = rdict[r[0]]
        if ok {
            rdict[r[0]] = append(rdict[r[0]], r[1])
        } else {
            rdict[r[0]] = make([]int, 1)
            rdict[r[0]][0] = r[1]
        }
    }
    // for anybody not in that list, mark them as the answer as there is nobody richer than them
    // so they are the quietest
    solve := make([]int, 0) // list of folks whom we could not resolve the richer but quietest folk.

    for i:=0; i < len(ret); i++ {
        _, ok := pdict[i]
        if !ok {
            ret[i]=i
        } else {
            solve = append(solve, i)
            ret[i] = -1
        }
    }

    for len(solve) > 0 {
        n := solve[0]
        missing := false
        curQ := quiet[n] // assume this person is the best we found
        qidx := n
        // look through all folks richer than n and find the best candidate
        for _, x := range(pdict[n]) { //<----1
            if ret[x] < 0 {
                missing = true
                break
            }
            if curQ > quiet[x] { // found a better person
                curQ = quiet[x]
                qidx = x
            }
        }

        solve = solve[1:]
        if missing {
            // put it back in the queue
            solve = append(solve, n)
            continue
        }
        ret[n] = qidx
        for _, x := range(rdict[n]) {
            // for each remaining entry, where n is richer than x, add n to the list
            // crude implementation of union find, because n in richer than x, you add it so that
            // finding best for x in the next iteration at <----1 becomes easier
            _, ok := pdict[x]
            if ok {
                pdict[x] = append(pdict[x], ret[n])
            }
        }
        delete(pdict, n)
    }
    return ret
}
