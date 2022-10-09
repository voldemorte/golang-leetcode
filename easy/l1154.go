package easy

import (
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	var days = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	var dateSplit = strings.Split(date, "-")
	year, _ := strconv.Atoi(dateSplit[0])
	month, _ := strconv.Atoi(dateSplit[1])
	day, _ := strconv.Atoi(dateSplit[2])

	var result int

	if ((year%4) == 0 && (year%100) != 0) || ((year % 400) == 0) {
		days[1]++ // Increment February if its a leap year
	}

	for i := 0; i < month-1; i++ {
		result = result + days[i]
	}

	return result + day
}
