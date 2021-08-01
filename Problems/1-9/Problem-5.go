/* Problem 5: Smallest multiple
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/
package main

import "fmt"

func main() {
	for i := 20; ; i++{
		if isDivisibleUpTo(i, 20) {
			fmt.Println(i)
			break
		}
	}
}

func isDivisibleUpTo(num, limit int) bool {
	for i := 2; i <= limit; i++ {
		if num % i != 0 {
			return false
		}
	}
	return true
}