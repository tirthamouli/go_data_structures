package linkedlist

// DeleteVal - Deletes by value, returns weather anything is deleted and the root
func (n *Node) DeleteVal(val int) (deleted bool, root *Node) {
	// Step 1: Set default value
	root = n
	var prev *Node

	// Step 2: Loop through all the nodes and delete if the value is found
	for i := n; i != nil; i = i.next {
		// Step 2.1: When found
		if i.Val == val {
			// Step 2.1.1: When we need to delete the root, the root is updated
			if prev == nil {
				root = i.next
			} else {
				prev.next = i.next
			}

			// Step 2.1.2: Set the deleted as true
			deleted = true
		}

		// Step 2.2: Setting the previous
		prev = i
	}

	// Step 3: Return the deleted and the root
	return deleted, root
}

// DeleteIndex - Delete the specified indexed, 0 indexed
func (n *Node) DeleteIndex(index int) (deleted bool, value int, root *Node) {
	// Step 1: Set default value
	root = n
	var prev *Node

	// Step 2: Check if we need to delete the index
	if index == 0 {
		root = n.next
		return false, n.Val, root
	}

	// Step 3: Loop through all the nodes and delete if the index is there
	for i := n; i != nil; i = i.next {
		// Step 3.1: When found
		if index == 0 {
			// Step 3.1.1: Delete index
			value = i.Val
			prev.next = i.next

			// Step 3.1.2: Set the deleted as true
			deleted = true
		}

		// Step 3.2: Setting the previous
		prev = i
		index--
	}

	// Step 4: Return the deleted and the root
	return deleted, value, root
}
