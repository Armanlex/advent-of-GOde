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

var grid = []string{}

func main() {

	grid = append(grid, "XX1XX")
	grid = append(grid, "X234X")
	grid = append(grid, "56789")
	grid = append(grid, "XABCX")
	grid = append(grid, "XXDXX")

	me := pos{x: 1, y: 3}
	word := ""
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		for _, r := range line {
			pre := me
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
			me.x, me.y = min(me.x, 5), min(me.y, 5)
			if grid[me.y-1][me.x-1] == 'X' {
				me = pre
			}
		}

		word += string(grid[me.y-1][me.x-1])

	}

	fmt.Println("Part 2: ", word)

}
