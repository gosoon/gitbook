package dfs

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m <= 0 {
		return 0
	}
	var ans int
	n := len(grid[0])
	var dfs func(i int, j int)
	dfs = func(i int, j int) {
		if i >= m || i < 0 || j >= n || j < 0 {
			return
		}
		if grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return ans
}
