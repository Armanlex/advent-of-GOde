package main

import (
	"fmt"
	"strconv"
	"strings"
)

var demo = `5 9 2 8
9 4 7 3
3 8 6 5`

var input string = `ENTER_YOUR_INPUT_HERE`

func main() {
	str := demo
	P2Answer := 0
	for _, line := range strings.Split(str, "\n") { //iterating every input line

		// //                                    | demo needs to use space char " " for delimiter,
		// //                                    v           while user input needs tab character "\t"
		// numberStrings := strings.Split(line, " ")

		numberStrings := strings.Split(line, "\t")

		for index, numStr := range numberStrings { //iterating over every input number

			numA, err := strconv.Atoi(string(numStr))
			if err != nil {
				fmt.Println("some kind of error? 1", err)
			}

			for i2 := index + 1; i2 < len(numberStrings); i2++ { //iterating every input number after the one of the loop above

				numB, err := strconv.Atoi(numberStrings[i2])
				if err != nil {
					fmt.Println("some kind of error? 2", err)
				}

				if numA%numB == 0 { //checking if evenly divisible
					P2Answer += numA / numB

				} else if numB%numA == 0 { //checking if evenly divisible the other way
					P2Answer += numB / numA
				}

			}

		}

	}
	fmt.Println("Part 2:", P2Answer)
}
