/* Problem 26: Reciprocal cycles
A unit fraction contains 1 in the numerator. The decimal representation of the unit fractions with denominators 2 to 10 are given:

1/2	= 	0.5
1/3	= 	0.(3)
1/4	= 	0.25
1/5	= 	0.2
1/6	= 	0.1(6)
1/7	= 	0.(142857)
1/8	= 	0.125
1/9	= 	0.(1)
1/10	= 	0.1
Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a 6-digit recurring cycle.

Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.
*/
package main

import "fmt"

func main() {
	d, longestCycle := 0, 0
	for i := 1; i < 1000; i++ {
		if cycle := reciprocalCycleLength(i); cycle > longestCycle {
			d = i
			longestCycle = cycle
		}
	}

	fmt.Printf("The value %d has a reciprocal cycle length of %d \n", d, longestCycle)
}

func reciprocalCycleLength(divisor int) int {
	remainderIndices := make(map[int]int)
	dividend := 1

	for currIndex := 0; dividend != 0; currIndex++ {
		remainder := dividend % divisor

		if firstIndex, ok := remainderIndices[remainder]; ok {
			return currIndex - firstIndex
		}
		
		remainderIndices[remainder] = currIndex
		dividend = remainder * 10
	}
	
	return 0
}