package main

import "strings"

func isCircularSentence(sentence string) bool {
	strList := strings.Split(sentence, " ")
	l := len(strList)
	lastL := len(strList[l-1])
	if strList[0][0] != strList[l-1][lastL-1] {
		return false
	}

	for i := 0; i < len(strList); i++ {
		first := strList[i]
		firstL := len(strList[i])
		if i+1 < len(strList) {
			second := strList[i+1]
			if first[firstL-1] != second[0] {
				return false
			}
		}

	}
	return true
}
