/* Problem 9: Special Pythagorean triplet
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	a, b, c := pythagoreanTripletWithSum(1000)
	fmt.Printf("a, b, c = %d, %d, %d\n", a, b, c)
	fmt.Println("a + b + c =", a+b+c)
	fmt.Println("a * b * c =", a*b*c)
}

func pythagoreanTripletWithSum(sum int) (a, b, c int) {
	for a = 1; ; a++ {
		if a*3+3 > sum {
			break
		}

		for b = a + 1; ; b++ {
			floatC := math.Sqrt(float64(a*a + b*b))
			c = int(floatC)

			if a+b+c > sum {
				break
			}

			if float64(c) == floatC && a+b+c == sum {
				return
			}
		}
	}
	panic("No Pythagorean Triplet exists with that sum.")
}
