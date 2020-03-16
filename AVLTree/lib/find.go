package lib

// Find - find the given value and return a pointer to that node
func (n *Node) Find(val int) *Node {
	// Step 1: Check if the node is the current node
	if n.Value == val {
		return n
	}

	// Step 2: Go to the right node, when value is greater
	if val > n.Value {
		// Step 2.1: Check if there is a right child node
		if n.Right == nil {
			return nil
		}

		// Step 2.2: Search the right node
		return n.Right.Find(val)
	}

	// Step 3: Check if there is a left child node
	if n.Left == nil {
		return nil
	}

	// Step 4: Find the value in the left node
	return n.Left.Find(val)
}
