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
// SumOfSquares is a function that takes the squares of the first N natural numbers and adds them up
func SumOfSquares(N int) int {
    var sumSquare int
    for i := 1; i < N+1; i++ {

        sumSquare += i * i
    }

    return sumSquare
}
// Difference is a function that subtracts the Sum of the Squares of the first N natural numbers from the Square of their sum 
func Difference(N int) (int) {
    squareSum := SquareOfSum(N)
    sumSquare := SumOfSquares(N)

    return squareSum - sumSquare
}
