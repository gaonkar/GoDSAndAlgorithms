'''

Given an array of integers, all occur in triplicate except 1 that is unique. Find that in O(N) time and O(c) space.

When you think of a similar question in base 2 system, given an array where all elements occur in duplicate, except 1.
We immediately think of XOR and the solution would be
uniq = 0
for _,x in enumerate(arr):
  uniq ^= x
return uniq

The idea can be expanded for a base 3 system, convert each number to base 3.
Now for each position, count how many 0s, 1s and 2s exist.
If this count is multiple of 3 (X%3==0) , that must be a part of integers who appear in triplicate.
If this count%3==1, that must be the unique integer, and if the count%3==2, the input has duplicates or uniqueness cannot be determined.

Reconstruct the unique number from these counts for each position.

The time complexity is O(K*N) and space is O(K)  where K is maximum number of digits of a number represented in ternary system.

Else memorize this solution :)
def special(lst):
    ones = 0
    twos = 0
    for x in lst:
        twos |= ones & x
        ones ^= x
        not_threes = ~(ones & twos)
        ones &= not_threes
        twos &= not_threes
    return ones

Here, we represent the number as array of strings to express the concept of the solution.
We can definitely bring down the space usage to O(K). Instead of generating number in string array,
directly compute the number of 0s, 1s, 2s in the number

def tern_opt(num, bit):
    pos = 0
    while num > 0:
        rem = num % 3
        bit[rem][pos] += 1
        num = num // 3
        pos += 1


'''

'''
convert a number to its ternary system, to show how to use the unique representation
'''
def tern(num):
    ret = []
    while num > 0:
        rem = num % 3
        ret.append(str(rem))
        num = num // 3
    return ret

def unique(arr):

    st = []
    maxlen = 0
    for x in range(len(arr)):
        st.append(tern(arr[x]))
        maxlen = max(maxlen, len(st[x]))
    for x in range(len(st)):
        diff = maxlen - len(st[x])
        for i in range(diff):
            st[x].append('0')
        st[x] = st[x][::-1]
    # unique
    uniq = [0] * maxlen

    for i in range(maxlen):
        bit = [0,0,0]
        for x in range(len(st)):
            idx = int(st[x][i]) - int('0')
            bit[idx] += 1
        if bit[0]%3 == 1:
            uniq[i] = 0
        if bit[1]%3 == 1:
            uniq[i] = 1
        if bit[2]%3 == 1:
            uniq[i] = 2

    print(st, uniq)

unique([1,20,3,1,20,3,1,20,3,10,10,10,100])
