/* Problem 4: Largest palindrome product
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/
package main

import "fmt"

func main() {
	max := 0
	for first := 999; first >= 100; first-- {
		for second := first; second >= 100; second -- {
			num := first * second
			if isPalindrome(num) && num > max {
				max = num
			}
		}
	}

	fmt.Println(max)
}

func isPalindrome(x int) bool {
	return x == reverse(x)
}

func reverse(x int) (reverse int) {
	for x > 0 {
		reverse *= 10
		reverse += x % 10
		x /= 10
	}
	return
}