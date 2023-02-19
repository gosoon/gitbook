package main

import "math/rand"

type Solution struct {
	size         int
	valueToIndex map[int]int
}

func Constructor(n int, blacklist []int) Solution {
	s := Solution{size: n - len(blacklist)}

	valueToIndex := make(map[int]int)
	last := n - 1
	blackMap := make(map[int]struct{})
	for _, black := range blacklist {
		blackMap[black] = struct{}{}
	}
	for _, black := range blacklist {
		if black < s.size {
			var found bool
			_, found := blackMap[last]
			for found {
				last--
				_, found = blackMap[last]
			}
			valueToIndex[black] = last
		}
	}
	s.valueToIndex = valueToIndex
	return s
}

func (this *Solution) Pick() int {
	index := rand.Intn(this.size)
	if value, found := this.valueToIndex[index]; found {
		return value
	}
	return index
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
