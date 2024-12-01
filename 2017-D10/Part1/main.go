package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

// var input string = `3, 4, 1, 5`

var arr [256]int

func main() {
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	pointer := 0
	skipSize := 0

	delimiter := ","

	if strings.Contains(input, ", ") { //there's a difference betwen the example and actual user input, so this should cover both
		delimiter = ", "
	}

	for _, numStr := range strings.Split(input, delimiter) {
		intVal, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Atoi error!")
		}
		reverse(&arr, pointer, intVal+pointer-1)
		pointer += intVal + skipSize
		skipSize++
	}
	fmt.Println(arr[0] * arr[1]) // Part 1 answer
}

func reverse(arr *[256]int, start int, end int) {
	a := 0
	b := 0
	for i := 0; i <= (end-start)/2; i++ {
		a = (start + i) % len(arr)
		b = (end - i) % len(arr)

		(*arr)[a], (*arr)[b] = (*arr)[b], (*arr)[a]
	}

}
