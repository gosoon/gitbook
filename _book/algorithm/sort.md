
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
func notRecursive(s []int, t int) {
	length := len(s)
	start := 0
	end := length - 1
	for start <= end {
		// mid := (end-start)/2 + start
		mid := low+((high-low)>>1)
		if t == s[mid] {
			fmt.Println(s[mid])
			return
		}
		if t > s[mid] {
			start = mid
		} else {
			end = mid
		}
	}
}
```


bubbling_sort.go
```
package main

import "fmt"

func main() {
	s := []int{3, 4, 2, 5, 1, 7}
	fmt.Println(s)
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				t := s[j]
				s[j] = s[j+1]
				s[j+1] = t
			}
		}
	}
	fmt.Println(s)
}

/*
6,1,2,3,4
// 改进的冒泡排序算法
int[] bubbleSort(int[] arr){
    int flag ;
    for(int i = 0; i < arr.length; i++){
        flag = 0;
        for(int j = 1; j < arr.length - i; j++){
            if(a[j-1] > a[j]){
                int temp = a[i];
                a[i] = a[j];
                a[j] = temp;
                flag = 1;
            }
        }
        // flag = 0 ,则表示内层循环中，没有发生相邻元素的交换，即已排序，直接跳出外层循环
        if(flag == 0)
            break;
    }
    return arr;
}
*/
```


select_sort.go
```
package main

import "fmt"

func main() {
	s := []int{3, 4, 2, 5, 1, 7}
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				t := s[i]
				s[i] = s[j]
				s[j] = t
			}

		}
		fmt.Println(s)
	}
	fmt.Println(s)
}
```
