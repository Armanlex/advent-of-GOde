package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

var input string = `YOUR_INPUT_HERE`

var cacheThrees = []byte{} //holds which letter is the first triplet, if any
var cacheFives = []byte{}  //holds which letter is the first pentuplet, if any
var hashSlice = []string{} //holds all the hashes after they've been generated

func main() {

	start := time.Now() //starting timer for benchmarking purposes

	var answer = "" //to keep track of how many password symbols have been generated

	var triplet byte = 255 //holds if the particular Nth hash contains a triplet, 255 means is has none

	Nth := 0
	for {
		//find triplet of currhash or generate it
		triplet = 255
		triplet = getTriplet(Nth)

		if triplet != 255 {
			//triplet found

			for i := Nth + 1; i < 1000+Nth+1; i++ {
				//checking the next 1k positions for a pentuplet
				if getPentuplet(i) == triplet {
					answer += string(rune(triplet))
					break
				}

			}
		}

		if len(answer) == 64 {
			fmt.Println("Part 2: ", Nth)
			fmt.Println("Calculation duration: ", time.Since(start))
			break
		}

		Nth++
	}

}

// returns if a particular Nth hash has a triplet,
// if the answer is missing from the cache then it's generated and then cached
func getTriplet(num int) byte {
	if len(cacheThrees) <= num {
		calcCombos(getHash(num))
	}
	return cacheThrees[num]
}

// returns if a particular Nth hash has a pentuplet,
// if the answer is missing from the cache then it's generated and then cached
func getPentuplet(num int) byte {
	if len(cacheThrees) <= num {
		calcCombos(getHash(num))
	}
	return cacheFives[num]
}

// calculates if a particular hash has a triplet or pentuplet, and then caches the result
func calcCombos(hash string) {
	combo3 := 0
	combo5 := 0
	var combosThree byte = 255
	var combosFive byte = 255

	for i := 0; i < 32-1; i++ {
		if combosThree == 255 {
			if hash[i] == hash[i+1] {
				combo3++
				if combo3 == 2 {
					combosThree = hash[i]
				}
			} else {
				combo3 = 0
			}
		}
		if combosFive == 255 {
			if hash[i] == hash[i+1] {
				combo5++
				if combo5 == 4 {
					combosFive = hash[i]
					break
				}
			} else {
				combo5 = 0
			}
		}
	}

	cacheThrees = append(cacheThrees, combosThree)
	cacheFives = append(cacheFives, combosFive)
}

// returns the hash string of an Nth hash, if it's not found in cache, it's calculated and then cached
func getHash(num int) string {
	if len(hashSlice) > num {
		return hashSlice[num]

	} else {
		number := strconv.Itoa(num)
		res := fmt.Sprintf("%x", md5.Sum([]byte(input+number)))
		hash_b := [16]byte{}
		hash_b_slice := hash_b[:]

		for i := 0; i < 2016; i++ {
			hash_b = md5.Sum([]byte(res))
			res = hex.EncodeToString(hash_b_slice)
		}

		hashSlice = append(hashSlice, res)
		return hashSlice[len(hashSlice)-1]
	}
}
