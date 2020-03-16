package lib

import "fmt"

// Print - Prints all the node
func (n *Node) Print() {
	fmt.Println(n)

	if n.Left != nil {
		n.Left.Print()
	}

	if n.Right != nil {
		n.Right.Print()
	}
}
