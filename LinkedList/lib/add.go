package linkedlist

// Add - add element in the linked list, 0 index
func (n *Node) Add(value int, index int) *Node {
	// Step 1: Create the new node
	newNode := &Node{
		Val: value,
	}

	// Step 2: Check if we need to insert at the root
	if index == 0 {
		newNode.next = n
		return newNode
	}

	// Step 2: Check if insert at the end
	if index == -1 {
		// Step 2.1 Go to the last node
		for i := n; ; i = i.next {
			// Step 2.1.1: Set the next node as the new node
			if i.next == nil {
				i.next = newNode
				break
			}
		}
	} else {
		// Step 2.1: Go to the index at which we are supposed to insert
		for i := n; i != nil; i = i.next {
			// 2.1.1: Insert and update the next
			if index == 1 {
				newNode.next = i.next
				i.next = newNode
			}
			index--
		}
	}

	// Step 3: Return the root
	return n
}
