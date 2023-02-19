package main

func checkInclusion(s1 string, s2 string) bool {
	needs := make(map[byte]int)
	windows := make(map[byte]int)

	for idx := range s1 {
		c := s1[idx]
		needs[c]++
	}

	var l, r int
	var valid int
	for l < len(s2) && r < len(s2) {
		// scale up windows
		c := s2[r]
		r++
		// handle windows data
		windows[c]++
		count := windows[c]
		if ct, found := needs[c]; found && count <= ct {
			valid++
		}

		for valid == len(s1) {
			// set start and windowsLen
			if r-l == len(s1) {
				return true
			}

			// scale down windows
			c := s2[l]
			l++
			// handle windows data
			if windows[c] == needs[c] {
				valid--
			}
			windows[c]--
		}
	}

	return false
}
