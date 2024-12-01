package main

import (
	"fmt"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

type pos struct {
	x int
	y int
}

func main() {

	me := pos{x: 2, y: 2}
	word := ""
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		for _, r := range line {
			switch r {
			case 'U':
				me.y--
			case 'R':
				me.x++
			case 'D':
				me.y++
			case 'L':
				me.x--
			}
			me.x, me.y = max(me.x, 1), max(me.y, 1)
			me.x, me.y = min(me.x, 3), min(me.y, 3)
		}
		word += string((me.y-1)*3 + me.x + '0')
	}

	fmt.Println("Part 1: ", word)

}
