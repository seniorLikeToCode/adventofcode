package main

import "fmt"

func problemB(matrix [][]int) {
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
			continue
		}

		// remove one element to make good array
		for i := 0; i < len(row); i++ {
			temp := make([]int, 0)

			for j := 0; j < len(row); j++ {
				if i == j {
					continue
				} else {
					temp = append(temp, row[j])
				}
			}

			if isGood(temp) {
				answer++
				break
			}
		}
	}

	fmt.Println(answer)
}
