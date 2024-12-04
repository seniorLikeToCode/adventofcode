package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getData() (string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	str := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		str += line
	}

	return str, nil
}

func findAllOccurrence(s, find string) []int {
	pos := make([]int, 0)
	for i := 0; i < len(s)-2; i++ {
		isMul := s[i : i+3]

		if isMul == find {
			pos = append(pos, i)
		}
	}

	return pos
}

func isValidMul(s string) (int, bool) {
	// we already know it has "mul"
	// check for bracket () , and number

	isValid := true

	// first postions of bracket ()
	bracketPos := []int{-1, -1}

	for i := 0; i < len(s); i++ {
		// fmt.Println(string(s[i]))

		if string(s[i]) == "(" {
			if bracketPos[0] == -1 {
				bracketPos[0] = i
			} else {
				isValid = false
				break
			}
		}

		if string(s[i]) == ")" {
			if bracketPos[0] == -1 {
				isValid = false
			}

			bracketPos[1] = i
			break
		}
	}

	if !isValid || bracketPos[0] == -1 || bracketPos[1] == -1 {
		return 0, false
	}

	text := s[bracketPos[0]+1 : bracketPos[1]]
	if len(text) < 3 {
		return 0, false
	}

	parts := strings.Split(text, ",")

	nums := make([]int, 0)
	for _, part := range parts {
		x, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println(s)
			fmt.Println("Error in changing string to int at", part)
			return 0, false
		}

		nums = append(nums, x)
	}

	if len(nums) != 2 {
		fmt.Println("Numbers are not equal to 2.")
		return 0, false
	}

	return nums[0] * nums[1], true
}
