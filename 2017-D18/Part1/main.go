package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

var arr ['z' - 'a']int

func main() {
	lastSoundPlayed := 0

	commands := strings.Split(input, "\n")

	for index := 0; index < len(commands); index++ {

		line := commands[index]

		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		num := 0

		switch parts[0] {

		case "add":
			num = regToIndex(parts[1])
			arr[num] += fetchNum(parts[2])

		case "mul":
			num = regToIndex(parts[1])
			arr[num] *= fetchNum(parts[2])

		case "mod":
			num = regToIndex(parts[1])
			arr[num] %= fetchNum(parts[2])

		case "snd":
			lastSoundPlayed = fetchNum(parts[1])

		case "set":
			num = regToIndex(parts[1])
			arr[num] = fetchNum(parts[2])

		case "rcv":
			num = regToIndex(parts[1])

			if num == 0 {
				fmt.Println("Part 1:", lastSoundPlayed)
				return
			}

		case "jgz":
			num = fetchNum(parts[1])

			if num > 0 {
				index += fetchNum(parts[2]) - 1
			}
		}

	}

}

func MyAtoi(str string) int { //to avoid copy pasting the error handling....
	val, ok := strconv.Atoi(str)
	if ok != nil {
		fmt.Print("Atoi error!: %", str, "%\n")
	}
	return val
}

func isReg(str string) bool { //true if the string is a register
	return str[0] >= 'a' && str[0] <= 'z'
}

func regToIndex(str string) int { //turns the register character to the index of the array
	return int(str[0] - 'a')
}

func fetchNum(str string) int { //if it's given a register, it goes and find the number in the register, if it's not then it just returns the value it was given
	if isReg(str) {
		return arr[regToIndex(str)]
	} else {
		return MyAtoi(str)
	}
}
