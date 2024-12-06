package main

import (
	"fmt"
)

var Ainput int = 0 //YOUR INPUT HERE
var Binput int = 0 //YOUR INPUT HERE

var genAfac int = 16807
var genBfac int = 48271
var divider int = 2147483647

func main() {

	p1answer := 0
	for i := 0; i < 5000000; i++ {

		Ainput = (Ainput * genAfac) % divider

		for Ainput%4 != 0 {
			Ainput = (Ainput * genAfac) % divider
		}

		Binput = (Binput * genBfac) % divider

		for Binput%8 != 0 {
			Binput = (Binput * genBfac) % divider
		}

		if Ainput%65536 == Binput%65536 { //this is the same as comparing the last 16 bits of the binary representation
			p1answer++
		}
	}
	fmt.Println("Part 2:", p1answer)

}
