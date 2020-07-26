

func reverseBits(num uint32) uint32 {
    L := 0xFFFFFFFF & int(num)
    O := 0
    i := 0
    for i=0; i < 32 && L > 0; i++ {
        L = L << 1
        //fmt.Println(num, L, 0xFFFFFFFF)
        if L >= 0xFFFFFFFF {
            O = O | 0x100000000
        }
        O = O >> 1
        L = L & 0xFFFFFFFF
        //fmt.Println(L,O)
    }
    O = O >> (32-i)
    return uint32(O)
}

result = 0
for(int i=0; i<32; i++){
    result = result << 1;
    //bitwise OR is cheaper than add
    result = result | n&1;
    n = n >> 1;
}

//swapping 16 bits at a time with masking
    uint32_t reverseBits(uint32_t n) {
        n = (n >> 16) | (n << 16);
        n = ((n & 0xff00ff00) >> 8) | ((n & 0x00ff00ff) << 8);
        n = ((n & 0xf0f0f0f0) >> 4) | ((n & 0x0f0f0f0f) << 4);
        n = ((n & 0xcccccccc) >> 2) | ((n & 0x33333333) << 2);
        n = ((n & 0xaaaaaaaa) >> 1) | ((n & 0x55555555) << 1);
        return n;
    }
}

func Power(n int) int {
    if n< 1 {return 1}
    return 2 * Power(n-1)
}
func grayCode(n int) []int {
    N := Power(n)
    r := make([]int, N)
    if n==0 {return r}
    r[1] = 1 // already have 1 bit grey code [0,1]
    n = 4 // we are starting with 2 bit
    for n <= N {
        /*
        for each element from n/2 to n-1, we or it with its mirror image
        */
        for i:=n/2; i < n;i++ {
            r[i] = (n/2) | r[n-i-1]
        }
        n *=2
        fmt.Println(r)
    }
    return r
}
