package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// this is a directed graph breadth-first search example

type Node struct {
	Data  int     `json:"data"`
	Nodes []*Node `json:"nodes"`
}

func (n *Node) addNext(node *Node) {
	n.Nodes = append(n.Nodes, node)
}

func main() {
	dat, err := ioutil.ReadFile("info.json")
	if err != nil {
		panic(err)
	}
	nodesHash := ingestJson(dat)
	fmt.Println(nodesHash[4].Nodes)

	// breadth first transversal first just print their names I guess
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
