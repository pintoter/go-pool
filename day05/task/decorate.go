package task

import (
	"day05/common/tree"
)

func unrollGirliand(root *tree.Node) []bool {
	if root == nil {
		return []bool{}
	}

	queue := make([]*tree.Node, 1)
	queue[0] = root

	values := make([]bool, 1)
	values[0] = root.HasToy

	isEvenLevel := true

	for len(queue) > 0 {
		levelSize := len(queue)

		reverseSlice(queue)

		for i := 0; i < levelSize; i++ {
			if isEvenLevel == true {
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
					values = append(values, queue[i].Left.HasToy)
				}
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
					values = append(values, queue[i].Right.HasToy)
				}
			} else {
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
					values = append(values, queue[i].Right.HasToy)
				}
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
					values = append(values, queue[i].Left.HasToy)
				}
			}
		}

		isEvenLevel = !isEvenLevel

		queue = queue[levelSize:]
	}

	return values
}

func reverseSlice(slice []*tree.Node) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
