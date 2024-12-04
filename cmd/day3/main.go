package main

import "fmt"

func main() {
	data, err := getData()

	if err != nil {
		fmt.Println(err)
		return
	}

	pos := findAllOccurrence(data, "mul")
	do := findAllOccurrence(data, "do()")
	dont := findAllOccurrence(data, "don't()")

	problemA(data, pos, do, dont)
}
