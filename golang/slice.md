1.
越界不会 panic 的一个示例，原因待分析
参考：https://golang.org/ref/spec#Slice_expressions


```
func main() {
	a := []int{1, 2, 3}
	fmt.Println(len(a), cap(a), a[3])
	b := a[3:]

	fmt.Println(b)
}
```


2. some example 
(1) array by group 
```
const (
    queryNodeCountsByGroup = 10
)

group := int(math.Ceil(float64(len(nodeList)) / float64(queryNodeCountsByGroup)))
for i := 0; i < group; i++ {
    var nodeListGroup []string
    if i+1 == group {
        nodeListGroup = nodeList[i*queryNodeCountsByGroup : len(nodeList)]
    } else {
        nodeListGroup = nodeList[i*queryNodeCountsByGroup : (i+1)*queryNodeCountsByGroup]
    } 
}
```




