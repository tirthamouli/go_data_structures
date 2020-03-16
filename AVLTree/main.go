package main

import (
	"github.com/Boom0027/go_data_structures/AVLTree/lib"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 9, 8, 7, 6, 12, 11, 45, 32, 12}

	n := &lib.Node{
		Value:  10,
		Height: 1,
	}

	for _, v := range arr {
		n, _ = n.Insert(v)
	}

	n.Print()
}
