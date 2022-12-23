package main

import (
	"fmt"
	"strings"

	"github.com/jgudka/advent-of-code-2k22/golang/utils"
)

func main() {
	// fmt.Println(partOne("example"))
	// fmt.Println(partOne("input"))

	fmt.Println(partTwo("example"))
	fmt.Println(partTwo("input"))
}

func partOne(filename string) int {
	input := parse(filename)
	rowCount, colCount := len(input), len(input[0])
	perimeterSize := (rowCount+colCount)*2 - 4
	return perimeterSize + getInnerVisibleCount(input, rowCount, colCount)
}

func partTwo(filename string) int {
	input := parse(filename)
	rowCount, colCount := len(input), len(input[0])
	return getMaxScenicScore(input, rowCount, colCount)
}

func getInnerVisibleCount(grid [][]int, rowCount, colCount int) int {
	visibleCount := 0
	for row := 1; row < rowCount-1; row++ {
		for col := 1; col < colCount-1; col++ {
			if isVisibleLeft(grid[row], col) || isVisibleRight(grid[row], col) ||
				isVisibleTop(grid, row, col) || isVisibleBottom(grid, row, col) {
				visibleCount++
			}
		}
	}
	return visibleCount
}

func getMaxScenicScore(grid [][]int, rowCount, colCount int) int {
	maxScenicScore := getScenicScore(grid, 1, 1)
	for row := 1; row < rowCount-1; row++ {
		for col := 1; col < colCount-1; col++ {
			scenicScore := getScenicScore(grid, row, col)
			// fmt.Printf("Scenic Score row %d, col %d: %d\n", row, col, scenicScore)
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}
	return maxScenicScore
}

func getScenicScore(grid [][]int, row, col int) int {
	return getViewingDistanceLeft(grid[row], col) * getViewingDistanceRight(grid[row], col) *
		getViewingDistanceTop(grid, row, col) * getViewingDistanceBottom(grid, row, col)
}

func getViewingDistanceLeft(row []int, index int) int {
	return getBlockingTreeDistance(row[index], utils.Reverse(row[:index]))
}

func getViewingDistanceRight(row []int, index int) int {
	return getBlockingTreeDistance(row[index], row[index+1:])
}

func getViewingDistanceTop(grid [][]int, row, col int) int {
	topTrees := make([]int, row)
	for i := range topTrees {
		topTrees[i] = grid[i][col]
	}
	return getBlockingTreeDistance(grid[row][col], utils.Reverse(topTrees))
}

func getViewingDistanceBottom(grid [][]int, row, col int) int {
	bottomTrees := make([]int, len(grid)-row-1)
	for i := range bottomTrees {
		bottomTrees[i] = grid[row+i+1][col]
	}
	return getBlockingTreeDistance(grid[row][col], bottomTrees)
}

func getBlockingTreeDistance(tree int, others []int) int {
	for i, otherTree := range others {
		if otherTree >= tree {
			return i + 1
		}
	}
	return len(others)
}

func isVisibleLeft(row []int, index int) bool {
	return isTreeTallerThanOthers(row[index], row[:index])
}

func isVisibleRight(row []int, index int) bool {
	return isTreeTallerThanOthers(row[index], row[index+1:])
}

func isVisibleTop(grid [][]int, row, col int) bool {
	topTrees := make([]int, row)
	for i := range topTrees {
		topTrees[i] = grid[i][col]
	}
	return isTreeTallerThanOthers(grid[row][col], topTrees)
}

func isVisibleBottom(grid [][]int, row, col int) bool {
	bottomTrees := make([]int, len(grid)-row-1)
	for i := range bottomTrees {
		bottomTrees[i] = grid[row+i+1][col]
	}
	return isTreeTallerThanOthers(grid[row][col], bottomTrees)
}

func isTreeTallerThanOthers(tree int, others []int) bool {
	for _, otherTree := range others {
		if otherTree >= tree {
			return false
		}
	}
	return true
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
