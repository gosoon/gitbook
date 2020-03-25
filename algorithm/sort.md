
**binary_search.go**
```
package main

import "fmt"

func main() {

	s := []int{2, 3, 4, 5, 6, 7}
	t := int(3)
	//binarySearch(s, t)
	notRecursive(s, t)
}

// 若求下标则不能使用该方法
func binarySearch(s []int, t int) {
	length := len(s)
	if t == s[length/2] {
		fmt.Println(s[length/2])
		return
	}
	if t > s[length/2] {
		binarySearch(s[length/2:], t)
	} else {
		binarySearch(s[0:length/2], t)
	}
}

// https://leetcode-cn.com/problems/binary-search/submissions/
func search(nums []int, target int) int {
    l := len(nums)
    if l == 0 {
        return -1
    }
    left, right := 0, l-1
    for left <= right {
        mid := (right-left)/2 + left
        if target == nums[mid] {
            return mid
        }
        if target > nums[mid] {
            left = mid + 1
        } else {
            right = mid - 1
        }

    }
    return -1
}

```


bubbling_sort.go
```
package main

import "fmt"

func main() {
    s := []int{3, 4, 2, 5, 1, 7, 1, -1, 99}
    fmt.Println(s)
    bubblingSort(s)
    fmt.Println(s)
}

func bubblingSort(nums []int) {
    l := len(nums)
    for i := 0; i < l; i++ {
        var flag bool
        for j := 1; j < l-i; j++ {
            if nums[j-1] > nums[j] {
                nums[j-1], nums[j] = nums[j], nums[j-1]
                flag = true
            }
        }
        if !flag {
            break
        }
    }
}
```


select_sort.go
```
package main

import "fmt"

func main() {
    s := []int{3, 4, 2, 5, 1, 7}
    fmt.Println(s)
    selectSort(s)
    fmt.Println(s)
}

func selectSort(nums []int) {
    l := len(nums)
    for i := 0; i < l; i++ {
        for j := i + 1; j < l; j++ {
            if nums[i] > nums[j] {
                nums[i], nums[j] = nums[j], nums[i]
            }
        }
    }
}
```


quickSort
```
// https://www.kancloud.cn/digest/batu-go/153531
package main

import "fmt"

func main() {
	//s := []int{3, 4, 2, 5, 1, 7, 9, 12, -1, 123}
	s := []int{0, 1, 1, 2, 0}
	fmt.Println(s)
	//quickSort(s, 0, len(s)-1)
	fmt.Println(sortArray(s))
}

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	i, j, temp := start, end, nums[start]
	for i != j {
		for j > i && nums[j] >= temp {
			j--
		}
		for i < j && nums[i] <= temp {
			i++
		}

		//swap
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[start] = nums[i]
	nums[i] = temp
	quickSort(nums, start, i-1)
	quickSort(nums, j+1, end)
}
```

