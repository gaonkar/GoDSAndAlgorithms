'''
  If we had only case 1, the answer would be just two states as all bulbs would only flip on or of.
  If we  had case 1 and 2, then we have to think of bulb 0 and 1, [00,01,10,11] then rest would be duplicates
  case 1,2 and 3, we need just 4 states as they effect only bulb 0 and 1, and rest would resemble these two bulbs.
  With case 4, bulb 1,4, 7 get effected so the unique values can be bulb 2,1,3,4 (remember 2 behaves like 0)


The 8 possible states

0  [0, 0, 0, 0]
65 [1, 0, 0, 1]
5  [0, 0, 1, 1]
68 [1, 0, 1, 0]
85 [1, 1, 1, 1]
20 [0, 1, 1, 0]
80 [1, 1, 0, 0]
17 [0, 1, 0, 1]

'''

import copy
arr = [0,0,0,0,0]

def flip1(x):
    if x == 0:
        return 1
    return 0

def flipall(a):
    arr = copy.deepcopy(a)
    for i in range(len(arr)):
        arr[i] = flip1(arr[i])
    return arr

def flipeven(a):
    arr = copy.deepcopy(a)
    arr[2] = flip1(arr[2])
    arr[4] = flip1(arr[4])
    return arr

def flipodd(a):
    arr = copy.deepcopy(a)
    arr[1] = flip1(arr[1])
    arr[3] = flip1(arr[3])
    return arr

def flipother(a):
    arr = copy.deepcopy(a)
    arr[1] = flip1(arr[1])
    arr[4] = flip1(arr[4])
    return arr

d = {}

def encode(arr):
    res = 0
    for x in arr[1:]:
        res = res * 4 + x
    return res

func = [flipall, flipeven, flipodd, flipother]

stk = []
stk.append(arr)
while len(stk) > 0:
    arr = stk.pop()
    x = encode(arr)
    if x in d:
        continue
    d[x] = 0
    print(x, arr[1:])
    for idx in range(0,4):
        arr1 = func[idx](arr)
        y = encode(arr1)
        if y in d:
            pass
        else:
            stk.append(arr1)
            d[x] = d[x] + 1
print(len(d))
