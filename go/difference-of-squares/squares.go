package diffsquares

import (
    "fmt"
)

// SquareOfSum is a function that produces the square of the sum of the
// first N natural numbers by adding them up and then squareing that sum
func SquareOfSum(N int) (int) {
    var sum int
    for i := 1; i < N+1; i++ {
        sum += i
    }
    square := sum * sum
    return square
}

func SumOfSquares(N int) int {
    var sumSquare int
    for i := 1; i < N+1; i++ {

        sumSquare += i * i
    }

    return sumSquare
}
// Difference is a function that ...
func Difference(N int) (int) {
    squareSum := SquareOfSum(N)
    sumSquare := SumOfSquares(N)

    return squareSum - sumSquare
}
