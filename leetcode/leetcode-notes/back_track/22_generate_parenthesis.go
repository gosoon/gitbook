package main

import "fmt"

func generateParenthesis(n int) []string {
	var ans []string
	var result []byte
	parenthesisList := []byte{'(', ')'}
	var backtrack func(left int, right int)
	backtrack = func(left int, right int) {
		fmt.Printf("result:%+v left:%v,right:%v\n", string(result), left, right)
		if left == n && right == n {
			ans = append(ans, string(append([]byte(nil), result...)))
			return
		}

		for _, parenthesis := range parenthesisList {
			if left <= n && right <= n {
				if parenthesis == ')' && left > right {
					result = append(result, parenthesis)
					backtrack(left, right+1)
					result = result[:len(result)-1]
				} else if parenthesis == '(' {
					result = append(result, parenthesis)
					backtrack(left+1, right)
					result = result[:len(result)-1]
				}
			}
		}
	}
	backtrack(0, 0)
	return ans
}
