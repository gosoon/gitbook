package main

import "fmt"

func sumOfNumberAndReverse(num int) bool {
	for i := num; i >= 0; i-- {
		reversNum := revers(i)
		fmt.Printf("num:%v reversNum:%v \n", i, reversNum)
		if reversNum+i == num {
			return true
		}
	}

	return false
}

func revers(num int) int {
	var reversNum int
	for num != 0 {
		remainder := num % 10
		num = num / 10
		reversNum = reversNum*10 + remainder
	}
	return reversNum
}
