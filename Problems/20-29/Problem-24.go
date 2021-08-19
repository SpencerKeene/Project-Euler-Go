/* Problem 24: Lexicographic permutations
A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4. If all of the permutations are listed numerically or alphabetically, we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	str := "0123456789"
	perms := getPermutations(str)
	sort.Strings(perms)

	fmt.Printf("The string \"%s\" has %d permutations\n", str, len(perms))
	fmt.Printf("The millionth permutation is %s\n", perms[999999])
}

func getPermutations(str string) []string {
	if len(str) < 2 {
		return []string{str}
	}

	var perms []string
	for i := range str {
		str = swapCharsInString(str, 0, i)

		for _, temp := range getPermutations(str[1:]) {
			perms = append(perms, str[:1]+temp)
		}

		str = swapCharsInString(str, 0, i)
	}
	return perms
}

func swapCharsInString(str string, i1, i2 int) string {
	r := []rune(str)
	r[i1], r[i2] = r[i2], r[i1]
	return string(r)
}
