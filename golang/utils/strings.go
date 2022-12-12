package utils

import "strconv"

func StringToInt(numberString string) int {
	number, _ := strconv.Atoi(numberString)
	return number
}

func StringIsInt(possibleNumberString string) bool {
	if _, err := strconv.Atoi(possibleNumberString); err == nil {
		return true
	}
	return false
}
