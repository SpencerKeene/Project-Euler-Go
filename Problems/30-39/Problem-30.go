/* Problem 30: Digit fifth powers
Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:

1634 = 14 + 64 + 34 + 44
8208 = 84 + 24 + 04 + 84
9474 = 94 + 44 + 74 + 44
As 1 = 14 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var numbers []int
	for i := 2; i < 1000000; i++ {
		if isSumOfFifthPowers(i) {
			numbers = append(numbers, i)
		}
	}

	var sum int
	for _, x := range numbers {
		sum += x
	}

	fmt.Println("The numbers which can be written as the sum of fifth powers of their digits are:\n", numbers)
	fmt.Println("The sum of these numbers is", sum)
}

func isSumOfFifthPowers(x int) bool {
	sum, original := 0, x
	for x > 0 {
		sum += int(math.Pow(float64(x%10), 5))
		x /= 10
	}
	return sum == original
}
