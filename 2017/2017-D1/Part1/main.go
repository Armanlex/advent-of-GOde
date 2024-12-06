package main

import (
	"fmt"
)

var demo = []string{"1122", "1111", "1234", "91212129"}
var input string = `ENTER_YOUR_INPUT_HERE`

func main() {
	// number := demo[0]
	number := input
	P1Answer := 0
	for i := 0; i < len(number)-1; i++ { //iterate over each position on the string
		if number[i] == number[i+1] { //if char is the same as the next char
			P1Answer += int(number[i] - '0') //turning char to int
		}
	}
	if number[0] == number[len(number)-1] { //checking if first matches the last one since the sequence is a circle
		P1Answer += int(number[0] - '0')
	}

	fmt.Println("Part 1:", P1Answer)

}
