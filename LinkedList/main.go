package main

import (
	"fmt"

	linkedlist "github.com/Boom0027/data_structure/LinkedList/lib"
)

func main() {
	// Creating the linked list
	n := &linkedlist.Node{
		Val: 123,
	}

	// Adding elements - not in root
	n.Add(3, -1)
	n.Add(4, -1)
	n.Add(5, -1)
	n.Add(6, -1)
	n.Add(7, -1)
	n.Add(8, -1)
	n.Add(9, -1)
	n.Add(111, 3) // 0 indexed

	// Delete by value and index
	_, n = n.DeleteVal(123)
	_, _, n = n.DeleteIndex(6)

	// Insert at root
	n = n.Add(32, 0) // Need to update root when inserting at 0
	n = n.Add(36, 4)
	n = n.Add(22, 4)

	// Printing result
	fmt.Println("Length:", n.Length()) // Lenght calculation
	fmt.Print("Values: ")
	n.Print()
	fmt.Print("Finding element: ")
	fmt.Println(n.Find(-9))
	fmt.Println("Middle element:", n.GetMid()) // Getting the middle element
}
