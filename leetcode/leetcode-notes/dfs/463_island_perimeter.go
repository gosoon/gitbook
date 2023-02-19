package dfs

func islandPerimeter(grid [][]int) int {
	m := len(grid)
	if m <= 0 {
		return 0
	}
	var perimeter int
	n := len(grid[0])
	var dfs func(i int, j int) int
	dfs = func(i int, j int) int {
		if i >= m || i < 0 || j >= n || j < 0 {
			return 1
		}
		if grid[i][j] == 0 {
			return 1
		}
		if grid[i][j] != 1 {
			return 0
		}
		grid[i][j] = 2
		return dfs(i-1, j) + dfs(i+1, j) +
			dfs(i, j-1) + dfs(i, j+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return dfs(i, j)
			}
		}
	}
	return perimeter
}
