package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jgudka/advent-of-code-2k22/golang/utils"
)

type monkey struct {
	number         int
	items          []int
	operation      func(old int) int
	testDivisbleBy int
	throwToTrue    int
	throwToFalse   int
	itemsInspected int
}

func main() {
	// fmt.Println(partOne("example"))
	// fmt.Println(partOne("input"))

	fmt.Println(partTwo("example"))
	fmt.Println(partTwo("input"))
}

func partOne(filename string) int {
	monkeys := parse(filename)
	rounds := 20
	for i := 0; i < rounds; i++ {
		simulateRoundOne(monkeys)
	}
	return calculateMonkeyBusiness(monkeys)
}

func partTwo(filename string) int {
	monkeys := parse(filename)
	rounds := 10_000
	lowestCommonMultiple := calculateLowestCommonMultiple(monkeys)
	for i := 0; i < rounds; i++ {
		simulateRoundTwo(monkeys, lowestCommonMultiple)
	}
	for number, monkey := range monkeys {
		fmt.Printf("Monkey %d inspected %d items \n", number, monkey.itemsInspected)
	}
	return calculateMonkeyBusiness(monkeys)
}

func calculateLowestCommonMultiple(monkeys map[int]monkey) int {
	lcm := 1
	for _, monkey := range monkeys {
		lcm *= monkey.testDivisbleBy
	}
	return lcm
}

func simulateRoundOne(monkeyMap map[int]monkey) {
	for i := 0; i < len(monkeyMap); i++ {
		monkey := monkeyMap[i]
		for _, item := range monkey.items {
			newItem := monkey.operation(item) / 3
			processNewItem(monkey, newItem, monkeyMap)
			monkey.itemsInspected++
		}
		monkey.items = []int{}
		monkeyMap[i] = monkey
	}
}

func simulateRoundTwo(monkeyMap map[int]monkey, lowestCommonMultiple int) {
	for i := 0; i < len(monkeyMap); i++ {
		monkey := monkeyMap[i]
		for _, item := range monkey.items {
			newItem := monkey.operation(item) % lowestCommonMultiple
			processNewItem(monkey, newItem, monkeyMap)
			monkey.itemsInspected++
		}
		monkey.items = []int{}
		monkeyMap[i] = monkey
	}
}

func processNewItem(monkey monkey, newItem int, monkeyMap map[int]monkey) {
	var throwToMonkeyNumber int
	if newItem%monkey.testDivisbleBy == 0 {
		throwToMonkeyNumber = monkey.throwToTrue
	} else {
		throwToMonkeyNumber = monkey.throwToFalse
	}
	throwToMonkey := monkeyMap[throwToMonkeyNumber]
	throwToMonkey.items = append(throwToMonkey.items, newItem)
	monkeyMap[throwToMonkeyNumber] = throwToMonkey
}

func calculateMonkeyBusiness(monkeyMap map[int]monkey) int {
	monkeys := make([]monkey, 0, len(monkeyMap))
	for _, monkey := range monkeyMap {
		monkeys = append(monkeys, monkey)
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].itemsInspected > monkeys[j].itemsInspected
	})
	return monkeys[0].itemsInspected * monkeys[1].itemsInspected
}

func parse(inputFileName string) map[int]monkey {
	scanner, _ := utils.GetScanner(inputFileName)
	var monkeyMap = map[int]monkey{}
	var partial []string
	for scanner.Scan() {
		next := scanner.Text()
		if next == "" {
			monkey := parseMonkey(partial)
			monkeyMap[monkey.number] = monkey
			partial = []string{}
		} else {
			partial = append(partial, next)
		}
	}
	monkey := parseMonkey(partial)
	monkeyMap[monkey.number] = monkey
	return monkeyMap
}

func parseMonkey(rawLines []string) monkey {
	return monkey{
		number:         parseNumber(rawLines[0]),
		items:          parseItems(rawLines[1]),
		operation:      parseOperation(rawLines[2]),
		testDivisbleBy: parseTestDivisibleBy(rawLines[3]),
		throwToTrue:    parseThrowTo(rawLines[4]),
		throwToFalse:   parseThrowTo(rawLines[5]),
		itemsInspected: 0,
	}
}

func parseNumber(rawLine string) int {
	return utils.ParseInts(rawLine, 1)[0]
}
func parseItems(rawLine string) []int {
	return utils.ParseInts(rawLine, -1)
}
func parseOperation(rawLine string) func(old int) int {
	substring := strings.SplitAfter(rawLine, "new = old ")[1]
	args := strings.Split(substring, " ")
	operator := args[0]
	rightOperand := args[1]
	if operator == "*" {
		if rightOperand == "old" {
			return func(old int) int {
				return old * old
			}
		} else {
			return func(old int) int {
				return old * utils.StringToInt(rightOperand)
			}
		}
	} else {
		if rightOperand == "old" {
			return func(old int) int {
				return old + old
			}
		} else {
			return func(old int) int {
				return old + utils.StringToInt(rightOperand)
			}
		}
	}
}
func parseTestDivisibleBy(rawLine string) int {
	return utils.ParseInts(rawLine, 1)[0]
}
func parseThrowTo(rawLine string) int {
	return utils.ParseInts(rawLine, 1)[0]
}
