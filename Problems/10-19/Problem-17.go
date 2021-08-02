/* Problem 17: Number letter counts
If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?

NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen)
contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.
*/
package main

import "fmt"

func main() {
	totalLetters := 0
	for i := 1; i <= 1000; i++ {
		totalLetters += countLetters(i)
	}

	fmt.Println(totalLetters)
}

var letters = map[int]int{
	1: len("one"),
	2: len("two"),
	3: len("three"),
	4: len("four"),
	5: len("five"),
	6: len("six"),
	7: len("seven"),
	8: len("eight"),
	9: len("nine"),

	10: len("ten"),
	11: len("eleven"),
	12: len("twelve"),
	13: len("thirteen"),
	14: len("fourteen"),
	15: len("fifteen"),
	16: len("sixteen"),
	17: len("seventeen"),
	18: len("eighteen"),
	19: len("nineteen"),

	20: len("twenty"),
	30: len("thirty"),
	40: len("forty"),
	50: len("fifty"),
	60: len("sixty"),
	70: len("seventy"),
	80: len("eighty"),
	90: len("ninety"),
}

func countLetters(num int) int {
	switch {
	case letters[num] != 0:
		return letters[num]
	case num == 1000:
		return len("one") + len("thousand")
	case num >= 100 && num % 100 == 0:
		return letters[num / 100] + len("hundred")
	case num >= 100:
		return letters[num / 100] + len("hundred") + len("and") + countLetters(num % 100)
	default:
		return letters[num / 10 * 10] + letters[num % 10]
	}
}