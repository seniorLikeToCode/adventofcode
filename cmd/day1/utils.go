package main

import "strconv"

func parseNumber(s string) []int {
	res := make([]int, 2)

	res[0], _ = strconv.Atoi(s[0:5])
	res[1], _ = strconv.Atoi(s[8:])

	return res
}
