package urlshort

import (
	"fmt"
	"strconv"
)

func calcDate(userInput string) string {
	yearStr := userInput[4:]
	monthStr := userInput[0:2]
	dayStr := userInput[2:4]
	century := -1
	preciseYear := -1
	monthInt := -1
	dayInt := -1
	first, second, third, fourth := 0, 0, 0, 0
	if i, err := strconv.Atoi(yearStr); err != nil {
		return "incorrect format"
	} else {
		fmt.Print(i)
		centuryStr := yearStr[0:2]
		preciseYearStr := yearStr[2:]
		if j, err := strconv.Atoi(centuryStr); err != nil {
			return "incorrect format"
		} else {
			century = j
		}
		if j, err := strconv.Atoi(preciseYearStr); err != nil {
			return "incorrect format"
		} else {
			preciseYear = j

		}
	}

	if i, err := strconv.Atoi(monthStr); err != nil {
		return "incorrect format"
	} else {
		monthInt = i
	}

	if i, err := strconv.Atoi(dayStr); err != nil {
		return "incorrect format"
	} else {
		dayInt = i
	}

	centuryCode := century % 4
	switch {
	case centuryCode == 0:
		first = 2
	case centuryCode == 1:
		first = 0
	case centuryCode == 2:
		first = 5
	case centuryCode == 3:
		first = 3
	}

	second = preciseYear / 12
	third = preciseYear % 12
	fourth = third / 4

	fingers := first + second + third + fourth
	doomsday := fingers % 7
	dayDay := -1
	switch {
	case monthInt == 1:
		if preciseYear%4 == 0 {
			dayDay = getDay(doomsday, 4, dayInt)
		} else {
			dayDay = getDay(doomsday, 3, dayInt)
		}
	case monthInt == 2:
		if preciseYear%4 == 0 {
			dayDay = getDay(doomsday, 29, dayInt)
		} else {
			dayDay = getDay(doomsday, 28, dayInt)
		}
	case monthInt == 3:
		dayDay = getDay(doomsday, 14, dayInt)
	case monthInt == 4:
		dayDay = getDay(doomsday, 4, dayInt)
	case monthInt == 5:
		dayDay = getDay(doomsday, 9, dayInt)
	case monthInt == 6:
		dayDay = getDay(doomsday, 6, dayInt)
	case monthInt == 7:
		dayDay = getDay(doomsday, 11, dayInt)
	case monthInt == 8:
		dayDay = getDay(doomsday, 8, dayInt)
	case monthInt == 9:
		dayDay = getDay(doomsday, 5, dayInt)
	case monthInt == 10:
		dayDay = getDay(doomsday, 10, dayInt)
	case monthInt == 11:
		dayDay = getDay(doomsday, 7, dayInt)
	case monthInt == 12:
		dayDay = getDay(doomsday, 12, dayInt)

	}
	yay := ""

	switch {
	case dayDay == 0:
		yay = "Sunday"
	case dayDay == 1:
		yay = "Monday"
	case dayDay == 2:
		yay = "Tuesday"
	case dayDay == 3:
		yay = "Wednesday"
	case dayDay == 4:
		yay = "Thursday"
	case dayDay == 5:
		yay = "Friday"
	case dayDay == 6:
		yay = "Saturday"

	}
	return yay
}

func getDay(doomsday int, monthSpec int, dayInt int) int {
	difference := 0
	dayDay := -1

	difference = dayInt - monthSpec
	for difference > 7 {
		difference -= 7
	}
	for difference < 0 {
		difference += 7
	}
	dayDay = doomsday + difference
	for dayDay > 7 {
		dayDay -= 7
	}
	return dayDay
}
