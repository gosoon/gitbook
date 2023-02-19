package main

func lengthOfLongestSubstring(s string) int {
	var l, r int

	windows := make(map[byte]int)
	var res int
	for r < len(s) {
		// scale up windows
		c := s[r]
		r++

		// update windows
		windows[c]++

		for windows[c] > 1 {
			c := s[l]
			l++
			windows[c]--
		}

		if r-l > res {
			res = r - l
		}
	}
	return res
}

/*
func lengthOfLongestSubstring(s string) int {
	fmt.Println(s)
	var l, r int

	windows := make(map[byte]int)
	var windowsLen int
	for r < len(s) {
		// scale up windows
		c := s[r]
		r++

		// update windows
		windows[c]++
		var haveRepeatChar bool
		if count := windows[c]; count > 1 {
			haveRepeatChar = true
		} else {
			if r-l > windowsLen {
				windowsLen = r - l
			}
		}

		fmt.Printf("windows:%v,l:%v,r:%v,windowsLen:%v \n", windows, l, r, windowsLen)
		for haveRepeatChar && l < r {
			// scale down windows
			fmt.Printf("-->windows:%v,l:%v,r:%v,windowsLen:%v \n", windows, l, r, windowsLen)
			c := s[l]
			l++

			windows[c]--
			haveRepeatChar = false
			for _, count := range windows {
				if count > 1 {
					haveRepeatChar = true
				}
			}
			fmt.Printf("-->windows:%v,l:%v,r:%v,windowsLen:%v,haveRepeatChar:%v \n", windows, l, r, windowsLen, haveRepeatChar)
		}
	}
	return windowsLen
}
*/
