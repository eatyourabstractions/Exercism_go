package diffsquares

func loop(n int)(int, int) {
    squareOfSum := 0
    sumOfsquares := 0 
    for i :=0; i <= n; i++ {
        squareOfSum += i
        sumOfsquares += i * i
    }
    return squareOfSum * squareOfSum, sumOfsquares
}

func SquareOfSum(n int) int {
    sqsum, _ :=loop(n)
    return sqsum
}

func SumOfSquares(n int) int {
    _, sumOfsqs := loop(n)
    return sumOfsqs
}

func Difference(n int) int {
    x,y := loop(n)
    return x - y
}