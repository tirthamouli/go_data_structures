package linkedlist

func findFromEnd(n *Node, index int) (bool, int) {
	// Step 1: Set the default
	var nth *Node

	// Step 2: Calculate the nth element
	for i := n; i != nil; i = i.next {
		if index == 0 {
			nth = i
			break
		}
		index--
	}

	// Step 3: Check if we were able to calculate the nth element
	if index != 0 {
		return false, 0
	}

	// Step 4: Get the nth element from the end
	j := n
	for i := nth; i != nil; i = i.next {
		j = j.next
	}

	// Step 5: Return the result
	return true, j.Val
}

func findFromStart(n *Node, index int) (found bool, res int) {
	// Step 1: Loop through all the elements
	for i := n; i != nil; i = i.next {
		if index == 0 {
			found = true
			res = i.Val
		}
		index--
	}

	// Step 2: Return the result
	return found, res
}

// Find - Gets the element at an index, can be negative as well
func (n *Node) Find(index int) (bool, int) {
	// Step 1: Check if the index is negative
	if index < 0 {
		return findFromEnd(n, index*-1)
	}

	// Step 2: Index is not negative here
	return findFromStart(n, index)
}

// GetMid - Gets the middle element (lower in case of even number)
func (n *Node) GetMid() *Node {
	s, m := n, n
	for s != nil {
		if s == nil || s.next == nil {
			break
		}
		m = m.next
		s = s.next.next
	}
	return m
}
