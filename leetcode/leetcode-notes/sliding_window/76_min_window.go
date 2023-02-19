package main

func minWindow(s string, t string) string {
	needs := make(map[byte]int)
	windows := make(map[byte]int)

	for idx := range t {
		needs[t[idx]]++
	}

	var left, right int
	var valid int
	var start int
	windowsLen := len(s) + 1
	for right < len(s) {
		c := s[right]
		windows[s[right]]++
		// scale up windows
		right++

		if windows[c] == needs[c] {
			valid++
		}

		//fmt.Printf("windows:%v, left:%v,right:%v,valid:%v \n", windows, left, right, valid)
		// check scale down windows
		for valid == len(needs) {
			if right-left < windowsLen {
				start = left
				windowsLen = right - left
			}

			d := s[left]
			left++
			// update windows
			if windows[d] == needs[d] {
				valid--
			}
			windows[d]--
		}
	}

	if windowsLen > len(s) {
		return string("")
	}
	return string(s[start : start+windowsLen])
}
