package dfs

func closedIsland(grid [][]int) int {
	m := len(grid)
	if m <= 0 {
		return 0
	}
	n := len(grid[0])

	var dfs func(i int, j int)
	dfs = func(i int, j int) {
		if i >= m || i < 0 || j >= n || j < 0 {
			return
		}
		if grid[i][j] == 1 {
			return
		}
		grid[i][j] = 1
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	// left right column and first end row set 1
	for i := 0; i < n; i++ {
		dfs(0, i)
		dfs(m-1, i)
	}
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}

	//for i := 0; i < m; i++ {
	//fmt.Printf("grid:%v \n", grid[i])
	//}

	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				ans++
				dfs(i, j)
			}
		}
	}
	return ans
}
