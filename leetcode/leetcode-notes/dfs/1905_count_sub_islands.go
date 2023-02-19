package dfs

import "fmt"

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m := len(grid2)
	if m <= 0 {
		return 0
	}
	var isSubIsLand bool
	var ans int
	n := len(grid2[0])
	var dfs func(i int, j int)
	dfs = func(i int, j int) {
		if i >= m || i < 0 || j >= n || j < 0 {
			return
		}
		if grid2[i][j] == 0 {
			return
		}

		if grid1[i][j] != 1 {
			fmt.Printf("---> i:%v j:%v != 1 \n", i, j)
			isSubIsLand = false
		}
		grid2[i][j] = 0
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				isSubIsLand = true
				dfs(i, j)
				fmt.Printf("isSubIsLand is %v \n", isSubIsLand)
				if isSubIsLand {
					ans++
				}
			}
		}
	}
	return ans
}
