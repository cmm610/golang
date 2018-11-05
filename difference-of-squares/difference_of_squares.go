package diffsquares

func SquareOfSum(val int) int {
	sum := 0
	for i:=1; i<=val; i++ {
		sum += i
	}
	return sum * sum
}

func SumOfSquares(val int) int {
	sum := 0
	for i:=1; i<=val; i++ {
		sq := i * i
		sum += sq
	}
	return sum
}

func Difference(val int) int {
	return SquareOfSum(val) - SumOfSquares(val)
}