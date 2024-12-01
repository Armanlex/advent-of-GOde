package main

import (
	"fmt"
	"strings"
)

var input string = `YOUR_INPUT_HERE` //careful, if you copy paste your shape here, vs code might eat the spaces on the left of each line. I fixed it by hitting ctrl z after pasting, and this step is removed.

var grid [][]rune

func main() {

	for _, line := range strings.Split(input, "\n") { //iterate over input and place it into a grid
		if line == "" {
			continue
		}
		row := []rune(line)
		grid = append(grid, row)
	}

	currX := 0
	currY := 0
	dirX := 0
	dirY := 1

	for x := 0; x < len(grid[0]); x++ { //finding where the program will start at
		if grid[0][x] == '|' {
			currX = x
			break
		}
	}

	p1answer := ""

	for {
		switch grid[currY][currX] {
		case '+':
			//find where to turn to
			if dirY == 0 { //going right or left
				if isValid(currX, currY+1) && grid[currY+1][currX] != ' ' {
					//go down
					dirY = 1
					dirX = 0
				} else if isValid(currX, currY-1) && grid[currY-1][currX] != ' ' {
					//go up
					dirY = -1
					dirX = 0
				}

			} else if dirX == 0 { //going up or down
				if isValid(currX+1, currY) && grid[currY][currX+1] != ' ' {
					//go down
					dirY = 0
					dirX = 1
				} else if isValid(currX-1, currY) && grid[currY][currX-1] != ' ' {
					//go up
					dirY = 0
					dirX = -1
				}

			}

		default:
			if grid[currY][currX] >= 'A' && grid[currY][currX] <= 'Z' {
				p1answer += string(grid[currY][currX])
			}
		}
		currX += dirX
		currY += dirY

		if grid[currY][currX] == ' ' {
			break
		}

	}
	fmt.Println("Part 1:", p1answer)

}

func isValid(x, y int) bool {
	return x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid)
}
