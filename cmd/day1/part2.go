package main

import "fmt"

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
