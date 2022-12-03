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

func partOne(filename string) int {
	input := parseOne(filename)
	var sharedItemTypes = getSharedItemTypes(input)
	return utils.Sum(getPriorities(sharedItemTypes))
}

func partTwo(filename string) int {
	input := parseTwo(filename)
	var badgeTypes = getBadgeTypes(input)
	return utils.Sum(getPriorities(badgeTypes))
}

func getSharedItemTypes(rucksackContents [][]string) []rune {
	var sharedItemTypes []rune
	for _, contents := range rucksackContents {
		firstCompartment, secondCompartment := contents[0], contents[1]
		for _, item := range secondCompartment {
			if strings.ContainsRune(firstCompartment, item) {
				sharedItemTypes = append(sharedItemTypes, item)
				break
			}
		}
	}
	return sharedItemTypes
}

func getBadgeTypes(rucksackContents [][]string) []rune {
	var badgeTypes []rune
	for _, contents := range rucksackContents {
		first, second, third := contents[0], contents[1], contents[2]
		for _, item := range first {
			if strings.ContainsRune(second, item) && strings.ContainsRune(third, item) {
				badgeTypes = append(badgeTypes, item)
				break
			}
		}
	}
	return badgeTypes
}

func getPriorities(sharedItemTypes []rune) []int {
	var priorities []int
	isLowerCase := func(item rune) bool {
		return item >= 'a'
	}
	for _, item := range sharedItemTypes {
		if isLowerCase(item) {
			priorities = append(priorities, int(item-'a'+1))
		} else {
			priorities = append(priorities, int(item-'A'+27))
		}
	}
	return priorities
}

func parseOne(inputFileName string) [][]string {
	scanner, _ := utils.GetScanner(inputFileName)
	var result [][]string
	for scanner.Scan() {
		next := scanner.Text()
		compartmentSize := len(next) / 2
		result = append(result, []string{next[:compartmentSize], next[compartmentSize:]})
	}
	return result
}
func parseTwo(inputFileName string) [][]string {
	scanner, _ := utils.GetScanner(inputFileName)
	var result [][]string
	var group []string
	lineCount := 0
	for scanner.Scan() {
		lineCount++
		group = append(group, scanner.Text())
		if lineCount%3 == 0 {
			result = append(result, group)
			group = []string{}
		}
	}
	return result
}
