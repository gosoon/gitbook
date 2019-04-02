浮点精度问题解决方案


```
package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	a := float64(0.00000001)
	dec := decimal.NewFromFloat(a)

	fmt.Println(dec.String())
}
```

参考： https://toutiao.io/posts/245965/app_preview
