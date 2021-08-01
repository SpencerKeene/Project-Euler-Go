/* Problem 15: Lattice paths
Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.

[see projecteuler.net for image]

How many such routes are there through a 20×20 grid?
*/
package main

import "fmt"

func main() {
	size := 20
	routes := countRoutes(size)
	fmt.Printf("There are %d routes for a %dx%d grid\n", routes, 20, 20)
}

func countRoutes(n int) int {
	n++
	grid := make([][]int, n) 
	for i := range grid {
		grid[i] = make([]int, n)
	}

	for i := range grid {
		grid[n-1][i] = 1
		grid[i][n-1] = 1
	}

	for i := n - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			grid[i][j] = grid[i][j+1] + grid[i+1][j]
		}
	}

	return grid[0][0]
}