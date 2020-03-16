package lib

// Finding max of two values
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Calculates the current height of the node
func heightUpdate(l, r *Node) int {
	// Step 1: 0 Height
	height := 0

	// Step 2: Calculate the max height
	if l == nil && r == nil {
		height = 1
	} else if l == nil {
		height = r.Height + 1
	} else if r == nil {
		height = l.Height + 1
	} else {
		height = max(l.Height, r.Height) + 1
	}

	// Step 3: Return the height
	return height
}

// Balance the tree
func (n *Node) balanceTree(first, second string) *Node {
	// Step 1: LL insertion
	if first == "LEFT" && second == "LEFT" {
		// Step 1.1: If right has greater height, then no balance required
		if n.Right != nil && (n.Right.Height >= n.Left.Height || n.Left.Height == n.Right.Height+1) {
			return n
		}

		// Step: 1.2: Balance required
		newR := n.Left
		n.Left.Right, n.Left = n, n.Left.Right

		// Step 1.3: Update height
		n.Height = heightUpdate(n.Left, n.Right)

		// Step 1.3: Return the root
		return newR
	}

	// Step 2: RR insertion
	if first == "RIGHT" && second == "RIGHT" {
		// Step 2.1: If right has greater height, then no balance required
		if n.Left != nil && (n.Left.Height >= n.Right.Height || n.Right.Height == n.Left.Height+1) {
			return n
		}

		// Step: 2.2: Balance required
		newR := n.Right
		n.Right.Left, n.Right = n, n.Right.Left

		// Step 2.3: Update height
		n.Height = heightUpdate(n.Left, n.Right)

		// Step 2.3: Return the root
		return newR
	}

	// Step 3: LR insertion
	if first == "LEFT" && second == "RIGHT" {
		// Step 3.1: If right has greater height, then no balance required
		if n.Right != nil && (n.Right.Height >= n.Left.Height || n.Left.Height == n.Right.Height+1) {
			return n
		}

		// Step 3.2: Balance required
		newR := n.Left.Right
		n.Left.Right.Left, n.Left.Right.Right, n.Left.Right, n.Left = n.Left, n, n.Left.Right.Left, n.Left.Right.Right

		// Step 3.3: Height update for the nodes
		n.Height = heightUpdate(n.Left, n.Right)
		newR.Left.Height = heightUpdate(newR.Left.Left, newR.Left.Right)

		// Step 3.4: Return result
		return newR
	}

	// Step 4: RL insertion
	if first == "RIGHT" && second == "LEFT" {
		// Step 4.1: If right has greater height, then no balance required
		if n.Left != nil && (n.Left.Height >= n.Right.Height || n.Right.Height == n.Left.Height+1) {
			return n
		}

		// Step 4.2: Balance required
		newR := n.Right.Left
		n.Right.Left.Right, n.Right.Left.Left, n.Right.Left, n.Right = n.Right, n, n.Right.Left.Right, n.Right.Left.Left

		// Step 4.3: Height update - previous root
		n.Height = heightUpdate(n.Left, n.Right)

		// Step 4.4: Current left
		newR.Right.Height = heightUpdate(newR.Right.Left, newR.Right.Right)

		// Step 4.5: Return result
		return newR
	}

	// Step 5: Default
	return n
}

// Insert - Insert a node and balance the tree if required. lR: NO, not inserted. LEFT - inserted to the left sub tree and RIGHT - inserted to the right subtree
func (n *Node) Insert(val int) (r *Node, lR string) {
	// Step 1: No need to insert if value already found
	if n.Value == val {
		return n, "NO"
	}

	// Step 2: Check if value is greater
	if val > n.Value {
		// Step 2.1: Check if there is a right sub tree
		if n.Right == nil {
			// Step 2.1.1: New node
			newNode := &Node{
				Value:  val,
				Height: 1,
			}

			// Step 2.1.2: Insert the new node and update height
			n.Right = newNode
			if n.Left == nil {
				n.Height = n.Right.Height + 1
			} else {
				n.Height = max(n.Left.Height, n.Right.Height) + 1
			}

			// Step 2.1.3: Return the root
			return n, "RIGHT"
		}

		// Step 2.2: Insert into right
		newN, pos := n.Right.Insert(val)
		n.Right = newN

		// Step 2.3: Balance tree
		n = n.balanceTree("RIGHT", pos)

		// Step 2.4: Update height
		if n.Left == nil {
			n.Height = n.Right.Height + 1
		} else {
			n.Height = max(n.Left.Height, n.Right.Height) + 1
		}

		// Step 2.5: Return the result
		return n, "RIGHT"
	}

	// Step 3: Insert into right
	if n.Left == nil {
		// Step 3.1: New node
		newNode := &Node{
			Value:  val,
			Height: 1,
		}

		// Step 3.2: Insert the new node and update height
		n.Left = newNode
		if n.Right == nil {
			n.Height = n.Left.Height + 1
		} else {
			n.Height = max(n.Left.Height, n.Right.Height) + 1
		}

		// Step 3.3: Return the root
		return n, "LEFT"
	}

	// Step 4: Insert into right
	newN, pos := n.Left.Insert(val)
	n.Left = newN

	// Step 5: Balance tree
	n = n.balanceTree("LEFT", pos)

	// Step 6: Update height
	if n.Right == nil {
		n.Height = n.Left.Height + 1
	} else {
		n.Height = max(n.Left.Height+1, n.Right.Height+1)
	}

	// Step 7: Return the result
	return n, "LEFT"
}
