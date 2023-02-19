package tree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var maxNum int
	var maxNumIndex int
	for idx, num := range nums {
		if num > maxNum {
			maxNum = num
			maxNumIndex = idx
		}
	}

	root := &TreeNode{Val: maxNum}
	root.Left = constructMaximumBinaryTree(nums[0:maxNumIndex])
	root.Right = constructMaximumBinaryTree(nums[maxNumIndex+1 : len(nums)])
	return root
}
