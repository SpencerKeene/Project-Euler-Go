/* Problem 27: Quadratic primes
Euler discovered the remarkable quadratic formula:
	n^2 + n + 41

It turns out that the formula will produce 40 primes for the consecutive integer values 0 <= n <= 39.
However, when n = 40, 40^2 + 40 + 41 = 40(40 + 1) + 41 is divisible by 41, and certainly
when n = 41, 41^2 + 41 + 41 is clearly divisible by 41.

The incredible formula n^2 + 79n + 1601 was discovered, which produces 80 primes for the consecutive
values 0 <= n <= 79. The product of the coefficients, −79 and 1601, is −126479.

Considering quadratics of the form:
	n^2 + an + b, where |a| < 1000 and |b| <= 1000

	where |n| is the modulus/absolute value of n
	e.g. |11| = 11 and |-4| = 4

Find the product of the coefficients, a and b, for the quadratic expression that produces the maximum
number of primes for consecutive values of n, starting with n = 0.
*/
package main

import (
	"fmt"
	"math"
)

type coefficients struct {
	a, b int
}

func main() {
	var (
		maxN     int
		maxCoeff coefficients
	)
	for a := -999; a < 1000; a++ {
		for b := -999; b <= 1000; b++ {
			coeff := coefficients{a, b}
			f := quadraticExpression(coeff)
			n := consecutivePrimesGenerated(f)

			if n > maxN {
				maxN = n
				maxCoeff = coeff
			}
		}
	}

	fmt.Printf("The coefficients a = %d and b = %d produce %d consecutive primes\n", maxCoeff.a, maxCoeff.b, maxN)
	fmt.Printf("Their product is %d\n", maxCoeff.a*maxCoeff.b)
}

func quadraticExpression(c coefficients) func(int) int {
	a, b := c.a, c.b
	return func(n int) int {
		return n*n + a*n + b
	}
}

func consecutivePrimesGenerated(f func(int) int) (n int) {
	for isPrime(f(n)) {
		n++
	}
	return
}

func isPrime(x int) bool {
	if x < 2 {
		return false
	}

	sqrt := int(math.Sqrt(float64(x)))
	for i := 2; i < sqrt; i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}
