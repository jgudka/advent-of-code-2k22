package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/jgudka/advent-of-code-2k22/golang/utils"
)

type step struct {
	quantity int
	from     int
	to       int
}

const prefixStackLabels = " 1"

func main() {
	fmt.Println(partOne("example"))
	fmt.Println(partOne("input"))

	fmt.Println(partTwo("example"))
	fmt.Println(partTwo("input"))
}

func partOne(filename string) string {
	stacks, steps := parse(filename)
	for _, step := range steps {
		stacks = executeStepPartOne(step, stacks)
	}
	return getTopCratesString(&stacks)
}

func partTwo(filename string) string {
	stacks, steps := parse(filename)
	for _, step := range steps {
		stacks = executeStepPartTwo(step, stacks)
	}
	return getTopCratesString(&stacks)
}

func executeStepPartOne(step step, stacks [][]string) [][]string {
	for i := 0; i < step.quantity; i++ {
		stacks = executeSingleMove(step.from, step.to, stacks)
	}
	return stacks
}

func executeStepPartTwo(step step, stacks [][]string) [][]string {
	var crates []string
	crates, stacks[step.from-1] = utils.PopMultiple(stacks[step.from-1], step.quantity)
	stacks[step.to-1] = append(stacks[step.to-1], crates...)
	return stacks
}

func executeSingleMove(fromStack, toStack int, stacks [][]string) [][]string {
	var crate string
	crate, stacks[fromStack-1] = utils.Pop(stacks[fromStack-1])
	stacks[toStack-1] = append(stacks[toStack-1], crate)
	return stacks
}

func getTopCratesString(stacks *[][]string) string {
	topCratesString := ""
	for _, stack := range *stacks {
		topCratesString += stack[len(stack)-1]
	}
	return topCratesString
}

func parse(inputFileName string) ([][]string, []step) {
	scanner, _ := utils.GetScanner(inputFileName)
	var rawLines []string
	var stackLabelIndex int
	index := 0
	for scanner.Scan() {
		next := scanner.Text()
		rawLines = append(rawLines, next)
		if strings.HasPrefix(next, prefixStackLabels) {
			stackLabelIndex = index
		}
		index++
	}
	crates, stackLabels, steps := rawLines[:stackLabelIndex], rawLines[stackLabelIndex], rawLines[stackLabelIndex+2:]
	stackIndexes := []int{}
	for i, char := range stackLabels {
		if char != ' ' {
			stackIndexes = append(stackIndexes, i)
		}
	}
	return parseStacks(stackIndexes, crates), parseSteps(steps)
}

func parseStacks(stackColIndexes []int, rawCrateLines []string) [][]string {
	stacks := make([][]string, len(stackColIndexes))
	for _, rawCrateLine := range utils.Reverse(rawCrateLines) {
		for stackNumber, stackColIndex := range stackColIndexes {
			crate := rune(rawCrateLine[stackColIndex])
			if utils.IsUpperCase(crate) {
				stacks[stackNumber] = append(stacks[stackNumber], string(crate))
			}
		}
	}
	return stacks
}

func parseSteps(rawSteps []string) []step {
	re := regexp.MustCompile("[0-9]+")
	steps := make([]step, len(rawSteps))
	for i, rawStep := range rawSteps {
		rawStepInfo := re.FindAllString(rawStep, 3)
		steps[i] = step{
			quantity: utils.StringToInt(rawStepInfo[0]),
			from:     utils.StringToInt(rawStepInfo[1]),
			to:       utils.StringToInt(rawStepInfo[2]),
		}
	}
	return steps
}
