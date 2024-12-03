package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isIncreasing(A []int) bool {
	good := true

	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			good = false
		}
	}

	return good
}

func isDecreasing(A []int) bool {
	good := true

	for i := 0; i < len(A)-1; i++ {
		if A[i] < A[i+1] {
			good = false
		}
	}

	return good
}

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

	answer := 0
	for _, row := range matrix {
		// checking for increasing
		if isIncreasing(row) {
			good := true
			for i := 0; i < len(row)-1; i++ {
				if row[i+1]-row[i] > 3 {
					good = false
				}

				if row[i+1]-row[i] < 1 {
					good = false
				}
			}

			if good {
				answer++
			}
		}

		// checking for decreasing
		if isDecreasing(row) {
			good := true
			for i := 0; i < len(row)-1; i++ {
				if row[i]-row[i+1] > 3 {
					good = false
				}

				if row[i]-row[i+1] < 1 {
					good = false
				}
			}

			if good {
				answer++
			}
		}

	}

	fmt.Println(answer)
}
