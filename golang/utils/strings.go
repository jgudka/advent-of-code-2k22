package utils

import (
	"math/big"
	"regexp"
	"strconv"
)

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

func ParseInts(rawText string, maxMatches int) []int {
	re := regexp.MustCompile("[0-9]+")
	stringInts := re.FindAllString(rawText, maxMatches)
	integers := make([]int, len(stringInts))
	for i, stringInt := range stringInts {
		integers[i] = StringToInt(stringInt)
	}
	return integers
}
func ParseBigInts(rawText string, maxMatches int) []big.Int {
	re := regexp.MustCompile("[0-9]+")
	stringInts := re.FindAllString(rawText, maxMatches)
	bigIntegers := make([]big.Int, len(stringInts))
	for i, stringInt := range stringInts {
		parsed, _ := new(big.Int).SetString(stringInt, 10)
		bigIntegers[i] = *parsed

	}
	return bigIntegers
}
