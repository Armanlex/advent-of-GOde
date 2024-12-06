package main

import "fmt"

var input int = 312930612983192 // << YOUR_INPUT_HERE

var grid = [70][70]pos{}

type pos struct {
	x      int
	y      int
	isWall bool
}

func main() {

	num := 0
	count := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			num = x*x + 3*x + 2*x*y + y + y*y + input
			count = 0
			grid[y][x].x = x
			grid[y][x].y = y
			for num > 0 {
				if num%2 == 1 {
					count++
				}
				num /= 2
			}
			if count%2 == 1 {
				grid[y][x].isWall = true

			}
		}
	}

	dirs := []pos{}
	dirs = append(dirs, pos{x: 0, y: 1})
	dirs = append(dirs, pos{x: 1, y: 0})
	dirs = append(dirs, pos{x: 0, y: -1})
	dirs = append(dirs, pos{x: -1, y: 0})

	frontSlice := []*pos{}
	tempSlice := []*pos{}
	frontSlice = append(frontSlice, &grid[1][1])

	grid[1][1].isWall = true
	counter := 1
	dist := 1
	for len(frontSlice) > 0 {
		for i := 0; i < len(frontSlice); i++ {
			currNode := frontSlice[i]
			for _, dir := range dirs { //iterate over each cardinal direction
				if isValidPos(currNode.x+dir.x, currNode.y+dir.y) { //check if new position is valid to avoid out of bounds slice access
					newNode := &grid[currNode.y+dir.y][currNode.x+dir.x]
					if !newNode.isWall {
						newNode.isWall = true
						counter++
						tempSlice = append(tempSlice, newNode)
					}
				}
			}
		}
		dist++
		if dist == 51 {
			break
		}
		frontSlice = tempSlice
	}

	fmt.Println("Part 2: ", counter)

}

func isValidPos(x int, y int) bool {
	return x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid)
}
