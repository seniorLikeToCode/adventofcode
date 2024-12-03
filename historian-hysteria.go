package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func parseNumber(s string) []int {
	res := make([]int, 2)

	res[0], _ = strconv.Atoi(s[0:5])
	res[1], _ = strconv.Atoi(s[8:])

	return res
}

func problemA(A, B []int) {
	N, answer := len(A), 0

	for i := 0; i < N; i++ {
		x := A[i] - B[i]
		if x < 0 {
			x *= -1
		}
		answer += x
	}

	fmt.Println(answer)
}

func problemB(A, B []int) {
	N, answer := len(A), 0

	mp := make(map[int]int)
	for i := 0; i < N; i++ {
		mp[B[i]]++
	}

	for i := 0; i < N; i++ {
		answer += A[i] * mp[A[i]]
	}

	fmt.Println(answer)
}

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
