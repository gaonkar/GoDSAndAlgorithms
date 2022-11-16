
def RGB(st):
    R = 0
    G = 0
    arr = [x for x in st]
    print(arr)
    B = len(arr) - 1
    swap = 0
    while G <= B:
        while arr[B] == 'B':
            B -= 1
        while arr[R] == 'R':
            R += 1

        G = max(R,G)
        if arr[G] == 'R':
            arr[R], arr[G] = arr[G], arr[R]
            R += 1
            G += 1
            swap += 1
        elif arr[G] == 'B':
            arr[B], arr[G] = arr[G], arr[B]
            B -= 1
            swap += 1
        else: # G
            G += 1
    return arr
print(RGB("RRRRRGGGGGGBBBBBB"))
