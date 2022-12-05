package utils

import "strconv"

func StringToInt(numberString string) int {
	number, _ := strconv.Atoi(numberString)
	return number
}
