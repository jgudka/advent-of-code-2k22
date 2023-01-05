package main

import (
	"fmt"

	"github.com/jgudka/advent-of-code-2k22/golang/utils"
)

const top, bottom, left, right = "top", "bottom", "left", "right"

type grid struct {
	squares [][]square
	height  int
	width   int
	start   *square
	end     *square
}

type square struct {
	elevation      rune
	x              int
	y              int
	isEnd          bool
	visited        bool
	routeChar      rune
	grid           *grid
	fromDirection  string
	previousSquare *square
}

func main() {
	// fmt.Println(partOne("example"))
	// fmt.Println(partOne("input"))

	fmt.Println(partTwo("example"))
	fmt.Println(partTwo("input"))
}

func partOne(filename string) int {
	grid := parse(filename)
	shortestRoute := grid.getShortestRoute()
	grid.printRoute()
	return shortestRoute
}

func partTwo(filename string) int {
	grid := parse(filename)
	mostDirectGrid := grid
	mostDirectRoute := grid.getShortestRoute()
	for _, row := range grid.squares {
		for _, square := range row {
			if square.elevation == 'a' {
				grid := parse(filename)
				grid.start = grid.getSquareAt(square.x, square.y)
				shortestRoute := grid.getShortestRoute()
				if shortestRoute > 0 && shortestRoute < mostDirectRoute {
					mostDirectRoute = shortestRoute
					mostDirectGrid = grid
				}
			}
		}
	}
	mostDirectGrid.printRoute()
	return mostDirectRoute
}

func (g *grid) getShortestRoute() int {
	routeLength := 0
	endReached := false
	g.start.visited = true
	currentSquares, nextSquares := []*square{g.start}, []*square{}
	for !endReached {
		for _, currentSquare := range currentSquares {
			candidateNextSquares := []*square{currentSquare.moveUp(), currentSquare.moveDown(), currentSquare.moveLeft(), currentSquare.moveRight()}
			for _, candidate := range candidateNextSquares {
				if candidate != nil {
					nextSquares = append(nextSquares, candidate)
					if candidate.isEnd {
						endReached = true
					}
				}
			}
		}
		if len(nextSquares) == 0 {
			return -1
		}
		currentSquares, nextSquares = nextSquares, []*square{}
		routeLength++
	}

	return routeLength
}

func (g *grid) printRoute() {
	currentSquare := g.getSquareAt(g.end.x, g.end.y)
	for currentSquare != g.start {
		if currentSquare.fromDirection == top {
			currentSquare.previousSquare.routeChar = 'v'
		} else if currentSquare.fromDirection == bottom {
			currentSquare.previousSquare.routeChar = '^'
		} else if currentSquare.fromDirection == left {
			currentSquare.previousSquare.routeChar = '>'
		} else if currentSquare.fromDirection == right {
			currentSquare.previousSquare.routeChar = '<'
		}
		currentSquare = currentSquare.previousSquare
	}

	for _, row := range g.squares {
		for _, square := range row {
			fmt.Printf("%c", square.routeChar)
		}
		fmt.Printf("\n")
	}
}

func (g *grid) getSquareAt(x, y int) *square {
	rowIndex, colIndex := g.height-1-y, x
	return &g.squares[rowIndex][colIndex]
}

func (s *square) moveUp() *square {
	if s.y == s.grid.height-1 {
		return nil
	}
	topSquare := s.grid.getSquareAt(s.x, s.y+1)
	if s.canMoveTo(topSquare) {
		topSquare.visited = true
		topSquare.fromDirection = bottom
		topSquare.previousSquare = s
		return topSquare
	} else {
		return nil
	}
}

func (s *square) moveDown() *square {
	if s.y == 0 {
		return nil
	}
	bottomSquare := s.grid.getSquareAt(s.x, s.y-1)
	if s.canMoveTo(bottomSquare) {
		bottomSquare.visited = true
		bottomSquare.fromDirection = top
		bottomSquare.previousSquare = s
		return bottomSquare
	} else {
		return nil
	}
}

func (s *square) moveLeft() *square {
	if s.x == 0 {
		return nil
	}
	leftSquare := s.grid.getSquareAt(s.x-1, s.y)
	if s.canMoveTo(leftSquare) {
		leftSquare.visited = true
		leftSquare.fromDirection = right
		leftSquare.previousSquare = s
		return leftSquare
	} else {
		return nil
	}
}

func (s *square) moveRight() *square {
	if s.x == s.grid.width-1 {
		return nil
	}
	rightSquare := s.grid.getSquareAt(s.x+1, s.y)
	if s.canMoveTo(rightSquare) {
		rightSquare.visited = true
		rightSquare.fromDirection = left
		rightSquare.previousSquare = s
		return rightSquare
	} else {
		return nil
	}
}

func (s *square) canMoveTo(targetSquare *square) bool {
	return !targetSquare.visited && targetSquare.elevation-s.elevation <= 1
}

func parse(inputFileName string) *grid {
	scanner, _ := utils.GetScanner(inputFileName)
	rawInput := []string{}
	for scanner.Scan() {
		rawInput = append(rawInput, scanner.Text())
	}
	grid := grid{}
	height := len(rawInput)
	squares := make([][]square, height)
	for rowIndex, line := range rawInput {
		row := make([]square, len(line))
		for colIndex, char := range line {
			elevation := char
			isEnd := false
			routeChar := '.'
			if char == 'S' {
				elevation = 'a'
			} else if char == 'E' {
				elevation = 'z'
				isEnd = true
				routeChar = 'E'
			}
			square := square{
				elevation: elevation,
				x:         colIndex,
				y:         height - 1 - rowIndex,
				isEnd:     isEnd,
				visited:   false,
				routeChar: routeChar,
				grid:      &grid,
			}
			row[colIndex] = square
			if char == 'S' {
				grid.start = &square
			} else if char == 'E' {
				grid.end = &square
			}
		}
		squares[rowIndex] = row
	}
	grid.squares = squares
	grid.height, grid.width = height, len(squares[0])
	return &grid
}
