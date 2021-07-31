/* Problem 6: Sum square difference
The sum of the squares of the first ten natural numbers is,

The square of the sum of the first ten natural numbers is,

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 - 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/
package main

import "fmt"

func main() {
	num := 100
	diff := squareOfSum(num) - sumOfSquares(num)
	fmt.Println(diff)
}

func sumOfSquares(x int) (sum int) {
	for i := 1; i <= x; i++ {
		sum += i * i
	}
	return
}

func squareOfSum(x int) int {
	sum := x * (x + 1) / 2
	return sum * sum
}