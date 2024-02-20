package diffsquares

func SquareOfSum(n int) int {
    sum := 0
	for i:= n; i>0; i-- {
		sum += i
    }
	return sum*sum
}

func SumOfSquares(n int) int {
	sqsum := 0;
    for i:= n; i>0; i-- {
        sqsum += i*i
    }
	return sqsum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}