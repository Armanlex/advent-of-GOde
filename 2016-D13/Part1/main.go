package main

//This is a super over engineered solution, I just wanted to do astar cause I've done all kinds of other
//	path finding algos in the past and I'm bored out of my mind with them

import (
	"fmt"
	"sort"
)

var input int = 12312412412412412 // << YOUR_INPUT_HERE

var grid = [50][50]node{} //size double of: 31,39

type node struct {
	parent  *node
	manDist int
	score   int
	length  int
	x       int
	y       int
	isWall  bool
	symb    rune
}

type pos struct {
	x int
	y int
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
			grid[y][x].manDist = abs(39-y) + abs(31-x)
			grid[y][x].symb = '.'
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

	frontSlice := []*node{}
	closedMap := make(map[*node]bool)
	frontMap := make(map[*node]bool)

	grid[1][1].score = grid[1][1].manDist + 0
	frontSlice = append(frontSlice, &grid[1][1])
	end := &grid[39][31]

	dirs := []pos{}
	dirs = append(dirs, pos{x: 0, y: 1})
	dirs = append(dirs, pos{x: 1, y: 0})
	dirs = append(dirs, pos{x: 0, y: -1})
	dirs = append(dirs, pos{x: -1, y: 0})

	var currNode *node
	finished := false
	for len(frontSlice) > 0 && !finished {
		// fmt.Println("loop")

		// for i := 0; i < len(frontSlice); i++ {
		// 	fmt.Print(frontSlice[i].x, ".", frontSlice[i].y, ".", frontSlice[i].score, ", ")
		// }

		// fmt.Println("")
		// printGrid()
		currNode = frontSlice[0]
		frontSlice = frontSlice[1:] //remove the first element
		delete(frontMap, currNode)
		closedMap[currNode] = true
		for _, dir := range dirs { //iterate over each cardinal direction
			if isValidPos(currNode.x+dir.x, currNode.y+dir.y) { //check if new position is valid to avoid out of bounds slice access
				newNode := &grid[currNode.y+dir.y][currNode.x+dir.x]
				if newNode.isWall == false {
					if newNode == end {
						end.parent = currNode
						finished = true
						break
					}
					_, isInClosed := closedMap[newNode]
					_, isInFront := frontMap[newNode]

					if !isInClosed && !isInFront {
						newNode.parent = currNode
						newNode.length = currNode.length + 1
						newNode.score = newNode.length + newNode.manDist
						frontSlice = InsertSorted(frontSlice, newNode)
						frontMap[newNode] = true
						newNode.symb = '?'

					} else if isInFront {
						if newNode.score > (currNode.length+1)+newNode.manDist {
							newNode.score = (currNode.length + 1) + newNode.manDist
							newNode.parent = currNode
							idx := sort.Search(len(frontSlice), func(i int) bool {
								return frontSlice[i] == newNode
							})
							if idx == len(frontSlice) {
								frontSlice = frontSlice[:len(frontSlice)-1]
							} else {
								fmt.Println(frontSlice, len(frontSlice), idx)
								frontSlice = append(frontSlice[:idx], frontSlice[idx+1:]...)
							}
							frontSlice = InsertSorted(frontSlice, newNode)
						}
					}
				}
			}

		}

	}

	counter := 0
	newcurrnode := end
	for newcurrnode.parent != nil {
		newcurrnode = newcurrnode.parent
		counter++
	}
	fmt.Println("Part 1: ", counter)
}

func InsertSorted(slice []*node, value *node) []*node { //from chatgpt then modified
	idx := sort.Search(len(slice), func(i int) bool {
		return slice[i].score >= value.score
	})
	slice = append(slice, &node{})   // Expand slice
	copy(slice[idx+1:], slice[idx:]) // Shift elements
	slice[idx] = value               // Insert value
	return slice
}

func isValidPos(x int, y int) bool {
	return x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func printGrid() {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x].isWall {
				fmt.Print("#")
			} else {
				fmt.Print(string(grid[y][x].symb))
			}

		}
		fmt.Println("")
	}
}
