package bst

// Find the required value, if not found then return the next greater element. Also insert if element is not found
func (n *Node) Find(val int) *Node {
	// Step 1: Check if the value is equal
	if n.Val == val {
		return n
	}

	// Step 2: Check for greater
	if val > n.Val {
		// When the right node is null
		if n.Right == nil {
			// Return value
			return nil
		}

		// When the right node is not null
		return n.Right.Find(val)
	}

	// Step 3: When the left node is null
	if n.Left == nil {
		// Return value
		return n
	}

	// Step 4: When not nil
	node := n.Left.Find(val)
	if node == nil {
		return n
	}

	// Step 5: Base return
	return node
}
