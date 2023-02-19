package tree

import "fmt"

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	rootLeftDistance := make(map[int][]int)
	rootRightDistance := make(map[int][]int)
	var targetInLeft bool
	var targetToRootDistance int
	var traverse func(root *TreeNode, distanceMap map[int][]int, level int)
	traverse = func(root *TreeNode, distanceMap map[int][]int, level int) {
		if root == nil {
			return
		}

		if root.Left == target {
			targetInLeft = true
		}

		if root.Left == target || root.Right == target {
			targetToRootDistance = level
		}

		if level > 0 {
			distanceMap[level] = append(distanceMap[level], root.Val)
		}

		traverse(root.Left, distanceMap, level+1)
		traverse(root.Right, distanceMap, level+1)
	}
	traverse(root.Left, rootLeftDistance, 1)
	traverse(root.Right, rootRightDistance, 1)
	if root.Left == target {
		targetInLeft = true
	}

	if root.Left == target || root.Right == target {
		targetToRootDistance = 1
	}

	var ans []int
	fmt.Printf("targetInLeft:%v targetToRootDistance:%v\n", targetInLeft, targetToRootDistance)
	fmt.Printf("rootLeftDistance:%+v rootRightDistance:%+v \n", rootLeftDistance, rootRightDistance)
	if targetInLeft {
		leftNodeList := rootLeftDistance[targetToRootDistance+k]
		rightNodeList := rootRightDistance[k-targetToRootDistance]
		ans = append(ans, leftNodeList...)
		ans = append(ans, rightNodeList...)
	} else {
		leftNodeList := rootLeftDistance[k-targetToRootDistance]
		rightNodeList := rootRightDistance[targetToRootDistance+k]
		ans = append(ans, leftNodeList...)
		ans = append(ans, rightNodeList...)
	}
	return ans
}
