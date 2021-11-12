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

