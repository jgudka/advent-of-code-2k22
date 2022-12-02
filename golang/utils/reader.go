package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetScanner(inputFileName string) (*bufio.Scanner, error) {
	file, err := os.Open(fmt.Sprintf("%s.txt", inputFileName))
	if err != nil {
		return nil, err
	}
	// defer file.Close()

	scanner := bufio.NewScanner(file)
	return scanner, nil
}

func ReadInput(inputFileName string) ([]string, error) {
	var lines []string

	scanner, _ := GetScanner(inputFileName)
	for scanner.Scan() {
		// number, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
