package main

import "fmt"

func offerIICoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			fmt.Printf("i:%v coin:%v dp[i]:%v 1+dp[i-coin]:%v \n", i, coin, dp[i], 1+dp[i-coin])
			dp[i] = min(dp[i], 1+dp[i-coin])
			//fmt.Println(dp)
		}
	}
	fmt.Println(dp)
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

//func min(x, y int) int {
//if x < y {
//return x
//}
//return y
//}
