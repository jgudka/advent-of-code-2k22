package main

import (
	"fmt"
	"strings"

	"github.com/jgudka/advent-of-code-2k22/golang/utils"
)

type instruction struct {
	command string
	arg     int
}

const noop, addx = "noop", "addx"

func main() {
	fmt.Println(partOne("example"))
	fmt.Println(partOne("input"))

	printCrt(partTwo("example"))
	fmt.Println()
	printCrt(partTwo("input"))
}

func partOne(filename string) int {
	cyclesOfInterest := []int{20, 60, 100, 140, 180, 220}
	remainingInstructions := parse(filename)
	var currentInstruction instruction
	X := 1
	correspondingXValues := make([]int, len(cyclesOfInterest))
	cyclesOfInterestIndex := 0
	addInProgress := false
	for cycle := 1; cycle <= cyclesOfInterest[len(cyclesOfInterest)-1]; cycle++ {
		if cycle == cyclesOfInterest[cyclesOfInterestIndex] {
			correspondingXValues[cyclesOfInterestIndex] = X
			cyclesOfInterestIndex++
		}
		if addInProgress {
			X += currentInstruction.arg
			addInProgress = false
		} else {
			currentInstruction, remainingInstructions = utils.Pop(remainingInstructions)
			if currentInstruction.command == addx {
				addInProgress = true
			}
		}
	}

	// fmt.Printf("%+v", correspondingXValues)
	return calculateSignalStrengthSum(cyclesOfInterest, correspondingXValues)
}

func partTwo(filename string) []rune {
	remainingInstructions := parse(filename)
	var currentInstruction instruction
	const totalCycles = 240
	crt := make([]rune, totalCycles)
	X := 1
	addInProgress := false
	for cycle := 0; cycle < totalCycles; cycle++ {
		crtXPosition := cycle % 40
		if crtXPosition <= X+1 && crtXPosition >= X-1 {
			crt[cycle] = '#'
		} else {
			crt[cycle] = '.'
		}
		if addInProgress {
			X += currentInstruction.arg
			addInProgress = false
		} else {
			currentInstruction, remainingInstructions = utils.Pop(remainingInstructions)
			if currentInstruction.command == addx {
				addInProgress = true
			}
		}
	}
	return crt
}

func printCrt(pixels []rune) {
	for i, pixel := range pixels {
		fmt.Printf("%c", pixel)
		if (i+1)%40 == 0 {
			fmt.Printf("\n")
		}
	}
}

func calculateSignalStrengthSum(cyclesOfInterest, correspondingXValues []int) int {
	sum := 0
	for i, cycle := range cyclesOfInterest {
		sum += cycle * correspondingXValues[i]
	}
	return sum
}

func parse(inputFileName string) []instruction {
	scanner, _ := utils.GetScanner(inputFileName)
	var result []instruction
	for scanner.Scan() {
		next := scanner.Text()
		rawLine := strings.Split(next, " ")
		instruction := instruction{command: rawLine[0]}
		if len(rawLine) == 2 {
			instruction.arg = utils.StringToInt(rawLine[1])
		}
		result = append(result, instruction)
	}
	return utils.Reverse(result)
}
