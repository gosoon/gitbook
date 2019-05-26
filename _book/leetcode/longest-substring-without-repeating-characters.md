

```
package main

import "fmt"

/*
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/
3. 无重复字符的最长子串
*/

func main() {
	s := "dvdf"
	fmt.Println(substring(s))
}

func substring(s string) int {
	m := make(map[string]int)
	i := 0
	j := 0
	max := 0
	for ; j < len(s); j++ {
		if _, exist := m[string(s[j])]; exist {
			if i < m[string(s[j])] {
				i = m[string(s[j])]
			}
		}
		if j-i+1 > max {
			max = j - i + 1
		}
		m[string(s[j])] = j + 1
	}
	return max
}
```
