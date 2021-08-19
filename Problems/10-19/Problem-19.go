/* Problem 19: Counting sundays
You are given the following information, but you may prefer to do some research for yourself.

- 1 Jan 1900 was a Monday.
- Thirty days has September,
	April, June and November.
	All the rest have thirty-one,
	Saving February alone,
	Which has twenty-eight, rain or shine.
	And on leap years, twenty-nine.
- A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
*/
package main

import "fmt"

const (
	SUNDAY    = 0
	MONDAY    = 1
	TUESDAY   = 2
	WEDNESDAY = 3
	THURSDAY  = 4
	FRIDAY    = 5
	SATURDAY  = 6

	JANUARY   = 0
	FEBRUARY  = 1
	MARCH     = 2
	APRIL     = 3
	MAY       = 4
	JUNE      = 5
	JULY      = 6
	AUGUST    = 7
	SEPTEMBER = 8
	OCTOBER   = 9
	NOVEMBER  = 10
	DECEMBER  = 11
)

func main() {
	sundayCount := 0
	dayOfWeek := TUESDAY
	for year := 1901; year < 2001; year++ {
		for month := JANUARY; month <= DECEMBER; month++ {
			if dayOfWeek == SUNDAY {
				sundayCount++
			}
			dayOfWeek += daysInMonth(month, year)
			dayOfWeek %= 7
		}
	}

	fmt.Println(sundayCount, "Sundays fell on the first of the month in the 20th century")
}

func isLeapYear(year int) bool {
	if year%4 == 0 {
		return year%100 != 0 || year%400 == 0
	}
	return false
}

func daysInMonth(month, year int) (days int) {
	daysPerMonth := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	days = daysPerMonth[month]
	if month == FEBRUARY && isLeapYear(year) {
		days++
	}
	return
}
