package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	data, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening the file.")
		return
	}

	defer data.Close()
	scanner := bufio.NewScanner(data)

	A := make([]int, 0)
	B := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		res := parseNumber(line)

		A = append(A, res[0])
		B = append(B, res[1])
	}

	sort.Ints(A)
	sort.Ints(B)

	// problemA(A, B)
	problemB(A, B)
}
