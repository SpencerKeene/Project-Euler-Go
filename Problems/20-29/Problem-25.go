/* Problem 25: 1000-digit Fibonacci number
The Fibonacci sequence is defined by the recurrence relation:

Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
Hence the first 12 terms will be:

F1 = 1
F2 = 1
F3 = 2
F4 = 3
F5 = 5
F6 = 8
F7 = 13
F8 = 21
F9 = 34
F10 = 55
F11 = 89
F12 = 144
The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
*/
package main

import (
	"fmt"
)

func main() {
	fib := fibonacci()
	index := 0
	for len(fib()) != 1000 {
		index++
	}

	fmt.Println("The first 1000-digit Fibonacci number is at index", index)
}

func fibonacci() func() []int {
	curr, next := []int{0}, []int{1}
	updateValues := func() {
		nextCopy := make([]int, len(next))
		copy(nextCopy, next)
		curr, next = nextCopy, add(curr, next)
	}

	return func() []int {
		defer updateValues()
		return curr
	}
}

func add(num1, num2 []int) []int {
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}
	
	var (
		carry int
		result []int
	)
	for len(num1) > 0 {
		digit1 := num1[len(num1) - 1]
		digit2 := num2[len(num2) - 1]

		sum := digit1 + digit2 + carry
		result = append([]int{sum % 10}, result...)
		carry = sum / 10

		num1 = num1[:len(num1) - 1]
		num2 = num2[:len(num2) - 1]
	}


	if carry != 0 {
		num2 = add(num2, []int{carry})
	}
	result = append(num2, result...)

	return result
}