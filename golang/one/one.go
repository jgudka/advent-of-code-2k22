package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/jgudka/advent-of-code-2k22/golang/utils"
)

func main() {
	input := parse("one/input")
	// fmt.Println(input)
	initial := input[:3]
	initialSums := make([]int, 3)
	for i, calories := range initial {
		initialSums[i] = sum(calories)
	}
	sort.Slice(initialSums, func(i, j int) bool {
		return initialSums[i] > initialSums[j]
	})

	first, second, third := initialSums[0], initialSums[1], initialSums[2]
	for i := 3; i < len(input); i++ {
		totalCalories := sum(input[i])
		if totalCalories > first {
			first, second, third = totalCalories, first, second
		} else if totalCalories > second {
			second, third = totalCalories, second
		} else if totalCalories > third {
			third = totalCalories
		}
	}
	fmt.Println(sum([]int{first, second, third}))
}

func sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func parse(inputFileName string) [][]int {
	scanner, _ := utils.GetScanner(inputFileName)
	var result [][]int
	var partial []int
	for scanner.Scan() {
		next := scanner.Text()
		if next == "" {
			result = append(result, partial)
			partial = []int{}
		} else {
			number, _ := strconv.Atoi(next)
			partial = append(partial, number)
		}
	}
	result = append(result, partial)
	return result
}
