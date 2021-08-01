/* Problem 14: Longest Collatz sequence
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/
package main

import "fmt"

func main() {
	var longest []int
	for i := 1; i < 1000000; i++ {
		sequence := collatzSequence(i)
		if len(sequence) > len(longest) {
			longest = sequence
		}
	}

	fmt.Println("The longest sequence starts with", longest[0])
	fmt.Println("Its length is", len(longest))
}

func collatzSequence(n int) (sequence []int) {
	for n != 1 {
		sequence = append(sequence, n)
		if n % 2 == 0 {
			n = n / 2
		} else {
			n = 3 * n + 1
		}
	}
	sequence = append(sequence, 1)
	return
}