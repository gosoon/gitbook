1. Go静态代码检查:

```
$ golangci-lint run ./...
```

2. 检测代码中的数据竞争:
```
$ go run -race x.go
```

example:
```
func main() {
	m := make(map[string]string)
	go func() {
		m["a"] = "a"
	}()

	m["t"] = "t"

	fmt.Println(m)
}
```

```
$ go run -race race.go
map[t:t]
==================
WARNING: DATA RACE
Write at 0x00c000124180 by goroutine 7:
  runtime.mapassign_faststr()
      /Users/feiyu/go/src/runtime/map_faststr.go:202 +0x0
  main.main.func1()
      /Users/feiyu/test/race.go:8 +0x47

Previous write at 0x00c000124180 by main goroutine:
  runtime.mapassign_faststr()
      /Users/feiyu/go/src/runtime/map_faststr.go:202 +0x0
  main.main()
      /Users/feiyu/test/race.go:11 +0xb5

Goroutine 7 (running) created at:
  main.main()
      /Users/feiyu/test/race.go:7 +0x98
==================
Found 1 data race(s)
exit status 66
```



[golangci-lint](https://github.com/golangci/golangci-lint)
