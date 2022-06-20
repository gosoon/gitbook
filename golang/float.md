### 1、浮点精度问题解决方案


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


### 2、异常浮点数的处理

```
func Decimal(value float64) float64 {
    if math.IsNaN(value) || math.IsInf(value, 0) {
        return float64(0)
    }

    value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
    return value
}
```
