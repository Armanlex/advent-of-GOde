package main

import (
	"fmt"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

var demo string = `0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5`

type node struct {
	Name        string
	ConnectedTo map[*node]int
}

var nodeMap map[string]*node //links the name of a node to the node pointer, for ease of access

var foundNodes map[*node]int //used to keep track of which nodes have been already visited

func main() {

	nodeMap = make(map[string]*node)
	foundNodes = make(map[*node]int)

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		slice := strings.Split(line, " <-> ")
		from := slice[0]

		n, ok := nodeMap[from] //testing to see if node exists in map

		fromNode := &node{}
		if ok == false { //if node doesn't exist, then update temp new node above, and add it to map
			fromNode.Name = from
			fromNode.ConnectedTo = make(map[*node]int)
			nodeMap[from] = fromNode
		} else { //if node does exist, replace temp new node with with the found node
			fromNode = n
		}

		for _, toNodeName := range strings.Split(slice[1], ", ") {
			if toNodeName == "" {
				continue
			}

			n, ok := nodeMap[toNodeName]
			toNode := &node{}
			if ok == false { //if node doesn't exist, then update temp new node above, and add it to map
				toNode.Name = toNodeName
				toNode.ConnectedTo = make(map[*node]int)
				nodeMap[toNodeName] = toNode

			} else { //if node does exist, replace temp new node with with the found node
				toNode = n
			}

			//creating the connectoin between the two nodes
			fromNode.ConnectedTo[toNode] = 1
			toNode.ConnectedTo[fromNode] = 1

		}
	}

	fmt.Println("Part 1:", countGroup(nodeMap["0"]))

}

// recursively takes the root node and enters all the children,
// but checks if it has entered the child already using foundNodes
func countGroup(root *node) int {

	foundNodes[root]++
	num := 0
	for n := range root.ConnectedTo {

		_, ok := foundNodes[n]
		if ok == false {
			num += countGroup(n)
		}
	}
	num++

	return num
}
