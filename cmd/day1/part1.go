package main

import "fmt"

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
