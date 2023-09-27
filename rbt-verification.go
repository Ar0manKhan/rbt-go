package main

import "fmt"

func verify_rbt_properties(tree *Tree) bool {
	// The root is black.
	fmt.Println("Testing #1 - SKIPPING - Every node is either red or black.")

	fmt.Println("Testing #2 - The root is black.")
	if tree.root != nil && tree.root.colorIsRed {
		fmt.Println("[ERROR] Root node is not black")
		return false
	}

	fmt.Println("Testing #3 - SKIPPING - Every leaf (NIL) is black.")

	fmt.Println("Testing #4 - If a node is red, then both its children are black.")
	for n := range traverseIterative(tree.root) {
		if n.colorIsRed && ((n.left != nil && n.left.colorIsRed) || (n.right != nil && n.right.colorIsRed)) {
			fmt.Printf("[ERROR] %v has red children\n", n)
			return false
		}
	}

	// TODO
	fmt.Println("Testing #5 - SKIPPING - For each node, all simple paths from the node to descendant leaves contain the same number of black nodes.")
	if !validateCountBlackNodeChildren(tree.root) {
		fmt.Println("[ERROR] Validation failed")
		return false
	}
	fmt.Println("[PASSED]")
	return true
}

func validateCountBlackNodeChildren(n *Node) bool {
	countMap := map[*Node]int{}
	var countLeft, countRight int
	for node := range traverseIterativePostOrder(n) {
		if node.left == nil {
			countLeft = 1
		} else {
			countLeft = countMap[node.left]
		}
		if node.right == nil {
			countRight = 1
		} else {
			countRight = countMap[node.right]
		}
		if countLeft != countRight {
			fmt.Println("Validation failed on", node.key)
			return false
		} else if node.colorIsRed {
			countMap[node] = countLeft
		} else {
			countMap[node] = countLeft + 1
		}
	}
	return true
}
