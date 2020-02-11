package linkedlist

import "fmt"

// Print - Print the linked list
func (n *Node) Print() {
	// Loop through all the values and print it
	for i := n; i != nil; i = i.next {
		fmt.Print(i.Val, " -> ")
	}
	fmt.Println(nil)
}
