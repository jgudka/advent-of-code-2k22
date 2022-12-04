package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jgudka/advent-of-code-2k22/golang/utils"
)

func main() {
	fmt.Println(partOne("example"))
	fmt.Println(partOne("input"))

	fmt.Println(partTwo("example"))
	fmt.Println(partTwo("input"))
}

func partOne(filename string) int {
	input := parse(filename)
	containingPairs := getOverlappingPairs(input, isPairContaining)
	return len(containingPairs)
}

func partTwo(filename string) int {
	input := parse(filename)
	containingPairs := getOverlappingPairs(input, isPairOverlapping)
	return len(containingPairs)
}

func getOverlappingPairs(pairs [][]int, overlapFn func(ass1, ass2 []int) bool) [][]int {
	var containingPairs [][]int
	for _, pair := range pairs {
		if overlapFn(pair[:2], pair[2:]) {
			containingPairs = append(containingPairs, pair)
		}
	}
	return containingPairs
}

func isPairContaining(ass1, ass2 []int) bool {
	return (ass1[0] <= ass2[0] && ass1[1] >= ass2[1]) || (ass1[0] >= ass2[0] && ass1[1] <= ass2[1])
}

func isPairOverlapping(ass1, ass2 []int) bool {
	return !(ass1[1] < ass2[0] || ass1[0] > ass2[1])
}

func parse(inputFileName string) [][]int {
	scanner, _ := utils.GetScanner(inputFileName)
	var result [][]int
	for scanner.Scan() {
		next := scanner.Text()
		assignments := strings.FieldsFunc(next, func(r rune) bool {
			return r == '-' || r == ','
		})
		pair := []int{}
		for _, time := range assignments {
			number, _ := strconv.Atoi(time)
			pair = append(pair, number)
		}
		result = append(result, pair)
	}
	return result
}
