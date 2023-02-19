package main

import "fmt"

func smallestValue(n int) int {

	for {
		factorList := factor(n)
		if factorList[0] == n {
			return n
		}
		fmt.Println(factorList)
		var sum int
		for _, r := range factorList {
			sum += r
		}

		if sum == n {
			return sum
		}

		n = sum
	}

	/*
		for i := 2; i < n/2; i++ {
			if n%i == 0 {
				res := n / i
				n = res + i

				fmt.Println(i, res)
				// reset
				i = 1
			}

		}
	*/
	return n
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
