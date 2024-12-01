package main

import (
	"fmt"
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

			if ok { //if child found in map, just update the parent pointer (cause it created itself and didn't have that info)
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

	for _, node := range lookupMap { //just trying to access any node on the tree, doesn't matter which one; as all of them should lead to the original parent
		for node.Parent != nil { //if node has parent, it's not the root
			node = node.Parent
		}
		fmt.Println(node.Name)
		return
	}

}
