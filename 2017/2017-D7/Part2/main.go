package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

type node struct {
	Name     string
	Parent   *node
	Children []*node
	Data     int
}

var lookupMap map[string]*node //will be used to quickly get the reference of a node by just using its name

func main() {

	lookupMap = make(map[string]*node)

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		//all the variables needed to construct a new node
		var mainName string
		var mainVal int
		var childrenNames []string

		if strings.Contains(line, "->") {
			//the node had other nodes on it

			//I'll use this string as an example:
			//"tknk (41) -> ugml, padx, fwft"

			temp := strings.Split(line, " -> ")
			//temp: "tknk (41)", "ugml, padx, fwft"

			children := temp[1] //saving for later

			main := temp[0]
			temp = strings.Split(main, " (")
			//temp: "tknk", "41)"

			mainName = temp[0]
			mainValStr := temp[1][:len(temp[1])-1] //removing the ) at the end
			val, err := strconv.Atoi(mainValStr)
			if err != nil {
				fmt.Println("Atoi error!")
			}
			mainVal = val

			childrenNames = strings.Split(children, ", ")
			//childrenNames: "ugml", "padx", "fwft"

		} else {
			//node doesn't have other nodes on it

			//I'll use this string as an example:
			//"tknk (41)"

			temp := strings.Split(line, " (")
			//temp: "tknk", "41)"

			mainName = temp[0]
			mainValStr := temp[1][:len(temp[1])-1] //removing the ) at the end
			val, err := strconv.Atoi(mainValStr)
			if err != nil {
				fmt.Println("Atoi error!")
			}
			mainVal = val
		}

		n, ok := lookupMap[mainName]
		mainNode := &node{} //generate a temporary new node pointer

		if ok { //if main node found in map then update its value (cause it was created by a parent who doesn't have that information)
			mainNode = n
			mainNode.Data = mainVal

		} else { //create the new node, give it a value and add it to the map
			mainNode.Name = mainName
			mainNode.Data = mainVal
			lookupMap[mainName] = mainNode
		}

		for _, childName := range childrenNames { //iterate over the children if there are any
			childNode, ok := lookupMap[childName]
			if ok { //if child found in map, just update the parent (cause it created itself)
				childNode.Parent = mainNode
			} else { //if child not found, then create it, update it best we can, then add it to map
				NewChildNode := &node{}
				NewChildNode.Name = childName
				NewChildNode.Parent = mainNode
				lookupMap[childName] = NewChildNode
				childNode = NewChildNode
			}

			mainNode.Children = append(mainNode.Children, childNode)
		}

	}

	root := &node{}

	for _, node := range lookupMap { //just trying to access any node on the tree, doesn't matter which one as all of them should lead to the original parent
		for node.Parent != nil {
			node = node.Parent
		}
		root = node
		break
	}

	recur(root)

}

type NodeWeight struct {
	load   int
	weight int
	total  int
}

//takes a node, and recurses into its children,
//if no children, then return with weight
//after all chidren have returned, check if their weights are correct and print the solution if imbalance is found

func recur(curNode *node) (int, int) {
	myWeight := curNode.Data
	myLoad := 0

	fattestChild := NodeWeight{load: 0, weight: 0, total: 0}
	lightestChild := NodeWeight{load: 9999999, weight: 9999999, total: 9999999 + 9999999}

	for _, childNode := range curNode.Children {

		childWeight, childLoad := recur(childNode)

		if fattestChild.total < childWeight+childLoad {
			fattestChild.load = childLoad
			fattestChild.weight = childWeight
			fattestChild.total = childWeight + childLoad

		}

		if lightestChild.total > childWeight+childLoad {
			lightestChild.load = childLoad
			lightestChild.weight = childWeight
			lightestChild.total = childWeight + childLoad
		}

		myLoad += childLoad + childWeight
	}

	if len(curNode.Children) != 0 {
		if lightestChild.total != fattestChild.total {
			//mismatch found!
			fmt.Println("Part 2:", fattestChild.weight-(fattestChild.total-lightestChild.total))
			os.Exit(0)
		}
	}

	return myWeight, myLoad
}
