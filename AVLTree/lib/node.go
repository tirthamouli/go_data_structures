package lib

// Node - node of an avl tree
type Node struct {
	Value  int
	Height int
	Left   *Node
	Right  *Node
}
