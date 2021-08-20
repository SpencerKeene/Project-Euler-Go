/* Problem 31: Coin sums
In the United Kingdom the currency is made up of pound (£) and pence (p). There are eight coins in general circulation:

1p, 2p, 5p, 10p, 20p, 50p, £1 (100p), and £2 (200p).
It is possible to make £2 in the following way:

1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p
How many different ways can £2 be made using any number of coins?
*/
package main

import "fmt"

func main() {
	coins := []int{200, 100, 50, 20, 10, 5, 2, 1}

	fmt.Printf("£2 can be made in %d different ways\n", combinationsOfCoins(coins, 200))
}

func combinationsOfCoins(coins []int, total int) int {
	if total == 0 {
		return 1
	}
	if len(coins) == 0 {
		return 0
	}

	var combinations int
	for i := total / coins[0]; i >= 0; i-- {
		combinations += combinationsOfCoins(coins[1:], total-coins[0]*i)
	}

	return combinations
}
