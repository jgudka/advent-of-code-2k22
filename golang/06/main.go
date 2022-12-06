package main

import (
	"fmt"
	"strings"

	"github.com/jgudka/advent-of-code-2k22/golang/utils"
)

func main() {
	fmt.Println(partOne("example"))
	fmt.Println(partOne("input"))

	fmt.Println(partTwo("example"))
	fmt.Println(partTwo("input"))
}

func partOne(filename string) []int {
	input, _ := utils.ReadInput(filename)
	markerIndices := make([]int, len(input))
	for i, message := range input {
		markerIndices[i] = calculateStartMarkerIndex(message, 4)
	}
	return markerIndices
}

func partTwo(filename string) []int {
	input, _ := utils.ReadInput(filename)
	markerIndices := make([]int, len(input))
	for i, message := range input {
		markerIndices[i] = calculateStartMarkerIndex(message, 14)
	}
	return markerIndices
}

func calculateStartMarkerIndex(message string, markerLength int) int {
	for i := markerLength; i <= len(message); i++ {
		if areCharactersUnique(message[i-markerLength : i]) {
			return i
		}
	}
	return 0
}

func areCharactersUnique(substring string) bool {
	set := utils.CreateSetFromList(strings.Split(substring, ""))
	return len(substring) == len(set)
}
