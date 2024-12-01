package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

// var demo string = `b inc 5 if a > 1
// a inc 1 if b < 5
// c dec -10 if a >= 1
// c inc -20 if c == 10`

var memory map[string]int

func main() {
	// input = demo
	p2answer := -99999999
	memory = make(map[string]int)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		ret := strings.Split(line, " ")

		//["c", "inc", "-20", "if", "c", "==", "10"]
		//	0     1      2      3    4    5     6

		targetReg := ret[0] //Register that we change

		oper := ret[1] //operator
		sign := 1      //will be multiplied with our value to easily change if the value is added or removed, withouth using logic flows like ifs or switches
		if oper == "dec" {
			sign = -1
		}

		valMain, err := strconv.Atoi(ret[2]) //value that we will add to target register
		if err != nil {
			fmt.Println("Atoi error!1")
		}

		checkReg := ret[4] //Register that we check

		comp := ret[5] //comparator

		checkVal, err := strconv.Atoi(ret[6]) //value that we check
		if err != nil {
			fmt.Println("Atoi error!2")
		}

		pass := false

		switch comp {
		case "<":
			pass = memory[checkReg] < checkVal
		case ">":
			pass = memory[checkReg] > checkVal
		case "<=":
			pass = memory[checkReg] <= checkVal
		case ">=":
			pass = memory[checkReg] >= checkVal
		case "==":
			pass = memory[checkReg] == checkVal
		case "!=":
			pass = memory[checkReg] != checkVal
		}

		if pass {
			memory[targetReg] += (valMain * sign)
			if p2answer < memory[targetReg] {
				p2answer = memory[targetReg]
			}
		}

	}

	fmt.Println("Part 2:", p2answer)

}
