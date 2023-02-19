回溯法模板：

```

```



#### 优化：

1、回溯+备忘录

2、在创建临时切片时， append 中使用 nil 切片比空切片更节省空间：

```
res = append(res, append([]int(nil), tmp...))
```

而不是使用：

```
res = append(res, append([]int{}, tmp...))
```

3、如果数组中有重复的元素，在计算的结果中不能有重复的数组时一般需要给元数组排序；

4、递归时，需要判断求的结果是子序列还是全排列，如果是子序列一般使用当前索引加1作为自递归的起始，如果是全排序一般使用索引0作为子递归遍历数组时的起始索引；

5、记忆化递归？？



注意：

1、某些题目需要列出全排列，则需要在 backtrack 中使用for循环遍历数组，比如"77. 组合"：

```
func combineBacktrack(preSet []int, start int, n int, k int, results *[][]int) {
	if len(preSet) == k {
		result := []int{}
		result = append(result, preSet...)
		*results = append(*results, result)
		return
	}

	for i := start; i <= n; i++ {
		preSet = append(preSet, i)
		combineBacktrack(preSet, i+1, n, k, results)
		preSet = preSet[:len(preSet)-1]
	}
}
```

有些题目不需要全排列只需要最终得到的目标数组符合预期，比如 “494. 目标和”

```
func findTargetSumWaysBackTrack(nums []int, start int, target int, sum int, result *int) {
    // some return state
    if start == len(nums) {
        if sum == target {
            *result += 1
        }
        return
    }
    num := nums[start]
    findTargetSumWaysBackTrack(nums, start+1, target, sum+num, result)
    findTargetSumWaysBackTrack(nums, start+1, target, sum-num, result)
}
```





