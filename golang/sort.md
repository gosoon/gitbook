```
package main

import (
	"sort"
	"fmt"
)

type StringSorter []string

// have StringSorter implement sort.Interface

func (s StringSorter) Len() int {
	return len(s)
}

func (s StringSorter) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s StringSorter) Swap(i, j int) {
	swap := s[i]
	s[i] = s[j]
	s[j] = swap
}

func main() {
	arr := []string{"5", "3", "20", "abc"}
	sort.Sort(StringSorter(arr))
	fmt.Println(arr)
}

```

replace with:

```
package main

import (
	"sort"
	"fmt"
)

func main() {
	arr := []string{"5", "3", "20", "abc"}
	sort.Sort(sort.Interface{
		Len: func() int { return len(arr) },
		Swap: func(i, j int) {
			swap := arr[i]
			arr[i] = arr[j]
			arr[j] = swap
		},
		Less: func(i, j int) bool { return arr[i] < arr[j] },
	})
	fmt.Println(arr)
}

```


[参考链接](https://dev.to/deanveloper/my-favorite-go-2-proposals-3lie)
