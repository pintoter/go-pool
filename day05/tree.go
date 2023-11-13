package main

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func areToysBalanced(t *TreeNode) bool {
	if t == nil {
		return false
	}
	return getSummary(t.Left) == getSummary(t.Right)
}

func getSummary(t *TreeNode) int {
	if t == nil {
		return 0
	}

	if t.HasToy {
		return getSummary(t.Left) + getSummary(t.Right) + 1
	}

	return getSummary(t.Left) + getSummary(t.Right)
}
