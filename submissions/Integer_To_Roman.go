/*
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII,
which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII.
Instead, the number four is written as IV. Because the one is before the five we subtract it making four.
The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.
*/

/*
 * Numbers can range from 1 to 3999
 * Iterate through roman numbers from largest to smallest, keep subtracting it from num until it goes to 0. Append the numerals to out and return it
 */
func intToRoman(num int) string {
    var keys []int
    roman := map[int] string {
        1000:"M",
        900:"CM",
        500:"D",
        400:"CD",
        100:"C",
        90:"XC",
        50:"L",
        40:"XL",
        10:"X",
        9: "IX",
        5:"V",
        4:"IV",
        3:"III",
        2:"II",
        1:"I",
    }
    out := ""

    for key  := range roman {
        keys = append(keys, key)
    }
    sort.Ints(keys)
    for i := len(keys)-1; i >= 0;  {
        if (num >= keys[i]) {
         num = num - keys[i]
         out = out + roman[keys[i]]
        } else {
            i--
        }
    }
    return out
}
