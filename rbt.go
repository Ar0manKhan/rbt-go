package main

type Node struct {
	key                 int
	colorIsRed          bool
	left, right, parent *Node
}

type Tree struct {
	root *Node
}
