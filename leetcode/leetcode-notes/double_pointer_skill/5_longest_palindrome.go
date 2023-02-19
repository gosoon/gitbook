package main

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return s
	}

	var resLeft, resRight int
	for i := 1; i < len(s); i++ {
		left := i - 1
		right := i
		for left >= 0 && right < len(s) {
			if s[left] == s[right] {
				if right-left > resRight-resLeft {
					resLeft = left
					resRight = right
				}
				left--
				right++
			} else {
				break
			}
		}

		left = i - 1
		right = i + 1
		for left >= 0 && right < len(s) {
			if s[left] == s[right] {
				if right-left > resRight-resLeft {
					resLeft = left
					resRight = right

				}
				left--
				right++
			} else {
				break
			}
		}
	}
	return s[resLeft : resRight+1]
}
