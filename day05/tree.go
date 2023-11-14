package main

import "fmt"

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

/*
				____1____
			 /		     \
			0			      1
		 / \		    /	   \
	  1	  1      1		  0
	 /\		/\    /\		  /\
	1  0 0  1  0  0    1  1
*/

func main() {
	t := Create(true)
	t.Left = Create(false)
	t.Left.Left = Create(true)
	t.Left.Left.Left = Create(true)
	t.Left.Left.Right = Create(false)
	t.Left.Right = Create(true)
	t.Left.Right.Left = Create(false)
	t.Left.Right.Right = Create(true)
	t.Right = Create(true)
	t.Right.Left = Create(true)
	t.Right.Left.Left = Create(false)
	t.Right.Left.Right = Create(false)
	t.Right.Right = Create(false)
	t.Right.Right.Left = Create(true)
	t.Right.Right.Right = Create(true)
	fmt.Println("Are toys balanced:", AreToysBalanced(t))
	fmt.Println("Unroll girland:", unrollGirliand(t))
}

func Create(value bool) *TreeNode {
	return &TreeNode{
		HasToy: value,
		Left:   nil,
		Right:  nil,
	}
}

func AreToysBalanced(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return getSummary(root.Left) == getSummary(root.Right)
}

func getSummary(root *TreeNode) int {
	switch {
	case root == nil:
		return 0
	case root.HasToy:
		return getSummary(root.Left) + getSummary(root.Right) + 1
	default:
		return getSummary(root.Left) + getSummary(root.Right)
	}
}

func unrollGirliand(root *TreeNode) []bool {
	if root == nil {
		return []bool{}
	}

	queue := make([]*TreeNode, 1)
	queue[0] = root

	values := make([]bool, 1)
	values[0] = root.HasToy

	isEvenLevel := true

	for len(queue) > 0 {
		levelSize := len(queue)

		for i, j := 0, levelSize; i < levelSize; i, j = i+1, j-1 {
			if isEvenLevel {
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
					values = append(values, queue[i].Left.HasToy)
				}
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
					values = append(values, queue[i].Right.HasToy)
				}
			} else {
				if queue[j-1].Right != nil {
					queue = append(queue, queue[j-1].Right)
					values = append(values, queue[j-1].Right.HasToy)
				}
				if queue[j-1].Left != nil {
					queue = append(queue, queue[j-1].Left)
					values = append(values, queue[j-1].Left.HasToy)
				}
			}
		}

		isEvenLevel = false
		queue = queue[levelSize:]
	}

	return values
}
