package main

import "fmt"

func main() {

	tree := GenerateTree()
	testArr := []int{1, 2, 3, 4, 5, 6, 7}

	for _, val := range testArr {
		fmt.Println("Inserting", val)
		tree.Insert(val)
	}
	verify_rbt_properties(tree)
}
