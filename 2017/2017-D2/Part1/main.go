package main

import (
	"fmt"
	"strconv"
	"strings"
)

var demo = `5 1 9 5
7 5 3
2 4 6 8`

var input string = `ENTER_YOUR_INPUT_HERE`

func main() {
	str := input
	P1Answer := 0
	for _, line := range strings.Split(str, "\n") {
		smallest := 999999999
		largest := -999999999
		//												  v demo needs to use space char for delimiter, while user input needs tab character
		for _, char := range strings.Split(line, string('\t')) {
			num, err := strconv.Atoi(char)
			if err != nil {
				fmt.Println("some kind of error? ", err)
			}
			if smallest > num {
				smallest = num
			}
			if largest < num {
				largest = num
			}
		}
		P1Answer += largest - smallest

	}
	fmt.Println("Part 1:", P1Answer)
}
