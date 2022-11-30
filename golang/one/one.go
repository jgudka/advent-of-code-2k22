package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := readInput("one/example")
	fmt.Println(input[2])
}

func readInput(inputFileName string) ([]string, error) {
	file, err := os.Open(fmt.Sprintf("%s.txt", inputFileName))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// number, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
