package task

import (
	"day05/common/tree"
)

func AreToysBalanced(root *tree.Node) bool {
	if root == nil {
		return false
	}
	return getSummary(root.Left) == getSummary(root.Right)
}

func getSummary(root *tree.Node) int {
	switch {
	case root == nil:
		return 0
	case root.HasToy:
		return getSummary(root.Left) + getSummary(root.Right) + 1
	default:
		return getSummary(root.Left) + getSummary(root.Right)
	}
}
