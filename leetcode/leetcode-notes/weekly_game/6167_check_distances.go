package main

import "fmt"

// s = "abaccb",
// distance = [1,3,0,5,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
func checkDistances(s string, distance []int) bool {
	l := 0
	r := len(s) - 1

	ok := true
	m := make(map[byte]int, len(s))
	for l < r {
		if _, found := m[s[l]]; !found {
			//start := len(m)
			for l < r && s[r] != s[l] {
				r--
			}

			if distance[s[l]-'a'] != (r - l - 1) {
				ok = false
			}
			r = len(s) - 1
			m[s[l]] = l
		}
		l++
	}
	fmt.Println(m)
	return ok
}
