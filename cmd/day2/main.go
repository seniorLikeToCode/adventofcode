package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	matrix := make([][]int, 0)

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split the line into fields
		line := scanner.Text()
		strNums := strings.Fields(line) // Split by whitespace

		// Convert strings to integers
		var row []int
		for _, str := range strNums {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				return
			}
			row = append(row, num)
		}

		matrix = append(matrix, row)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// problemA(matrix)
	problemB(matrix)
}
