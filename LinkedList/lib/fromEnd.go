package linkedlist

// FromEnd - Returns the nth value from the end
func (n *Node) FromEnd(num int) (int, bool) {
	ptr1, ptr2 := n, n

	// Step 1: Increase the first pointer
	for ; ptr1 != nil && num >= 1; ptr1 = ptr1.next {
		num--
	}

	// Step 2: Check if number of element is not enough
	if num != 1 {
		return 0, false
	}

	// Step 3: Travel using two pointers
	for ; ptr1 != nil; ptr1 = ptr1.next {
		ptr2 = ptr2.next
	}

	// Step 4: Return value
	return ptr2.Val, true
}
