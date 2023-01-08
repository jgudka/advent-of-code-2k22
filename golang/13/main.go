package main

import (
	"fmt"
	"sort"

	"github.com/jgudka/advent-of-code-2k22/golang/utils"
)

const separator, startList, endList = ',', '[', ']'
const leftSide, rightSide, same = "left", "right", "same"
const bothInts, bothLists, leftInt, rightInt = "bothInts", "bothLists", "leftInt", "rightInt"
const rightOrder, wrongOrder, noResult = "rightOrder", "wrongOrder", "noResult"

type packetPair struct {
	left, right []interface{}
}

func main() {
	fmt.Println(partOne("example"))
	fmt.Println(partOne("input"))

	fmt.Println(partTwo("example_part_2"))
	fmt.Println(partTwo("input_part_2"))
}

func partOne(filename string) int {
	input := parseOne(filename)
	// fmt.Printf("%+v\n", input)
	rightOrderIndices := []int{}
	for i, pair := range input {
		if compareLists(pair.left, pair.right) == rightOrder {
			rightOrderIndices = append(rightOrderIndices, i+1)
		}
	}
	return utils.Sum(rightOrderIndices)
}

func partTwo(filename string) int {
	packets := parseTwo(filename)
	// fmt.Printf("%+v\n", input)
	sort.Slice(packets, func(i, j int) bool {
		order := compareLists(packets[i], packets[j])
		return order == rightOrder
	})
	return utils.Product(getDividerIndices(packets))
}

func getDividerIndices(packets [][]interface{}) []int {
	dividerIndices := []int{}
	isDivider := func(packet []interface{}) bool {
		if len(packet) == 1 && utils.IsSlice(packet[0]) {
			firstEl := utils.InterfaceToSlice(packet[0])
			return len(firstEl) == 1 && (firstEl[0] == 2 || firstEl[0] == 6)
		}
		return false
	}
	for i, packet := range packets {
		if isDivider(packet) {
			dividerIndices = append(dividerIndices, i+1)
		}
	}
	return dividerIndices
}

func compareLists(left, right []interface{}) string {
	smaller := leftSide
	minSize := len(left)
	if len(left) > len(right) {
		smaller = rightSide
		minSize = len(right)
	} else if len(left) == len(right) {
		smaller = same
	}
	for i := 0; i < minSize; i++ {
		leftEl, rightEl := left[i], right[i]
		types := getTypes(leftEl, rightEl)
		switch types {
		case bothLists:
			result := compareLists(utils.InterfaceToSlice(leftEl), utils.InterfaceToSlice(rightEl))
			if result != noResult {
				return result
			}
		case bothInts:
			leftInt, rightInt := utils.InterfaceToInt(leftEl), utils.InterfaceToInt(rightEl)
			if leftInt < rightInt {
				return rightOrder
			} else if leftInt > rightInt {
				return wrongOrder
			}
		case leftInt:
			leftList := []interface{}{}
			leftList = append(leftList, leftEl)
			result := compareLists(leftList, utils.InterfaceToSlice(rightEl))
			if result != noResult {
				return result
			}
		case rightInt:
			rightList := []interface{}{}
			rightList = append(rightList, rightEl)
			result := compareLists(utils.InterfaceToSlice(leftEl), rightList)
			if result != noResult {
				return result
			}
		}
	}

	if smaller == leftSide {
		return rightOrder
	} else if smaller == rightSide {
		return wrongOrder
	} else {
		return noResult
	}
}

func getTypes(left, right interface{}) string {
	leftIsSlice, rightIsSlice := utils.IsSlice(left), utils.IsSlice(right)
	if leftIsSlice && rightIsSlice {
		return bothLists
	} else if !leftIsSlice && !rightIsSlice {
		return bothInts
	} else if leftIsSlice {
		return rightInt
	} else {
		return leftInt
	}
}

func parseOne(inputFileName string) []packetPair {
	scanner, _ := utils.GetScanner(inputFileName)
	var result []packetPair
	var partial []string
	for scanner.Scan() {
		next := scanner.Text()
		if next == "" {
			packetPair := packetPair{left: parseList(partial[0]), right: parseList(partial[1])}
			result = append(result, packetPair)
			partial = []string{}
		} else {
			partial = append(partial, next)
		}
	}
	packetPair := packetPair{left: parseList(partial[0]), right: parseList(partial[1])}
	result = append(result, packetPair)
	return result
}

func parseTwo(inputFileName string) [][]interface{} {
	scanner, _ := utils.GetScanner(inputFileName)
	var result [][]interface{}
	for scanner.Scan() {
		next := scanner.Text()
		if next != "" {
			result = append(result, parseList(next))
		}
	}
	return result
}

func parseList(rawList string) []interface{} {
	rawListChars := []rune(rawList)
	list := []interface{}{}
	if rawListChars[0] == startList && rawListChars[1] == endList {
		return list
	}
	i := 1
	partial := ""
	for i < len(rawListChars) {
		if rawListChars[i] == startList {
			endIndex := findEndIndex(rawList[i:])
			list = append(list, parseList(rawList[i:i+endIndex+1]))
			i += endIndex + 2
		} else if rawListChars[i] == separator || rawListChars[i] == endList {
			list = append(list, utils.StringToInt(partial))
			i++
			partial = ""
		} else {
			partial += string(rawListChars[i])
			i++
		}
	}

	return list
}

func findEndIndex(rawList string) int {
	bracketCount := 0
	for i, char := range rawList {
		if char == startList {
			bracketCount++
		} else if char == endList {
			bracketCount--
		}
		if bracketCount == 0 {
			return i
		}
	}
	return 0
}
