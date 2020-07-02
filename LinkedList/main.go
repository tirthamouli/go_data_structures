package main

import (
	"fmt"

	linkedlist "github.com/Boom0027/go_data_structures/LinkedList/lib"
)

func main() {
	// Creating the linked list
	n := &linkedlist.Node{
		Val: 123,
	}

	// Adding elements - not in root
	n.Add(1, -1)
	n.Add(2, -1)
	n.Add(3, -1)
	n.Add(4, -1)
	n.Add(3, -1)
	n.Add(2, -1)
	n.Add(1, -1)
	n.Add(111, 3) // 0 indexed

	// Delete by value and index
	_, n = n.DeleteVal(123)
	_, _, n = n.DeleteIndex(2)

	// Insert at root
	n = n.Add(7, 0) // Need to update root when inserting at 0
	n = n.Add(7, -1)

	// Printing result
	fmt.Println("Length:", n.Length()) // Lenght calculation
	fmt.Print("Values: ")
	n.Print()
	fmt.Print("Finding element: ")
	fmt.Println(n.Find(-9))
	fmt.Println("Middle element:", n.GetMid()) // Getting the middle element
	fmt.Println("Palindrome check:", n.CheckPalindrome())
	elem, found := n.FromEnd(3)
	fmt.Println("4th element from the end:", elem, "Found:", found)
	n.Print()
}
