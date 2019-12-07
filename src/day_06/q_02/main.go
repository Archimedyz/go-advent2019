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

	root, you, san := GenerateOrbitMap(string(input))

	UpdateOrbitMap(root)

	ancestor := FindClosestAncestor(you, san)

	transfers := (*you).level + (*san).level - 2 - ((*ancestor).level * 2)

	fmt.Printf(">>> ancestor: %v (%d)\n>>> you:      %v (%d)\n>>> santa:    %v (%d)\n", (*ancestor).name, (*ancestor).level, (*you).name, (*you).level, (*san).name, (*san).level)

	fmt.Printf("Total transfers: %d\n\nTerminated.", transfers)
}

func FindClosestAncestor(a *Node, b *Node) *Node {
	var ancestor *Node

	seenNodes := make(map[*Node]bool)
	
	currA := a
	currB := b

	for currA != nil && currB != nil {
		if currA != nil {
			if seenNodes[currA] {
				ancestor = currA
				break
			}
			seenNodes[currA] = true
			currA = (*currA).parent
		}

		if currB != nil {
			if seenNodes[currB] {
				ancestor = currB
				break
			}
			seenNodes[currB] = true
			currB = (*currB).parent
		}
	}

	return ancestor
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

func GenerateOrbitMap(input string) (*Node, *Node, *Node) {
	nodeMap := make(map[string]*Node)
	var root *Node
	var you *Node
	var san *Node

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

		if orbitKey == "YOU" {
			you = orbitNode
		} else if orbitKey == "SAN" {
			san = orbitNode
		}
	}

	return root, you, san
}

func GetNodeKeys(input string) (string, string) {
	split := strings.Split(input, ")")
	return split[0], split[1]
}
