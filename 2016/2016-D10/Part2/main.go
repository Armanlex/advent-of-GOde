package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

type bot struct {
	id int

	valH int
	valL int

	giveHigh int
	botHigh  bool

	giveLow int
	botLow  bool
}

var bots = [1000]bot{}
var output = [1000]int{}

func main() {

	for i := 0; i < 1000; i++ {
		botInit(&bots[i], i)
		output[i] = -1
	}

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")

		if parts[0] == "bot" {
			//1,6,11
			who, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Atoi error here")
			}

			typeLow := parts[5]

			low, err := strconv.Atoi(parts[6])
			if err != nil {
				fmt.Println("Atoi error here")
			}

			bots[who].giveLow = low
			bots[who].botLow = (typeLow == "bot")

			typeHigh := parts[10]
			high, err := strconv.Atoi(parts[11])
			if err != nil {
				fmt.Println("Atoi error here")
			}

			bots[who].giveHigh = high
			bots[who].botHigh = (typeHigh == "bot")

		} else {
			toWho, err := strconv.Atoi(parts[5])
			if err != nil {
				fmt.Println("Atoi error here")
			}

			what, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Atoi error here")
			}

			giveBotANum(what, &bots[toWho])

		}

	}

	for {
		for _, aBot := range bots {
			// fmt.Println(bots[i])
			checkBot(&aBot)
			if output[0] >= 0 && output[1] >= 0 && output[2] >= 0 { //jesus so stupid... the exercise didn't specify I need to stop when all these are written over the first time, I thought the system had to settle
				fmt.Println("Part 2: ", output[0]*output[1]*output[2])
				return
			}

		}
	}

}

func botInit(robot *bot, id int) {
	robot.id = id
	robot.valH = -1
	robot.valL = -2
	robot.giveHigh = -1
	robot.botHigh = false
	robot.giveLow = -1
	robot.botLow = false
}

func giveBotANum(num int, robot *bot) {
	if num > robot.valH {
		num, robot.valH = robot.valH, num
		if num > robot.valL {
			num, robot.valL = robot.valL, num
		}
	} else {
		if num > robot.valL {
			num, robot.valL = robot.valL, num
		}
	}
}

func checkBot(robot *bot) {
	if robot.valH >= 0 && robot.valL >= 0 {
		if robot.botHigh {
			giveBotANum(robot.valH, &bots[robot.giveHigh])
		} else {
			output[robot.giveHigh] = robot.valH
		}

		if robot.botLow {
			giveBotANum(robot.valL, &bots[robot.giveLow])
		} else {
			output[robot.giveLow] = robot.valL
		}

		robot.valH = -1
		robot.valL = -2

		checkBot(&bots[robot.giveHigh])

		checkBot(&bots[robot.giveLow])

	}

}
