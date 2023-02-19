package main

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string(nil)
	}

	var ans []string
	var m = map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}
	var result []byte
	var backtrack func(start int)
	backtrack = func(start int) {
		if start > len(digits)-1 {
			ans = append(ans, string(append([]byte(nil), result...)))
			return
		}

		b := m[digits[start]]

		for i := 0; i < len(b); i++ {
			result = append(result, b[i])
			backtrack(start + 1)
			result = result[:len(result)-1]
		}
	}
	backtrack(0)
	return ans
}
