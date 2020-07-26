/*
Given a string of numbers and operators, return all possible results from computing all the different possible ways to group numbers and operators. The valid operators are +, - and *.

Example 1:

Input: "2-1-1"
Output: [0, 2]
Explanation:
((2-1)-1) = 0
(2-(1-1)) = 2
Example 2:

Input: "2*3-4*5"
Output: [-34, -14, -10, -10, 10]
Explanation:
(2*(3-(4*5))) = -34
((2*3)-(4*5)) = -14
((2*(3-4))*5) = -10
(2*((3-4)*5)) = -10
(((2*3)-4)*5) = 10
*/

def MyOps(op: str, x: int, y: int) -> int:
    if op == '+':
        return x + y
    if op == '-':
        return x - y
    if op == '*':
        return x * y
    assert(True)

class Solution:
    def diffWaysToCompute(self, input: str) -> List[int]:
        tokens = re.split('(\D)', input)
        nums = list(map(int, tokens[::2]))
        ops = tokens[1::2]
        #print(nums, ops)
        def Solve(lo:int, hi:int) -> List[int]:
            r = []
            if lo == hi:
                return [nums[lo]]
            for i in range(lo, hi):
                a = Solve(lo, i)
                b = Solve(i+1, hi)
                for x in a:
                    for y in b:
                        r.append(MyOps(ops[i], x, y))
            assert(len(r) > 0)
            return r
        return Solve(0, len(nums)-1)


/*
Implement a basic calculator to evaluate a simple expression string.

The expression string may contain open ( and closing parentheses ), the plus + or minus sign -, non-negative integers and empty spaces .

Example 1:

Input: "1 + 1"
Output: 2
Example 2:

Input: " 2-1 + 2 "
Output: 3
Example 3:

Input: "(1+(4+5+2)-3)+(6+8)"
Output: 23
*/

class Solution:
    def calculate(self, s: str) -> int:
        nums = []
        ops = []
        num = 0
        rst = 0
        sign = 1
        for x in s:
            if x == ' ':
                continue
            if '0' <= x and x <= '9':
                num = num * 10 + int(x)
                continue
            rst = rst + sign * num
            num = 0
            if x == '+':
                sign = 1
            if x == '-':
                sign = -1
            if x == '(':
                nums.append(rst)
                ops.append(sign)
                rst = 0
                sign = 1
            if x == ')':
                rst = ops.pop() * rst + nums.pop()
        rst = rst + sign * num
        return rst
