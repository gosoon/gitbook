package main

// ()())))()
func minInsertions(s string) int {
	var needRight int
	var res int
	for _, b := range s {
		if b == '(' {
			needRight += 2
			if needRight%2 == 1 {
				res++
				needRight--
			}
		} else {
			needRight--
			if needRight == -1 {
				res++
				needRight = 1
			}
		}
	}
	return res + needRight
}
