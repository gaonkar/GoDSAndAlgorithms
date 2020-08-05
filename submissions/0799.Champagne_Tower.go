/*
Champagne Tower (pascal triangle)

*/
func champagneTower(poured int, query_row int, query_glass int) float64 {
    row := 1
    cup := make([][]float64, 104)
    cup[0] = make([]float64, 1)
    cup[0][0] = float64(poured)
    flag := true
    for row < 104 && flag {
        cup[row] = make([]float64, row+1)
        flag = false
        for col:=0; col < row; col++ {
            if cup[row-1][col] > 1 {
                r := cup[row-1][col] - 1
                cup[row-1][col] = 1
                cup[row][col] += r/2
                cup[row][col+1]+= r/2
                flag = true
            }
        }
        row++
        //fmt.Println(cup[row-1])

    }
    //fmt.Println(cup[49])
    if len(cup[query_row]) == 0 {return 0}
    return cup[query_row][query_glass]
}
