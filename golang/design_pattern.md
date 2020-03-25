
### single



```
package main

import (
	"fmt"
	"sync"
)

type Interface interface {
	GetName() string
}

type demo struct {
	Name string
}

var demoOnce sync.Once

func NewDemo(name string) *demo {
	return &demo{Name: name}
}

func (d *demo) GetName() string {
	return d.Name
}

var d Interface

func getDemo(name string) Interface {
	demoOnce.Do(func() {
		d = NewDemo(name)
		fmt.Println("d.Name:", d.GetName())
	})

	if d == nil {
		fmt.Println("=== nil")
	}

	return d
}

func main() {
	d := getDemo("test1")
	fmt.Println("d.Name:", d.GetName())

	d = getDemo("test2")
	fmt.Println("d.Name:", d.GetName())
	//fmt.Println("d2.Name:", d2.Name)
}
```


