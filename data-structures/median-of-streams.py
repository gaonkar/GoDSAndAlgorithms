import heapq

'''

Time Complexity O(log(N))
Space Complexity O(N)

'''

class Median:
    def __init__(self):
        self.left = [] # list of elements that are smaller
        self.right = [] # list of elements that are larger
        self.median = 0

    def median(self):
        return self.median

    def add(self, val):
        if val < self.median:
            heapq.heappush(self.left, -val)
        else:
            heapq.heappush(self.right,val)
        Nl = len(self.left)
        Nr = len(self.right)
        if Nl == Nr:
            self.median = (-self.left[0] + self.right[0]) / 2
        elif Nl +1 == Nr:
            self.median = self.right[0]
        elif Nl == 1 + Nr:
            self.median = -self.left[0]
        elif Nl + 2 == Nr:
            out = -heapq.heappop(self.right)
            heapq.heappush(self.left, out)
            self.median = (self.left[0] - self.right[0]) / 2
        elif Nl == 2 + Nr:
            out = -heapq.heappop(self.left)
            heapq.heappush(self.right, out)
            self.median = (self.left[0] - self.right[0]) / 2
        else:
            assert(False)

        self.print()

    def print(self):
        print(self.median, self.left, self.right)

m = Median()
m.add(4)
m.add(3)
m.add(2)
m.add(5)
