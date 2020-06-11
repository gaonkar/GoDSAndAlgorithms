/*

Given an array of words and a width maxWidth, format the text such that each line has exactly maxWidth characters and is fully (left and right) justified.

You should pack your words in a greedy approach; that is, pack as many words as you can in each line. Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.

Extra spaces between words should be distributed as evenly as possible. If the number of spaces on a line do not divide evenly between words, the empty slots on the left will be assigned more spaces than the slots on the right.

For the last line of text, it should be left justified and no extra space is inserted between words.

Note:

A word is defined as a character sequence consisting of non-space characters only.
Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
The input array words contains at least one word.
Example 1:

Input:
words = ["This", "is", "an", "example", "of", "text", "justification."]
maxWidth = 16
Output:
[
   "This    is    an",
   "example  of text",
   "justification.  "
]
Example 2:

Input:
words = ["What","must","be","acknowledgment","shall","be"]
maxWidth = 16
Output:
[
  "What   must   be",
  "acknowledgment  ",
  "shall be        "
]
Explanation: Note that the last line is "shall be    " instead of "shall     be",
             because the last line must be left-justified instead of fully-justified.
             Note that the second line is also left-justified becase it contains only one word.
Example 3:

Input:
words = ["Science","is","what","we","understand","well","enough","to","explain",
         "to","a","computer.","Art","is","everything","else","we","do"]
maxWidth = 20
Output:
[
  "Science  is  what we",
  "understand      well",
  "enough to explain to",
  "a  computer.  Art is",
  "everything  else  we",
  "do                  "
]


Not difficult, but painful
*/
func Justify(w []string, s, W int, last bool) (ret string) {
    if len(w) == 0 {return ""}
    if len(w) == 1 || last{
        W = W - len(w[0])
        ret = w[0]
        for i:=1; i < len(w); i++ {
            ret = ret + " " + w[i]
            W = W - len(w[i]) - 1
        }
        for i:=0; i < W;i++ {
            ret += " "
        }
        //fmt.Println(ret, len(ret))
        return ret
    }

    N := len(w)-1
    diff := W - s
    step := 0
    sp :=make([]int, N)
    for d:=0; d < diff; {
        step++
        for i:=0; i < len(sp) && d < diff; i++ {
            sp[i]++
            d++
        }

    }
    space := " "
    for i:=0; i < step; i++ {
        space = space + " "
    }
    ret = w[0]

    for i:=1; i < len(w) ; i++ {
        ret = ret + space[:sp[i-1]] + w[i]

    }
    //fmt.Println(ret, len(ret), step, diff)
    return ret
}

func fullJustify(words []string, maxWidth int) (r []string) {

    j := 0
    for j < len(words) {
        i := j
        s := 0
        L := 0
        for i < len(words) && s + len(words[i]) <= maxWidth{
            s += len(words[i]) + 1
            L += len(words[i])
            i++
        }
        r = append(r, Justify(words[j:i],L, maxWidth, i==len(words)))
        j = i
    }
    return r
}
