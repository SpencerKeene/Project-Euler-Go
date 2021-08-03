/* Problem 20: Factorial Digit Sum
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/
package main

import (
	"fmt"
)


func main() {
	factorial := []int{1}
	for i := 2; i <= 100; i++ {
		factorial = multiply(factorial, i)
	}

	sum := 0
	for _, digit := range factorial {
		sum += digit
	}

	fmt.Println("The sum of the digits of 100! is", sum)
}

func multiply(num []int, x int) (result []int) {
	var carry int
	for i := range num {
		product := num[len(num) - 1 - i] * x + carry
		result = append([]int{product % 10}, result...)
		carry = product / 10
	}
	for carry != 0 {
		result = append([]int{carry % 10}, result...)
		carry /= 10
	}
	return
}
