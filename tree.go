package main

import (
	"fmt"
)

type Node struct {
	value int
	children []*Node
}

func newNode(value int) *Node {
	node := Node{value: value}
	return &node
}

var root Node

func addNode(value int, node Node){
	newNode := newNode(value)
	node.children = append(node.children, newNode)
}

func printTree(){
	fmt.Println("Tree!")
	fmt.Printf("Root node kids: %d", len(root.children))
}