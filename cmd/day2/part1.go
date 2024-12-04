package main

import "fmt"

func problemA(matrix [][]int) {
	// increasing

	isGood := func(t []int) bool {
		good := true
		for i := 1; i < len(t); i++ {
			if t[i-1] > t[i] {
				good = false
			}

			if t[i]-t[i-1] < 1 {
				good = false
			}

			if t[i]-t[i-1] > 3 {
				good = false
			}
		}

		if good {
			return true
		}

		good = true
		// decreasing

		for i := 1; i < len(t); i++ {
			if t[i-1] < t[i] {
				good = false
			}

			if t[i-1]-t[i] < 1 {
				good = false
			}

			if t[i-1]-t[i] > 3 {
				good = false
			}
		}

		return good
	}

	answer := 0
	for _, row := range matrix {
		if isGood(row) {
			answer++
		}
	}

	fmt.Println(answer)
}
