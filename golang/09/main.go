package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/jgudka/advent-of-code-2k22/golang/utils"
)

type point struct {
	x int
	y int
}

type motion struct {
	direction string
	distance  int
}

const up, down, left, right = "U", "D", "L", "R"
const upLeft, upRight, downLeft, downRight = "UL", "UR", "DL", "DR"

var movementVectors = map[string][]int{
	up: {0, 1}, down: {0, -1}, left: {-1, 0}, right: {1, 0},
	upLeft: {-1, 1}, upRight: {1, 1}, downLeft: {-1, -1}, downRight: {1, -1},
}

var tailPositionMap = make(map[point]struct{})
var tailPositionMapTwo = make(map[point]struct{})

func main() {
	// fmt.Println(partOne("example"))
	// fmt.Println(partOne("input"))

	// fmt.Println(partTwo("example"))
	// fmt.Println(partTwo("example_two"))
	fmt.Println(partTwo("input"))
}

func partOne(filename string) int {
	input := parse(filename)
	head, tail := point{x: 0, y: 0}, point{x: 0, y: 0}
	for _, motion := range input {
		head, tail = executeMotionOne(head, tail, motion)
	}

	// fmt.Printf("%+v", tailPositionMap)
	return len(tailPositionMap)
}

func partTwo(filename string) int {
	input := parse(filename)
	knots := make([]point, 10)
	for i := 0; i < 10; i++ {
		knots[i] = point{x: 0, y: 0}
	}
	for _, motion := range input {
		knots = executeMotionTwo(knots, motion)
	}
	return len(tailPositionMapTwo)
}

func executeMotionTwo(knots []point, motion motion) []point {
	for i := 0; i < motion.distance; i++ {
		knots[0] = moveKnot(knots[0], motion.direction)
		for j := 1; j < len(knots); j++ {
			knots[j] = moveTail(knots[j-1], knots[j])
		}
		tailPositionMapTwo[knots[9]] = struct{}{}
	}
	return knots
}

func executeMotionOne(head, tail point, motion motion) (point, point) {
	for i := 0; i < motion.distance; i++ {
		head = moveKnot(head, motion.direction)
		tail = moveTail(head, tail)
		tailPositionMap[tail] = struct{}{}
	}
	return head, tail
}

func moveTail(head, tail point) point {
	if isTouching(head, tail) {
		return tail
	}
	if head.y == tail.y {
		if head.x > tail.x {
			return moveKnot(tail, right)
		} else {
			return moveKnot(tail, left)
		}
	} else if head.x == tail.x {
		if head.y > tail.y {
			return moveKnot(tail, up)
		} else {
			return moveKnot(tail, down)
		}
	} else {
		if head.x > tail.x {
			if head.y > tail.y {
				return moveKnot(tail, upRight)
			} else {
				return moveKnot(tail, downRight)
			}
		} else {
			if head.y > tail.y {
				return moveKnot(tail, upLeft)
			} else {
				return moveKnot(tail, downLeft)
			}
		}
	}
}

func isTouching(p1, p2 point) bool {
	return math.Abs(float64(p1.x-p2.x)) <= 1 && math.Abs(float64(p1.y-p2.y)) <= 1
}

func moveKnot(startPoint point, direction string) point {
	vector := movementVectors[direction]
	startPoint.x += vector[0]
	startPoint.y += vector[1]
	return startPoint
}

func parse(inputFileName string) []motion {
	scanner, _ := utils.GetScanner(inputFileName)
	var result []motion
	for scanner.Scan() {
		next := scanner.Text()
		rawLine := strings.Split(next, " ")
		result = append(result,
			motion{direction: rawLine[0], distance: utils.StringToInt(rawLine[1])},
		)
	}
	return result
}
