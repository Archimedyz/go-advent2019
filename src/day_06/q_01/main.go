package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Node struct {
	parent   *Node
	children []*Node
	level    int
	name     string
}

func main() {
	input, _ := ioutil.ReadFile("../input.txt")
	// fmt.Printf("Contents:\n%v\n", string(input))

	root := GenerateOrbitMap(string(input))

	UpdateOrbitMap(root)

	result := GetLevelSum(root)

	fmt.Printf("Total orbits: %d\n\nTerminated.", result)
}

func GetLevelSum(n *Node) int {
	sum := (*n).level

	for _, ptrChild := range (*n).children {
		sum += GetLevelSum(ptrChild)
	}

	return sum
}

func UpdateOrbitMap(root *Node) {
	UpdateNode(root, 0)
}

func UpdateNode(n *Node, lvl int) {
	(*n).level = lvl

	for _, ptrChild := range n.children {
		UpdateNode(ptrChild, lvl+1)
	}
}

func GenerateOrbitMap(input string) *Node {
	nodeMap := make(map[string]*Node)
	var root *Node

	for _, v := range strings.Split(input, "\n") {
		centerKey, orbitKey := GetNodeKeys(v)

		var centerNode *Node
		var orbitNode *Node
		var exists bool
		// check to see if either node is already in the map
		// if not, add it
		if centerNode, exists = nodeMap[centerKey]; !exists {
			centerNode = new(Node)
			(*centerNode).name = centerKey
			nodeMap[centerKey] = centerNode
		}

		if orbitNode, exists = nodeMap[orbitKey]; !exists {
			orbitNode = new(Node)
			(*orbitNode).name = orbitKey
			nodeMap[orbitKey] = orbitNode
		}

		// now just do the appropriate assignments
		(*orbitNode).parent = centerNode
		(*centerNode).children = append((*centerNode).children, orbitNode)

		// finally, if we found our root, COM, assign it
		if centerKey == "COM" {
			root = centerNode
		}
	}

	return root
}

func GetNodeKeys(input string) (string, string) {
	split := strings.Split(input, ")")
	return split[0], split[1]
}
