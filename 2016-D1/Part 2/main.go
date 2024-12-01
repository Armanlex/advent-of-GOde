package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

type pos struct {
	x int
	y int
}

func main() {
	mymap := make(map[pos]int)

	dirs := []pos{}
	dirs = append(dirs, pos{x: 0, y: -1}) //up
	dirs = append(dirs, pos{x: 1, y: 0})  //right
	dirs = append(dirs, pos{x: 0, y: +1}) //down
	dirs = append(dirs, pos{x: -1, y: 0}) //left

	me := pos{x: 0, y: 0}

	dir := 0 // up 0, right 1, down 2, left 3

	for _, chunk := range strings.Split(input, ", ") {
		if chunk == "" {
			continue
		}

		switch chunk[0] {
		case 'R':
			dir++

		case 'L':
			dir--
		}

		dir = dir % 4
		if dir < 0 {
			dir = 3
		}

		mult, err := strconv.Atoi(chunk[1:])
		if err != nil {
			fmt.Println("yo! there's an error here dawg")
		}
		for i := 0; i < mult; i++ {
			me.x += dirs[dir].x
			me.y += dirs[dir].y
			fmt.Println(me)
			_, ok := mymap[me]
			if ok {
				fmt.Println("Part 2: ", max(me.x, -me.x)+max(me.y, -me.y))
				return
			}
			mymap[me] = 1
		}

	}

}
