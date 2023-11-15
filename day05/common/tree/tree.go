package tree

type Node struct {
	HasToy bool
	Left   *Node
	Right  *Node
}

func Create(value bool) *Node {
	return &Node{
		HasToy: value,
		Left:   nil,
		Right:  nil,
	}
}
