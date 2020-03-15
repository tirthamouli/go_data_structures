package main

import "fmt"

type node struct {
	val   int
	left  *node
	right *node
}

func (n *node) insert(val int) {
	// No need to insert if element already exists
	if val == n.val {
		return
	}

	// Insert in the right or left
	if val > n.val {
		// Check if right sub tree is empty
		if n.right == nil {
			// Create a new node
			newNode := &node{
				val:   val,
				right: nil,
				left:  nil,
			}

			// Add the node to the current node
			n.right = newNode
		} else {
			n.right.insert(val)
		}
	} else {
		// Check if left sub tree is empty
		if n.left == nil {
			// Create a new node
			newNode := &node{
				val:   val,
				right: nil,
				left:  nil,
			}

			// Add the node to the current node
			n.left = newNode
		} else {
			n.left.insert(val)
		}
	}
}

// Find the required value, if not found then return the next greater element
func (n *node) find(val int) int {
	// Value found
	if val == n.val {
		return val
	}

	if val < n.val {
		if n.left == nil {
			return n.val
		}

		ans := n.left.find(val)
		if ans != 0 {
			return ans
		}

		return n.val

	}

	if n.right == nil {
		return 0
	}

	return n.right.find(val)
}

func main() {
	a := &node{
		val: 23,
	}
	fmt.Println(a)
	a.insert(22)
	a.insert(10)
	a.insert(5)
	fmt.Println(a.find(1))
}
