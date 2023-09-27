package main

func (tree *Tree) leftRotate(x *Node) {
	// save y
	y := x.right
	// set y-left to x-right
	x.right = y.left
	if y.left != tree.null {
		y.left.parent = x
	}
	// link x's parent to y
	y.parent = x.parent
	// handle root case, left child or right child case
	if x.parent == tree.null {
		tree.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	// put x on y's left
	y.left = x
	x.parent = y
}

func (tree *Tree) rightRotate(y *Node) {
	x := y.left
	y.left = x.right
	if x.right != tree.null {
		x.right.parent = y
	}
	x.parent = y.parent
	if y.parent == tree.null {
		tree.root = x
	} else if y == y.parent.left {
		y.parent.left = x
	} else {
		y.parent.right = x
	}
	x.right = y
	y.parent = x
}

func (tree *Tree) Insert(key int) {
	z := &Node{key, true, tree.null, tree.null, tree.null}

	// init y and x - y will be parent of x and x will be the place where z is to
	// be stored eventually
	var y *Node = tree.null
	x := tree.root
	for x != tree.null {
		y = x
		if key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}

	// after finding correct position for x/z, set z.p to y and check if z will be
	// root node, left child or right child and set it accordingly
	z.parent = y
	if y == tree.null {
		tree.root = z
	} else if key < y.key {
		y.left = z
	} else {
		y.right = z
	}

	// call insert fixup function on z
	tree.insertFixup(z)
}

func (tree *Tree) insertFixup(z *Node) {
	for z.parent.colorIsRed {
		if z.parent == z.parent.parent.left {
			if uncle := z.parent.parent.right; uncle.colorIsRed {
				z.parent.colorIsRed = false
				uncle.colorIsRed = false
				z = z.parent.parent
				z.colorIsRed = true
				continue
			}
			if z == z.parent.right {
				z = z.parent
				tree.leftRotate(z)
			}
			z.parent.colorIsRed = false
			// if z.parent.parent != nil {
			z.parent.parent.colorIsRed = true
			tree.rightRotate(z.parent.parent)
			// }
		} else {
			if uncle := z.parent.parent.left; uncle.colorIsRed {
				z.parent.colorIsRed = false
				uncle.colorIsRed = false
				z = z.parent.parent
				z.colorIsRed = true
				continue
			}
			if z == z.parent.left {
				z = z.parent
				tree.rightRotate(z)
			}
			z.parent.colorIsRed = false
			// if z.parent.parent != nil {
			z.parent.parent.colorIsRed = true
			tree.leftRotate(z.parent.parent)
			// }
		}
	}
	tree.root.colorIsRed = false
}
