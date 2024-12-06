package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

// var demo = `rect 3x2
// rotate column x=1 by 1
// rotate row y=0 by 4
// rotate column x=1 by 1
// `

var screen [6][50]int

func main() {

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		line_parts := strings.Split(line, " ")
		// fmt.Println(line_parts)

		switch line_parts[0] {
		case "rect":
			//rect
			rect_parts := strings.Split(line_parts[1], "x")
			valX, err := strconv.Atoi(rect_parts[0])
			if err != nil {
				fmt.Println("yo there's an error here")
			}
			valY, _ := strconv.Atoi(rect_parts[1])
			if err != nil {
				fmt.Println("yo there's an error here")
			}
			for y := 0; y < valY; y++ {
				for x := 0; x < valX; x++ {
					screen[y][x] = 1
				}
			}

		case "rotate":
			if line_parts[1] == "column" {
				//column
				parts := strings.Split(line_parts[2], "=")
				valx, err := strconv.Atoi(parts[1])
				if err != nil {
					fmt.Println("yo there's an error here")
				}
				reps, err := strconv.Atoi(line_parts[4])
				if err != nil {
					fmt.Println("yo there's an error here")
				}
				slicecopy := []int{}
				for y := 0; y < len(screen); y++ {
					slicecopy = append(slicecopy, screen[y][valx])
				}
				for y := 0; y < len(screen); y++ {
					newnum := (y - reps)
					if newnum < 0 {
						newnum = len(screen) - max(newnum, -newnum)
					}
					screen[y][valx] = slicecopy[newnum]
				}

			} else {
				//row
				parts := strings.Split(line_parts[2], "=")
				valy, err := strconv.Atoi(parts[1])
				if err != nil {
					fmt.Println("yo there's an error here")
				}
				reps, err := strconv.Atoi(line_parts[4])
				if err != nil {
					fmt.Println("yo there's an error here")
				}
				slicecopy := []int{}
				for x := 0; x < len(screen[0]); x++ {
					slicecopy = append(slicecopy, screen[valy][x])
				}
				for x := 0; x < len(screen[0]); x++ {
					newnum := (x - reps)
					if newnum < 0 {
						newnum = len(screen[0]) - max(newnum, -newnum)
					}
					screen[valy][x] = slicecopy[newnum]
				}

			}
		}
	}
	counter := 0
	for y := 0; y < len(screen); y++ {
		for x := 0; x < len(screen[0]); x++ {
			if screen[y][x] == 1 {
				counter++
			}
		}
	}
	fmt.Println("Part 1: ", counter)

}

func printscreen() {
	for y := 0; y < len(screen); y++ {
		for x := 0; x < len(screen[0]); x++ {
			if screen[y][x] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Print("\n")
	}
}
