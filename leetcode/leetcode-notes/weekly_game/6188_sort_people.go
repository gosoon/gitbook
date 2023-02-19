package main

import (
	"fmt"
	"sort"
)

func sortPeople(names []string, heights []int) []string {
	m := make(map[int]string)

	for i := 0; i < len(names); i++ {
		m[heights[i]] = names[i]
	}

	sort.Ints(heights)

	fmt.Println(m)
	fmt.Println(heights)
	j := 0
	for i := len(heights) - 1; i >= 0; i-- {
		fmt.Println(i)

		names[j] = m[heights[i]]
		j++
	}
	return names
}
