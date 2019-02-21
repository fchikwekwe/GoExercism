package main

import "fmt"

// Square is a function that uses the input as a guide for how many times to
// double the int 1 and return an array showing the number of grains on each square
func Square(squareNum int) []int {
	var nums []int
	num := 1
	for i := 0; i < squareNum; i++ {
		fmt.Println(i, num)
		nums = append(nums, num)
		num *= 2
	}
	Total(nums)
}

// Total is a function that returns the total number of grains on all squares
func Total(nums []int) int {
	squares := Square(number)
	total := squares[len(squares)-1]
	return total
	// var total int
	// for i := 0; i < len(squares); i++ {
	// 	fmt.Println(total)
	// 	total += squares[i]
	// }
	// return total
}

func main() {
	fmt.Println(Total(16))

}

// squareNum = 4
// 0 = 1
// 1 = 1 * 2 = 2
// 2 = 2 * 2 = 4
// 3 = 4 * 2
