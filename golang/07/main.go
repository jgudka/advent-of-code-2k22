package main

import (
	"fmt"
	"strings"

	"github.com/jgudka/advent-of-code-2k22/golang/utils"
)

type fileNode struct {
	name string
	size int
}

type directoryNode struct {
	name            string
	directFilesSize int
	totalSize       int
	parent          *directoryNode
	children        []*directoryNode
	files           []fileNode
}

func main() {
	// fmt.Println(partOne("example"))
	// fmt.Println(partOne("input"))

	fmt.Println(partTwo("example"))
	fmt.Println(partTwo("input"))
}

func partOne(filename string) int {
	directoryRoot := parse(filename)
	_, smallDirectories :=
		calculateTotalSizeAndGetDirectoriesMatchingSizeCriteriaRecursive(
			directoryRoot,
			func(totalSize int) bool { return totalSize <= 100_000 },
		)
	return sumDirectorySizes(smallDirectories)
}

func partTwo(filename string) int {
	directoryRoot := parse(filename)

	const totalSpace, requiredSpace = 70_000_000, 30_000_000
	usedSpace, _ :=
		calculateTotalSizeAndGetDirectoriesMatchingSizeCriteriaRecursive(
			directoryRoot,
			func(totalSize int) bool { return false },
		)

	minSpaceToClear := requiredSpace - (totalSpace - usedSpace)
	_, directoriesToClear :=
		calculateTotalSizeAndGetDirectoriesMatchingSizeCriteriaRecursive(
			directoryRoot,
			func(totalSize int) bool { return totalSize >= minSpaceToClear })
	return getSmallestDirectorySize(directoriesToClear)
}

func calculateTotalSizeAndGetDirectoriesMatchingSizeCriteriaRecursive(node *directoryNode, sizePredicate func(int) bool) (int, []directoryNode) {
	totalSize := node.directFilesSize
	smallDirectories := []directoryNode{}
	for _, child := range node.children {
		childTotalSize, childSmallDirectories := calculateTotalSizeAndGetDirectoriesMatchingSizeCriteriaRecursive(child, sizePredicate)
		totalSize += childTotalSize
		smallDirectories = append(smallDirectories, childSmallDirectories...)
	}
	node.totalSize = totalSize
	if sizePredicate(totalSize) {
		smallDirectories = append(smallDirectories, *node)
	}
	return totalSize, smallDirectories
}

func sumDirectorySizes(directories []directoryNode) int {
	size := 0
	for _, directory := range directories {
		size += directory.totalSize
	}
	return size
}

func getSmallestDirectorySize(directories []directoryNode) int {
	size := directories[0].totalSize
	for _, directory := range directories {
		if directory.totalSize < size {
			size = directory.totalSize
		}
	}
	return size
}

const commandPrefix = "$"
const cd, ls = "cd", "ls"
const directoryKeyword = "dir"

func parse(inputFileName string) *directoryNode {
	scanner, _ := utils.GetScanner(inputFileName)
	var commands [][]string
	for scanner.Scan() {
		next := scanner.Text()
		commands = append(commands, strings.Split(next, " "))
	}
	rootNode := directoryNode{name: "/"}
	currentNode := &rootNode
	for _, command := range commands {
		if command[0] == commandPrefix {
			if command[1] == cd {
				currentNode.directFilesSize = calculateDirectFileSizes(currentNode.files)
				currentNode = changeDirectory(currentNode, &rootNode, command[2])
			}
		} else {
			if command[0] == directoryKeyword {
				currentNode.children = append(currentNode.children, &directoryNode{name: command[1], parent: currentNode})
			} else {
				currentNode.files = append(currentNode.files, fileNode{size: utils.StringToInt(command[0]), name: command[1]})
			}
		}
	}
	currentNode.directFilesSize = calculateDirectFileSizes(currentNode.files)

	return &rootNode
}

func changeDirectory(currentNode, rootNode *directoryNode, target string) *directoryNode {
	if target == "/" {
		return rootNode
	} else if target == ".." {
		return currentNode.parent
	} else {
		for _, child := range currentNode.children {
			if child.name == target {
				return child
			}
		}
		return currentNode
	}
}

func calculateDirectFileSizes(files []fileNode) int {
	size := 0
	for _, file := range files {
		size += file.size
	}
	return size
}
