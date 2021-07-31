/* Problem 7: 10001st prime
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	primes := sieveOfEratosthenes(120000)
	fmt.Println(primes[10000])
}

func sieveOfEratosthenes(n int) []int {
	isPrime := make([]bool, n + 1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true;
	}

	limit := int(math.Sqrt(float64(n)))
	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			for j := i*2; j <= n ; j += i {
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