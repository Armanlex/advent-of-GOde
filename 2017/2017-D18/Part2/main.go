package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

var p2answer int

var arrA ['z' - 'a']int
var arrB ['z' - 'a']int

var queueA []int
var queueB []int

var ProgramAWaiting bool = false
var ProgramAOutputSlice []int

var ProgramBWaiting bool = false
var ProgramBOutputSlice []int

//basically each program operates once, if it sends somethng, it's added to it's queue, if it received an input, it reads the other program's
//queue, if it finds nothing, it returns and tries again. eventually both programs get stuck, then we print the result.

func main() {

	commands := strings.Split(input, "\n")
	arrA[regToIndex("p")] = 0 //this was already the case by default
	arrB[regToIndex("p")] = 1

	indexA := 0
	indexB := 0
	for {

		if indexA >= 0 && indexA < len(commands) {
			execute(0, &indexA, &arrA, &queueA, &queueB, commands[indexA])
		}

		if indexB >= 0 && indexB < len(commands) {
			execute(1, &indexB, &arrB, &queueB, &queueA, commands[indexB])
		}

	}

}

//takes a ref to AN array,
//a ref to a queue to add outputs into
//the command to execute
//then executes it on the array or w/e it needs to do

func execute(ID int, index *int, array *['z' - 'a']int, myqueue *[]int, othersQueue *[]int, line string) {
	if line == "" {
		*index++
		return
	}

	parts := strings.Split(line, " ")

	num := 0

	// fmt.Print("\n", line, "\n")

	switch parts[0] {

	case "add":
		num = regToIndex(parts[1])
		(*array)[num] += fetchNum(array, parts[2])
		*index++
		return

	case "mul":
		num = regToIndex(parts[1])
		(*array)[num] *= fetchNum(array, parts[2])
		*index++
		return

	case "mod":
		num = regToIndex(parts[1])
		(*array)[num] %= fetchNum(array, parts[2])
		*index++
		return

	case "snd":
		if ID == 1 {
			p2answer++
		}
		*myqueue = append(*myqueue, fetchNum(array, parts[1]))
		*index++
		return

	case "set":
		num = regToIndex(parts[1])
		(*array)[num] = fetchNum(array, parts[2])
		*index++
		return

	case "rcv":
		if len(*othersQueue) > 0 {
			num = regToIndex(parts[1])
			(*array)[num] = (*othersQueue)[0]
			*othersQueue = (*othersQueue)[1:]
			*index++

			if ID == 0 {
				ProgramAWaiting = false
			}

			if ID == 1 {
				ProgramBWaiting = false
			}

			return
		} else {
			//Do not increment, so that we "wait"
			if ID == 0 {
				ProgramAWaiting = true
			}

			if ID == 1 {
				ProgramBWaiting = true
				if ProgramAWaiting {
					fmt.Println("Part 2:", p2answer)
					os.Exit(0)
				}
			}

			return
		}

	case "jgz":
		num = fetchNum(array, parts[1])
		if num > 0 {
			*index += fetchNum(array, parts[2]) - 1
		}
		*index++
		return
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

func fetchNum(array *['z' - 'a']int, str string) int { //if it's given a register, it goes and find the number in the register, if it's not then it just returns the value it was given
	if isReg(str) {
		return (*array)[regToIndex(str)]
	} else {
		return MyAtoi(str)
	}
}
