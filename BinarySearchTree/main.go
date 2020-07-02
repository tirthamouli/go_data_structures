package main

import (
	bst "github.com/Boom0027/go_data_structures/BinarySearchTree/lib"
)

func main() {
	a := &bst.Node{
		Val:   23,
		Left:  nil,
		Right: nil,
	}

	// Insert all the nodes
	vals := []int{2, 45, 1, 21, 34, 123, 22, 15, 9, 96, 45}
	for _, val := range vals {
		a.Insert(val)
	}
}
