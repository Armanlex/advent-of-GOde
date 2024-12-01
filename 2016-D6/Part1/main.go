package main

import (
	"fmt"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

func main() {

	counter := [26][8]int{}

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		for i, r := range line {
			counter[int(r-'a')][i]++
		}

	}
	word := ""

	for i := 0; i < 8; i++ {
		largest := 0
		tempRune := '1'
		for i2 := 0; i2 < 26; i2++ {
			if largest < counter[i2][i] {
				tempRune = rune(i2) + 'a'
				largest = counter[i2][i]
			}
		}
		word += string(tempRune)
	}

	fmt.Println("Part 1: ", word)

}
