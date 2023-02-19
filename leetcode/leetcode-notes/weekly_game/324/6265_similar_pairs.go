package main

import (
	"fmt"
	"sort"
	"strings"
)

func similarPairs(words []string) int {
	m := make(map[int]string)

	for idx, word := range words {
		set := make(map[int32]struct{})
		for _, b := range word {
			set[b] = struct{}{}
		}

		var keys []int
		for k, _ := range set {
			keys = append(keys, int(k))
		}
		sort.Ints(keys)
		str := strings.Replace(strings.Trim(fmt.Sprint(keys), "[]"), " ", "-", -1)
		m[idx] = str
	}

	var result int
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if m[i] == m[j] {
				result++
			}
		}
	}

	return result
}
