/* Problem 32: Pandigital products
We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.

HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	pandigitals := make(map[int]struct{})

	for m := 1; m <= 9; m++ {
		for n := 1000; n <= 9999; n++ {
			prod := m * n
			if isPandigital(m, n, prod) {
				pandigitals[prod] = struct{}{}
			}
		}
	}

	for m := 10; m <= 99; m++ {
		for n := 100; n <= 999; n++ {
			prod := m * n
			if isPandigital(m, n, prod) {
				pandigitals[prod] = struct{}{}
			}
		}
	}

	var sum int
	for val := range pandigitals {
		sum += val
	}

	fmt.Println("The sum of all pandigital numbers is", sum)
}

func isPandigital(multiplicand, multiplier, product int) bool {
	var digits []int

	digits = appendDigits(digits, multiplicand)
	digits = appendDigits(digits, multiplier)
	digits = appendDigits(digits, product)

	sort.Ints(digits)
	for i, val := range digits {
		if i != val-1 {
			return false
		}
	}
	return true
}

func appendDigits(digits []int, x int) []int {
	for x > 0 {
		digits = append(digits, x%10)
		x /= 10
	}
	return digits
}
