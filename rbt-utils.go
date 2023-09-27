package main

func (node *Node) isRootNode() bool {
	return node.parent == nil
}

func (node *Node) isLeftChild() bool {
	return node != nil && node.parent != nil && node.parent.left != nil && node.parent.left == node
}

func (x *Node) isRightChild() bool {
	return !x.isRootNode() && !x.isLeftChild()
}

func (x *Node) getLeftSibling() *Node {
	if x == nil || x.parent == nil {
		return nil
	}
	return x.parent.left
}

func (x *Node) getRightSibling() *Node {
	if x == nil || x.parent == nil {
		return nil
	}
	return x.parent.right
}

func (x *Node) getSibling() *Node {
	if x == nil || x.parent == nil {
		return nil
	}
	if x.isLeftChild() {
		return x.parent.right
	}
	return x.parent.left
}
