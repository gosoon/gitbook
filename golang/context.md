golang context 


[深入Golang之context](http://www.opscoder.info/golang_context.html)
[context 用法](https://github.com/Kevin-fqh/learning-k8s-source-code/blob/master/Go-package%E4%BD%BF%E7%94%A8/Context%E7%94%A8%E6%B3%95.md)

```
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		fmt.Println("time sleep")
		time.Sleep(2 * time.Second)
	}(ctx)
	time.Sleep(1 * time.Second)
	cancel()
}
```


```
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	select {
	case <-time.After(2 * time.Minute):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}
```


