package bst

// Create a mew node
func create(val int) *Node {
	return &Node{
		Val:   val,
		Right: nil,
		Left:  nil,
	}
}

// Insert - Inserts a new node in the bst at the correct position
func (n *Node) Insert(val int) {
	// Step 1: Check if value already exists
	if n.Val == val {
		return
	}

	// Step 2: Check if value is greater
	if val > n.Val {
		// Step 2.1: When there is no right node
		if n.Right == nil {
			n.Right = create(val)
			return
		}

		// Step 2.2 When there is a right node
		n.Right.Insert(val)
		return
	}

	// Step 3: When less
	if n.Left == nil {
		n.Left = create(val)
		return
	}

	// Step 4: There is a left node
	n.Left.Insert(val)
}
