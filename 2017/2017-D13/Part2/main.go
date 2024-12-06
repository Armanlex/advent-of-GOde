package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Slow solution, takes several seconds to get my answer on my strong desktop cpu
//definitely can be optimized to run instantly, by using the cycles and offsets of the scanners to calculate the answer in a mathematical way.
//or my brute force method can be improved by an addition where multiples of known failed delays values, are skipped and not calculated or simulated

var input string = `YOUR_INPUT_HERE`

var layerMap map[int]*layerT

type layerT struct {
	ScannerPos int
	Length     int
	Dir        int //1 means down, -1 means up
}

func main() {
	firewallLen := 0
	layerMap = make(map[int]*layerT)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		slice := strings.Split(line, ": ")

		layer, ok := strconv.Atoi(slice[0])
		if ok != nil {
			fmt.Println("Atoi error!")
		}

		len, ok := strconv.Atoi(slice[1])
		if ok != nil {
			fmt.Println("Atoi error!")
		}

		firewallLen = max(firewallLen, layer)
		l := &layerT{}
		l.Length = len
		l.Dir = 1

		layerMap[layer] = l

	}

	delay := 0

	for i := 0; i <= firewallLen; i++ {
		_, ok := layerMap[i]
		if ok {
			for x := 0; x < i; x++ {
				if layerMap[i].ScannerPos+layerMap[i].Dir >= layerMap[i].Length || layerMap[i].ScannerPos+layerMap[i].Dir < 0 {
					layerMap[i].Dir *= -1
				}
				layerMap[i].ScannerPos += layerMap[i].Dir
			}
		}
	}

	for {
		caught := false

		for _, node := range layerMap {
			if node.ScannerPos == 0 {
				caught = true
				break
			}
		}

		if caught == false {
			fmt.Println("Part 2:", delay)
			return
		}

		step()

		delay += 1

	}

}

func step() {

	for key := range layerMap {

		if layerMap[key].ScannerPos+layerMap[key].Dir >= layerMap[key].Length || layerMap[key].ScannerPos+layerMap[key].Dir < 0 {
			layerMap[key].Dir *= -1
		}
		layerMap[key].ScannerPos += layerMap[key].Dir

	}
}
