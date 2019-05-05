make 的三个参数：

type,length,cap

cap must be re-slice and use

```
package main

import "fmt"

func main(){
	a := make([]int, 10, 20)
	fmt.Printf("%d, %d\n", len(a), cap(a))
	fmt.Println(a)
	b := a[:cap(a)]
	fmt.Println(b)
}
```
