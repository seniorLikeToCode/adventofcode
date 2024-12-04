package main

import "fmt"

func problemA(data string, pos, do, dont []int) {
	ans := 0
	for i := 0; i < len(pos); i++ {
		text := ""
		if i+1 == len(pos) {
			text = data[pos[i]:]
		} else {
			text = data[pos[i]:pos[i+1]]
		}

		// fmt.Println(text)
		val, isValid := isValidMul(text)
		if isValid {
			ans += val
		}
	}

	fmt.Println(ans)
}
