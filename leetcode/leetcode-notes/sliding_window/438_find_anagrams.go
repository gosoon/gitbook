package main

func findAnagrams(s string, p string) []int {
	needs := make(map[byte]int)
	for i := range p {
		needs[p[i]]++
	}

	windows := make(map[byte]int)
	l, r := 0, 0

	var result []int
	var valid int
	for r < len(s) {
		// scale up windows
		c := s[r]
		r++
		windows[c]++

		if windows[c] == needs[c] {
			valid++
		}

		for valid == len(needs) {
			// update res
			if r-l == len(p) {
				result = append(result, l)
			}

			// scale down windows
			c := s[l]
			l++

			if windows[c] == needs[c] {
				valid--
			}
			windows[c]--
		}
	}
	return result
}
