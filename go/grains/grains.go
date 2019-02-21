package grains

import (
	"errors"
	"math"
)

// Square is a function that uses the input as a guide for how many times to
// double the int 1 and return an array showing the number of grains on each square
func Square(squareNum int) (uint64, error) {
	if squareNum > 64 || squareNum < 1 {
		return 0, errors.New("this selection is invalid")
	}
	sum := math.Pow(2, float64(squareNum-1))
	return uint64(sum), nil
}

// Total is a function that returns the total number of grains on all squares
func Total() uint64 {
	// squares := Square(number)
	// total := squares[len(squares)-1]
	// return total
	var total uint64
	for i := 1; i <= 64; i++ {
		num, _ := Square(i)
		total += num
	}
	return total
}

func main() {

}
