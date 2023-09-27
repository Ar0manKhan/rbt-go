package main

func main() {

	tree := Tree{}
	testArr := []int{984, 899, 487, 641, 34, 501, 551, 286, 814, 935, 301, 822, 788, 918, 532, 901, 387, 837, 41, 378, 1, 153, 85, 176, 305, 955}

	for _, val := range testArr {
		tree.Insert(val)
	}
	verify_rbt_properties(&tree)
}
