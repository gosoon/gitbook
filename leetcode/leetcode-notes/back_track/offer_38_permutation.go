package main

func permutation(s string) []string {
	if s == "" {
		return []string(nil)
	}

	globleCount := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		globleCount[s[i]]++
	}
	globleUsed := make(map[byte]int)

	var ans []string
	var result []byte
	var backtrack func()
	backtrack = func() {
		if len(result) == len(s) {
			ans = append(ans, string(append([]byte(nil), result...)))
			return
		}

		used := make(map[byte]bool)
		for i := 0; i < len(s); i++ {
			count := globleCount[s[i]]
			usedCount := globleUsed[s[i]]
			usedCount++
			if usedCount <= count && !used[s[i]] {
				result = append(result, s[i])
				used[s[i]] = true
				globleUsed[s[i]]++
				backtrack()
				result = result[:len(result)-1]
				globleUsed[s[i]]--
			}
		}
	}
	backtrack()
	return ans
}
