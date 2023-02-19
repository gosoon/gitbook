package main

import "fmt"

func distinctPrimeFactors(nums []int) int {
	var result []int
	for _, num := range nums {
		result = append(result, factor(num)...)
	}
	fmt.Println(result)
	m := make(map[int]struct{})
	for _, res := range result {
		m[res] = struct{}{}
	}
	return len(m)
}

func factor(n int) []int {
	var result []int
	for i, j := 2, n; i <= n; i++ {
		if j%i == 0 {
			result = append(result, i)
			j /= i
			for j%i == 0 {
				result = append(result, i)
				j /= i
			}
		}
	}
	return result
}
