package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

func main() {

	registers := [4]int{}
	instructions := strings.Split(input, "\n")
	for i := 0; i < len(instructions); i++ {
		line := instructions[i]
		if line == "" {
			continue
		}

		// fmt.Println(line, registers)

		parts := strings.Split(line, " ")
		if parts[len(parts)-1] == "" {
			parts = parts[:len(parts)-1]
		}

		switch parts[0] {
		case "cpy":
			if parts[1][0] >= 'a' && parts[1][0] <= 'd' {
				//reg to reg
				registers[int(parts[2][0]-'a')] = registers[int(parts[1][0]-'a')]
			} else {
				val, err := strconv.Atoi(parts[1])
				if err != nil {
					fmt.Println("yo atoi error here 1")
				}
				registers[int(parts[2][0]-'a')] = val

			}
		case "inc":
			registers[int(parts[1][0]-'a')]++

		case "dec":
			registers[int(parts[1][0]-'a')]--

		case "jnz":
			val, err := strconv.Atoi(parts[2])

			if err != nil {
				fmt.Println("yo atoi error here 1")
			}

			if parts[1][0] >= 'a' && parts[1][0] <= 'd' {
				if registers[int(parts[1][0]-'a')] != 0 {
					i += (val - 1)
				}
			} else {
				val2, err := strconv.Atoi(parts[1])
				if err != nil {
					fmt.Println("yo atoi error here 1")
				}
				if val2 != 0 {
					i += (val - 1)
				}

			}

		}

	}

	fmt.Println("Part 1: ", registers[0])
}
