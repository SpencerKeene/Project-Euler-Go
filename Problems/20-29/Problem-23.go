/* Problem 23: Non-abundant sums
A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	abundantNums := make(map[int]struct{})
	for i := 1; i <= 28123; i++ {
		if isAbundant(i) {
			abundantNums[i] = struct{}{}
		}
	}

	total := 0
INTEGER_LOOP:
	for i := 1; i <= 28123; i++ {
		for abNum := range abundantNums {
			required := i - abNum
			if _, ok := abundantNums[required]; ok {
				continue INTEGER_LOOP
			}
		}
		total += i
	}

	fmt.Println("The sum is", total)
}

func isAbundant(x int) bool {
	return sum(properDivisors(x)...) > x
}

func sum(nums ...int) int {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	return sum
}

func properDivisors(x int) []int {
	sqrt := math.Sqrt(float64(x))
	divisors := []int{1}

	for i := 2; i <= int(sqrt); i++ {
		if x%i == 0 {
			divisors = append(divisors, i, x/i)
		}
	}

	if sqrt == float64(int(sqrt)) {
		divisors = divisors[:len(divisors)-1]
	}

	return divisors
}
