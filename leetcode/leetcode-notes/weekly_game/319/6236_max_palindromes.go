package main

func maxPalindromes(s string, k int) int {
	var l, r int
	var window []byte
	for r < len(s) {
		b := s[r]
		r++
		window = append(window,b)
		if len(window) == k 

	}

}
