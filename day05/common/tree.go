package common

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

// func main() {
// 	t := Create(true)
// 	t.Left = Create(false)
// 	t.Left.Left = Create(true)
// 	t.Left.Left.Left = Create(true)
// 	t.Left.Left.Right = Create(false)
// 	t.Left.Right = Create(true)
// 	t.Left.Right.Left = Create(false)
// 	t.Left.Right.Right = Create(true)
// 	t.Right = Create(true)
// 	t.Right.Left = Create(true)
// 	t.Right.Left.Left = Create(false)
// 	t.Right.Left.Right = Create(false)
// 	t.Right.Right = Create(false)
// 	t.Right.Right.Left = Create(true)
// 	t.Right.Right.Right = Create(true)
// 	fmt.Println("Are toys balanced:", AreToysBalanced(t))
// 	fmt.Println("Unroll girland:", unrollGirliand(t))
// }

func Create(value bool) *TreeNode {
	return &TreeNode{
		HasToy: value,
		Left:   nil,
		Right:  nil,
	}
}
