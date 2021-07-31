/* Problem 3: Largest prime factor
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	num := 600851475143
	ans := largestPrimeFactor(num)
	fmt.Println(ans)
}

func largestPrimeFactor(num int) int {
	limit := int(math.Sqrt(float64(num)))
	for curr := 2; curr < limit; curr++ {
		if num % curr == 0 {
			num /= curr
			limit = int(math.Sqrt(float64(num)))
			curr--
		}
	}
	return num
}