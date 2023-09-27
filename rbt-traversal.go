package main

// Traverse in ascending order - i.e left first, then parent and then right
func traverseIterativeInAscOrder(root *Node) <-chan *Node {
	cl := make(chan *Node)
	go func() {
		traverseInAscOrder(root, cl)
		close(cl)
	}()
	return cl
}
func traverseInAscOrder(n *Node, cl chan *Node) {
	if n == nil {
		return
	}
	traverseInAscOrder(n.left, cl)
	cl <- n
	traverseInAscOrder(n.right, cl)
}

// Post Order Traversal left->right->parent
func traverseIterativePostOrder(tree *Tree, root *Node) <-chan *Node {
	cl := make(chan *Node)
	go func() {
		stk := []*Node{root}
		visited := map[*Node]bool{}
		var tempNode *Node
		for len(stk) > 0 {
			tempNode = stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			if tempNode == tree.null {
				continue
			}
			if visited[tempNode] {
				cl <- tempNode
				continue
			}
			visited[tempNode] = true
			stk = append(stk, tempNode, tempNode.right, tempNode.left)
		}
		close(cl)
	}()
	return cl
}

func traverseIterative(tree *Tree) <-chan *Node {
	cl := make(chan *Node)
	go func() {
		stk := []*Node{tree.root}
		var tempNode *Node
		for len(stk) > 0 {
			tempNode = stk[0]
			stk = stk[1:]
			if tempNode != tree.null {
				cl <- tempNode
				stk = append(stk, tempNode.left, tempNode.right)
			}
		}
		close(cl)
	}()
	return cl
}
