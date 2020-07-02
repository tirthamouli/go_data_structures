package main

import "fmt"

func main() {
	a := &node{
		val: 23,
	}

	// Insert all the nodes
	vals := []int{2, 45, 1, 21, 34, 123, 22, 15, 9, 96, 45}
	for _, val := range vals {
		fmt.Println(a.insert(val))
	}

}
