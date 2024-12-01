package main

import (
	"fmt"
)

var demo = []string{"1212", "1221", "123425", "123123", "12131415"}

var input string = `ENTER_YOUR_INPUT_HERE`

func main() {
	// number := demo[4]
	number := input
	P2Answer := 0
	halfLen := len(number) / 2
	size := len(number)
	for i := 0; i < len(number); i++ { //iterate over each position on the string
		a := number[i]
		b := number[(i+halfLen)%size]
		if a == b { //if char is the same as the next char
			P2Answer += int(number[i] - '0') //turning char to int
		}
	}

	fmt.Println("Part 2:", P2Answer)

}
