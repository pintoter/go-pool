package task

import (
	"day05/common"
)

func unrollGirliand(root *common.TreeNode) []bool {
	if root == nil {
		return []bool{}
	}

	queue := make([]*common.TreeNode, 1)
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
