package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

// var input string = `s1,x3/4,pe/b`

//x0/5
//s2
//pg/k

var arr [16]rune
var arr2 [16]rune  //used to temporarily hold a copy of the above
var arrmap [16]int //used as a map to quickly find on what index all the numbers are in, for quick char based swaps

func main() {

	for i := 'a'; i <= 'p'; i++ { //intializing the array with the rune characters, the array is in the initial dance state
		arr[i-'a'] = i
		arrmap[i-'a'] = int(i - 'a')
	}

	if strings.Contains(input, "\n") { //remove if there's a new line char in the input
		input = strings.Split(input, "\n")[0]
	}

	for _, command := range strings.Split(input, ",") {
		if command == "" {
			continue
		}

		switch command[0] { //switch using the first letter of the command, it can only be x, s or p
		case 'x':
			res := strings.Split(command, "/")
			//generate the 2 indeces
			Aint, err := strconv.Atoi(res[0][1:]) //index 1

			if err != nil {
				fmt.Println("atoi error!!!")
			}

			Bint, err := strconv.Atoi(res[1]) //index 2

			if err != nil {
				fmt.Println("atoi error!!!")
			}

			aletter := arr[Aint] - 'a' //offset so that for example a becomes 0, so that the map can be updated
			bletter := arr[Bint] - 'a'
			arr[Aint], arr[Bint] = arr[Bint], arr[Aint]                         //swap them
			arrmap[aletter], arrmap[bletter] = arrmap[bletter], arrmap[aletter] //update the map

		case 'p':
			res := strings.Split(command, "/")

			Aint := res[0][1:][0] - 'a' //offset them so that the map can be used
			Bint := res[1][0] - 'a'

			//using the map, we can immediately find what index those letters are standing on
			//or else we would need to iterate over the entire array until we find it
			arr[arrmap[Aint]], arr[arrmap[Bint]] = arr[arrmap[Bint]], arr[arrmap[Aint]]
			arrmap[Aint], arrmap[Bint] = arrmap[Bint], arrmap[Aint] //update the map

		case 's':

			theInt, err := strconv.Atoi(command[1:])

			if err != nil {
				fmt.Println("atoi error!!! 3")
			}

			arr2 = arr                //keeping a copy of the original position, so that when the og is being updated, there's no lost data
			for i := 0; i < 16; i++ { //iterating over
				x := i - theInt
				if x < 0 {
					arr[i] = arr2[16+x]
				} else {
					arr[i] = arr2[x]
				}

				arrmap[arr[i]-'a'] = i //updating the map with the new location
			}

		}

	}
	fmt.Print("Part 1: ")
	for i := 0; i < len(arr); i++ {
		fmt.Print(string(arr[i]))
	}
	fmt.Print("\n")

}
