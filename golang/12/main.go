package main

import (
	"fmt"
	"strings"

	"github.com/jgudka/advent-of-code-2k22/golang/utils"
)

type square struct {
	elevation rune
	x         int
	y         int
	isEnd     bool
	visited   bool
	routeChar rune
}

func main() {
	fmt.Println(partOne("example"))
	fmt.Println(partOne("input"))

	// fmt.Println(partTwo("example"))
	// fmt.Println(partTwo("input"))
}

func partOne(filename string) int {
	input := parse(filename)
	return 1
}

func partTwo(filename string) int {
	input := parse(filename)

	return 1
}

func parse(inputFileName string) [][]int {
	scanner, _ := utils.GetScanner(inputFileName)
	var result [][]int
	for scanner.Scan() {
		next := scanner.Text()
		row := make([]int, len(next))
		for i, numberString := range strings.Split(next, "") {
			row[i] = utils.StringToInt(numberString)
		}
		result = append(result, row)
	}
	return result
}
