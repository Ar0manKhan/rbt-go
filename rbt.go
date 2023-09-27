package main

type Node struct {
	key                 int
	colorIsRed          bool
	left, right, parent *Node
}

type Tree struct {
	root *Node
	null *Node
}

func GenerateTree() *Tree {
	t := &Tree{null: &Node{key: 0, colorIsRed: false}}
	t.null.left = t.null
	t.null.right = t.null
	t.null.parent = t.null
	t.root = t.null
	return t
}
