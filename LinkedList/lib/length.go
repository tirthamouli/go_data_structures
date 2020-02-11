package linkedlist

// Length - Returns the length of the linked list
func (n *Node) Length() (length int) {
	// Step 1: Go through all the nodes
	for i := n; i != nil; i = i.next {
		length++
	}

	// Step 2: Return the length
	return length
}
