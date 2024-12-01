package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

func main() {

	intSlice := []int{}

	for _, inst := range strings.Split(input, "\n") {
		if inst != "" {

			val, err := strconv.Atoi(inst)
			if err != nil {
				fmt.Println("Atoi error!")
			}
			intSlice = append(intSlice, val)
		}
	}

	p2answer := 0
	pointer := 0

	for {

		p2answer++
		if intSlice[pointer] >= 3 { //Part 2 change        ->                        V
			pointer, intSlice[pointer] = pointer+intSlice[pointer], intSlice[pointer]-1
		} else { //                                                                  V
			pointer, intSlice[pointer] = pointer+intSlice[pointer], intSlice[pointer]+1
		}

		if pointer < 0 {
			pointer += len(intSlice)
		} else if pointer >= len(intSlice) {
			pointer -= len(intSlice)
			fmt.Println("Part 2:", p2answer)
			return
		}

	}

}
