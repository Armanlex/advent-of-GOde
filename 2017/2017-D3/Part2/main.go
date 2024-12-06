package main

import (
	"fmt"
	"os"
	"strconv"
)

type pos struct {
	x int
	y int
}

var input string = `YOUR_INPUT_HERE`
var myInt int
var grid [99][99]int
var offset int = (99 / 2) + 1
var currPos pos

func main() {
	val, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Atoi Error!")
	}
	myInt = val
	currPos.x = 0
	currPos.y = 0

	grid[offset][offset] = 1
	for ring := 1; ring < 100; ring++ {
		endingVal := (ring + ring + 1) * (ring + ring + 1)
		endingValofPrevRing := (ring + ring - 1) * (ring + ring - 1)
		sideLen := (endingVal - endingValofPrevRing) / 4

		currPos.x += 1
		for i := 0; i < sideLen-1; i++ {

			addVal(currPos.x, currPos.y)

			currPos.y -= 1
		}
		for i := 0; i < sideLen; i++ {
			addVal(currPos.x, currPos.y)

			currPos.x -= 1
		}
		for i := 0; i < sideLen; i++ {
			addVal(currPos.x, currPos.y)

			currPos.y += 1
		}
		for i := 0; i < sideLen+1; i++ {
			addVal(currPos.x, currPos.y)

			if i < sideLen {
				currPos.x += 1
			}
		}

		// for i := 0; i < len(grid); i++ {
		// 	fmt.Println(grid[i])
		// }

	}

}

func addVal(curx, cury int) {
	curx += offset
	cury += offset
	grid[cury][curx] += grid[cury-1][curx]
	grid[cury][curx] += grid[cury+1][curx]
	grid[cury][curx] += grid[cury][curx-1]
	grid[cury][curx] += grid[cury][curx+1]
	grid[cury][curx] += grid[cury-1][curx-1]
	grid[cury][curx] += grid[cury+1][curx+1]
	grid[cury][curx] += grid[cury-1][curx+1]
	grid[cury][curx] += grid[cury+1][curx-1]
	if grid[cury][curx] > myInt {
		fmt.Println("Part 2: ", grid[cury][curx])

		os.Exit(0)
	}

}
