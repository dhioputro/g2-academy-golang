package main

import (
	"fmt"
)

func main() {

	// t := time.Now()
	// startDay := t.Day()

	// var initYear int = 1920
	var year int = 2020
	var blankDays int = 3

	var months = make([]string, 13)
	months[1] = "January"
	months[2] = "February"
	months[3] = "March"
	months[4] = "April"
	months[5] = "May"
	months[6] = "June"
	months[7] = "July"
	months[8] = "August"
	months[9] = "September"
	months[10] = "October"
	months[11] = "November"
	months[12] = "December"

	// fmt.Println("months", months)

	days := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// if ((((year-1)%4 == 0) && ((year-1)%100 != 0)) || ((year-1)%400 == 0)) && year > initYear {
	// 	blankDays += (year - (year - initYear) + 1) % 7
	// } else if year > initYear {
	// 	blankDays += (year - (year - initYear)) % 7
	// }

	for m := 1; m <= 12; m++ {
		if (((year%4 == 0) && (year%100 != 0)) || (year%400 == 0)) && m == 2 {
			days[m] = 29
		}

		fmt.Println("        ", months[m], " ", year)
		fmt.Println("====================================")
		fmt.Println("  Sun  Mon  Tue  Wed  Thu  Fri  Sat")
		// fmt.Println("initday", blankDays)
		blankDays = (days[m-1] + blankDays) % 7
		// fmt.Println("initday", blankDays)

		for i := 0; i < blankDays; i++ {
			fmt.Print("     ")
		}

		for i := 1; i <= days[m]; i++ {
			fmt.Printf("%5d", i)
			if ((i+blankDays)%7 == 0) || (i == days[m]) {
				fmt.Println()
			}
		}
		fmt.Println()
	}
}
