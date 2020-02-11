package linkedlist

// CheckPalindrome - check if the given linked list is a plindrome
func (n *Node) CheckPalindrome() bool {
	// Step 1: Slow, fast and prev init
	s, f := n, n
	var prev *Node

	// Step 2: Loop and find the mid element and reverse all the elements on the way
	for true {
		if f == nil || f.next == nil {
			if f == nil {
				s, prev, f = prev, s, s
			} else {
				s, prev, f = prev, s, s.next
			}
			break
		}
		f = f.next.next
		s, s.next, prev = s.next, prev, s
	}

	// Step 3: Is palindrome init
	isPalindrome := true

	// Step 4: Check the elements and reverse again on the way to make them straight
	for f != nil {
		if s.Val != f.Val {
			isPalindrome = false
		}
		s, s.next, prev = s.next, prev, s
		f = f.next
	}

	// Step 5: Return the result
	return isPalindrome
}
