/* Problem 10: Summation of primes
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	primes := sieveOfEratosthenes(2000000)

	sum := 0
	for _, prime := range primes {
		sum += prime
	}

	fmt.Println(sum)
}

func sieveOfEratosthenes(n int) []int {
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	limit := int(math.Sqrt(float64(n)))
	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			for j := i * 2; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	primes := make([]int, 0)
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}
