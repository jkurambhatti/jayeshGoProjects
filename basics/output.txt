package main

import (
	"fmt"
)

type Date struct {
	day, month, year int
}

var (
	from_date   = Date{2, 9, 1992}
	todays_date = Date{6, 5, 2016}
	total_days  int
)

func main() {

	calculate_days(from_date)
	totd := total_days
	fmt.Println("total days :", totd)
}

func calculate_days(f_date Date) {

	if !((f_date.month == todays_date.month) && (f_date.year == todays_date.year)) {

		total_days += days_in_month(f_date)

		if f_date.month == 12 {
			f_date.year += 1
			f_date.month = 0
		}

		f_date.month += 1
		//fmt.Println(total_days)
		calculate_days(f_date)

	} else if (todays_date.day - from_date.day) < 0 {
		total_days += days_in_month(f_date) + (todays_date.day - from_date.day)
	} else {
		total_days += todays_date.day - from_date.day
	}

}

func days_in_month(fd Date) int {

	switch fd.month {

	case 1, 3, 5, 7, 8, 10, 12:
		return 31
		//total_days += 31
	case 2:
		if isleap(fd.year) {
			//total_days += 29
			return 29
		} else {
			//total_days += 28
			return 28
		}
	default:
		//total_days += 30
		return 30
	}
}

func isleap(year int) bool {
	if year%4 == 0 {
		return true
	} else {
		return false
	}
}
