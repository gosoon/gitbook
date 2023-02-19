package main

import "fmt"

func getDistances(arr []int) []int64 {
	ans := make([]int64, len(arr))
	m := make(map[int][]int)

	for i := 0; i < len(arr); i++ {
		m[arr[i]] = append(m[arr[i]], i)
	}

	for _, indexList := range m {
		n := len(indexList)
		left := make([]int, n)
		right := make([]int, n)

		i := 1
		j := n - 2
		for i < n {
			left[i] = left[i-1] + i*(indexList[i]-indexList[i-1])
			right[j] = right[j+1] + (n-j-1)*(indexList[j+1]-indexList[j])
			i++
			j--
		}

		fmt.Printf("indexList is %+v,left is %+v,right is %+v \n", indexList, left, right)
		for idx, index := range indexList {
			ans[int64(index)] = int64(left[idx]) + int64(right[idx])
		}

	}

	return ans
}
