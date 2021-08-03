/* Problem 21: Amicable numbers
Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
*/
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()

	amicableNums := make(map[int]struct{})
	for a := 1; a < 10000; a++ {
		_, ok := amicableNums[a]
		if ok {
			continue
		}

		b := d(a)
		if a != b && a == d(b) {
			amicableNums[a] = struct{}{}
			amicableNums[b] = struct{}{}
		}
	}

	sum := 0
	for num := range amicableNums {
		sum += num
	}

	fmt.Println("The sum of all amicable numbers below 10000 is", sum)
	
	fmt.Println("The program completed in", time.Since(start))
}

func d(n int) (sum int) {
	propDivs := getProperDivisors(n)
	for _, x := range propDivs {
		sum += x
	}
	return
}

func getProperDivisors(n int) (properDivisors []int) {
	sqrt := math.Sqrt(float64(n))
	for i := 1; i < int(sqrt); i++ {
		if n % i == 0 {
			properDivisors = append(properDivisors, i)
		}
	}
	if sqrt == float64(int(sqrt)) {
		properDivisors = append(properDivisors, int(sqrt))
	}
	for i := len(properDivisors)-1; i > 0; i-- {
		properDivisors = append(properDivisors, n / properDivisors[i])
	}
	return
}