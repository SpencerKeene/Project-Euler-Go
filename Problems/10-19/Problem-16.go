/* Problem 16: Power digit sum
2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?
*/
package main

import "fmt"

func main() {
	num := "1"
	for i := 0; i < 1000; i++ {
		num = multiplyBy2(num)
	}

	sum := 0
	for i := range num {
		sum += int(num[i] - '0')
	}

	fmt.Println("The sum of the digits of 2^100 is", sum)
}

func multiplyBy2(num string) (result string) {
	var carry byte
	for i := range num {
		digit := num[len(num) - 1 - i] - '0'
		curr := digit * 2 + carry
		
		result = string(curr % 10 + '0') + result
		carry = curr / 10
	}
	if carry != 0 {
		result = string(carry + '0') + result
	}
	return
}