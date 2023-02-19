package main

import "fmt"

func closestNodes(root *TreeNode, queries []int) [][]int {
	ans := make([][]int, len(queries))
	def := -1
	var treeValList []int
	for idx := range ans {
		ans[idx] = []int{def, def}
	}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		/*
			for i := 0; i < len(queries); i++ {
				if root.Val <= queries[i] && root.Val >= ans[i][0] {
					ans[i][0] = root.Val
				}
				if root.Val >= queries[i] {
					if ans[i][1] == def {
						ans[i][1] = root.Val
					} else if root.Val <= ans[i][1] {
						ans[i][1] = root.Val
					}
				}
			}
		*/
		treeValList = append(treeValList, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	for _, treeVal := range treeValList {
		for i := 0; i < len(queries); i++ {
			if treeVal <= queries[i] && treeVal >= ans[i][0] {
				ans[i][0] = treeVal
			}
			if treeVal >= queries[i] {
				if ans[i][1] == def {
					ans[i][1] = treeVal
				} else if treeVal <= ans[i][1] {
					ans[i][1] = treeVal
				}
			}
		}
	}

	fmt.Printf("ans:%+v \n", ans)
	return ans
}
