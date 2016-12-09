package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// This program reads a JSON file representing
// an acyclic directed graph and performs a depth-first
// and breadth-first traversal.

type Node struct {
	Data  int     `json:"data"`
	Nodes []*Node `json:"nodes"`
}

type WrapNode struct {
	Node     *Node
	Distance int
}

func (n *Node) addNext(node *Node) {
	n.Nodes = append(n.Nodes, node)
}

func main() {
	// read the file to a buffer
	dat, err := ioutil.ReadFile("info.json")
	if err != nil {
		panic(err)
	}

	// parse the json in the buffer and return a hash of json objects
	nodesHash := ingestJson(dat)

	// breadth first traversal
	// print the Data and Distance values of each node
	fmt.Println("Breadth First")
	queue := []*WrapNode{}

	// pull out the root node of the graph which is always at [1]
	// set its distance to 0
	startNode := WrapNode{nodesHash[1], 0}

	// add rootnode to queue
	queue = append(queue, &startNode)

	for len(queue) != 0 {
		// pop node from front of queue
		curr := queue[0]

		// move start of queue up one index
		queue = queue[1:]

		// print the current node's Data and Distance
		fmt.Printf("%v %v", curr.Node.Data, curr.Distance)
		fmt.Println()

		// put the current node's children onto the queue
		children := curr.Node.Nodes
		for _, node := range children {
			queue = append(queue, &WrapNode{node, curr.Distance + 1})
		}
	}

	// depth first transversal
	// print the Data and Distance values of each node
	fmt.Println("Depth First")
	stack := []*WrapNode{}

	// add root node to stack
	stack = append(stack, &startNode)

	for len(stack) != 0 {
		// pop node from back of stack
		curr := stack[len(stack)-1]

		// move end of stack back one index
		stack = stack[:len(stack)-1]

		// print the current node's Data and Distance
		fmt.Printf("%v %v", curr.Node.Data, curr.Distance)
		fmt.Println()

		// put the current node's children onto the stack
		children := curr.Node.Nodes
		for _, node := range children {
			stack = append(stack, &WrapNode{node, curr.Distance + 1})
		}
	}
}

func ingestJson(dat []byte) map[int]*Node {
	var nodesRaw []map[string]interface{}
	nodesHash := map[int]*Node{}
	if err := json.Unmarshal(dat, &nodesRaw); err != nil {
		panic(err)
	}
	// putting all nodes in a hash
	for _, node := range nodesRaw {
		rawData := node["data"].(float64)
		data := int(rawData)
		nodesHash[data] = &Node{data, []*Node{}}
	}
	// then linking up nodes
	for _, node := range nodesRaw {
		rawData := node["data"].(float64)
		data := int(rawData)
		theNode := nodesHash[data]
		rawNextArray := node["nodes"]
		nextArray := rawNextArray.([]interface{})
		for _, rawNodeNum := range nextArray {
			nodeNum := int(rawNodeNum.(float64))
			childNode := nodesHash[nodeNum]
			theNode.addNext(childNode)
		}
	}
	return nodesHash
}
