package dfs

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	if m <= 0 {
		return 0
	}
	n := len(grid[0])

	var area int
	var dfs func(i int, j int)
	dfs = func(i int, j int) {
		if i >= m || i < 0 || j >= n || j < 0 {
			return
		}
		if grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		area++
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
		return
	}

	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				area = 0
				dfs(i, j)
				//fmt.Printf("i:%v j:%v area:%v \n", i, j, area)
				if area > ans {
					ans = area
				}
			}
		}
	}

	return ans
}
