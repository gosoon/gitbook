```
func StringListEqual(slice1 []string, slice2 []string) bool {
	var diff []string

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}
	if len(diff) != 0 {
		return true
	}
	return false
}

// diff slice1 and slice2,
// if in slice2 and not in slice1,need delete it,
// if in slice1 and not in slice2,need add it
func StringDifference(slice1 []string, slice2 []string) ([]string, []string) {
	var addList []string
	var deleteList []string

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found && i == 0 {
				addList = append(addList, s1)
			} else if !found && i == 1 {
				deleteList = append(deleteList, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return addList, deleteList
}
```

ref:
https://stackoverflow.com/questions/19374219/how-to-find-the-difference-between-two-slices-of-strings/19374861#19374861

