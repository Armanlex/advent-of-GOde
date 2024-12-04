package main

import (
	"fmt"
	"os"
	"strings"
)

type pos struct {
	x int
	y int
}

func main() {
	input := getInput()
	grid := textToGrid(input)
	fmt.Println("Part 1 answer:", part1(grid))
	fmt.Println("Part 2 answer:", part2(grid))
}

func part1(grid [][]rune) string {

	directions := make([]pos, 8)
	directions[0] = pos{x: 0, y: -1}  //up
	directions[1] = pos{x: 1, y: -1}  //up right
	directions[2] = pos{x: 1, y: 0}   //right
	directions[3] = pos{x: 1, y: 1}   //down right
	directions[4] = pos{x: 0, y: 1}   //down
	directions[5] = pos{x: -1, y: 1}  //down left
	directions[6] = pos{x: -1, y: 0}  //left
	directions[7] = pos{x: -1, y: -1} //left up

	XMASCounter := 0

	correctWord := [4]rune{'X', 'M', 'A', 'S'}
	wordBuffer := [4]rune{}
	reset := [4]rune{}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid); x++ {

			for _, dir := range directions { //pick a direction
				for dist := 0; dist < 4; dist++ { //travel that direction this many times
					newPos := pos{x: x + dir.x*dist, y: y + dir.y*dist}
					if !inBounds(newPos, grid) {
						break
					}
					wordBuffer[dist] = grid[newPos.y][newPos.x] //if the new position is in bounds, save it on the buffer
				}
				if wordBuffer == correctWord { //check the buffer
					XMASCounter++
				}
				wordBuffer = reset
			}
		}
	}

	return fmt.Sprint(XMASCounter)

}

func part2(grid [][]rune) string {
	directions := make([]pos, 4)

	directions[0] = pos{x: 1, y: -1}  //up right
	directions[1] = pos{x: 1, y: 1}   //down right
	directions[2] = pos{x: -1, y: 1}  //down left
	directions[3] = pos{x: -1, y: -1} //left up

	MASCounter := 0

	correctWord := [3]rune{'M', 'A', 'S'}
	wordBuffer := [3]rune{}
	reset := [3]rune{}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid); x++ {
			hitCounter := 0                  //it needs to hit 2 times to know that we hit an X-MAS
			for _, dir := range directions { //pick a direction
				for dist := -1; dist < 2; dist++ { //travel that direction this many times, but since it starts from -1, we take a step back first
					newPos := pos{x: x + dir.x*dist, y: y + dir.y*dist}
					if !inBounds(newPos, grid) {
						break
					}
					wordBuffer[dist+1] = grid[newPos.y][newPos.x] //if the new position is in bounds, save it on the buffer
				}
				if wordBuffer == correctWord { //check the buffer
					hitCounter++
				}
				wordBuffer = reset
			}
			if hitCounter == 2 {
				MASCounter++
			}
		}
	}

	return fmt.Sprint(MASCounter)
}

func textToGrid(str string) [][]rune {
	lines := strings.Split(str, "\n")

	grid := make([][]rune, len(lines))

	for li, line := range lines {
		for ci := 0; ci < len(line); ci++ {
			grid[li] = append(grid[li], rune(line[ci]))
		}
	}
	return grid
}

func inBounds(p pos, grid [][]rune) bool {
	return p.x >= 0 && p.x < len(grid[0]) && p.y >= 0 && p.y < len(grid)
}

func getInput() string {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	if len(data) == 0 {
		fmt.Println(fileName, " file is empty")
		os.Exit(1)
	}
	input := strings.ReplaceAll(string(data), "\r\n", "\n") //doing this replace so it can handle both linux and window text format
	return strings.TrimSpace(input)                         //doing this cause usually there's an extra new line at the bottom of the input
}
