package main

import (
	"fmt"

	"github.com/jgudka/advent-of-code-2k22/golang/utils"
)

const rock, paper, scissors string = "Rock", "Paper", "Scissors"
const lose, draw, win string = "Lose", "Draw", "Win"

var shapeKeyMap = map[string]string{"A": rock, "B": paper, "C": scissors}
var roundOutcomeMap = map[string]string{"X": lose, "Y": draw, "Z": win}
var shapeScoreMap = map[string]int{rock: 1, paper: 2, scissors: 3}
var outcomeScoreMap = map[string]int{lose: 0, draw: 3, win: 6}

func main() {
	input := parse("input")

	score := 0
	for _, round := range input {
		opp, outcome := round[0], round[1]
		score += getShapeScore(opp, outcome) + outcomeScoreMap[outcome]
	}
	fmt.Println(score)
}

func getShapeScore(opp, outcome string) int {
	var me string
	if outcome == draw {
		me = opp
	} else if opp == rock {
		if outcome == win {
			me = paper
		} else {
			me = scissors
		}
	} else if opp == paper {
		if outcome == win {
			me = scissors
		} else {
			me = rock
		}
	} else if opp == scissors {
		if outcome == win {
			me = rock
		} else {
			me = paper
		}
	}
	return shapeScoreMap[me]
}

func parse(inputFileName string) [][]string {
	scanner, _ := utils.GetScanner(inputFileName)
	var result [][]string
	for scanner.Scan() {
		next := scanner.Text()
		result = append(result, []string{shapeKeyMap[next[0:1]], roundOutcomeMap[next[2:3]]})
	}
	return result
}
