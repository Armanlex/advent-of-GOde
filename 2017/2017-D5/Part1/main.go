package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

func main() {

	intSlice := []int{}

	//turning the multiline string input to an int slice
	for _, inst := range strings.Split(input, "\n") {
		if inst != "" {
			val, err := strconv.Atoi(inst)
			if err != nil {
				fmt.Println("Atoi error!")
			}
			intSlice = append(intSlice, val)
		}
	}

	p1answer := 0
	pointer := 0

	for {
		p1answer++

		//move the pointer, while simultaneously incrementing the value where the pointer used to point to
		pointer, intSlice[pointer] = pointer+intSlice[pointer], intSlice[pointer]+1

		if pointer < 0 { //probably doesn't happen, but just in case
			pointer += len(intSlice)

		} else if pointer >= len(intSlice) { //this is the instruction escape condition
			pointer -= len(intSlice)
			fmt.Println("Part 1:", p1answer)
			return
		}
	}

}
