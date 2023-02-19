package main

func reverseString(s []byte) {
	i := 0
	j := len(s) - 1
	var t byte
	for i < j {
		t = s[i]
		s[i] = s[j]
		s[j] = t
		i++
		j--
	}
}
