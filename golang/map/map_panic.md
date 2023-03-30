map panic error:
```
fatal error: concurrent map iteration and map write

fatal error: concurrent map writes

fatal error: concurrent map read and map write
```

Go 的运行时在并发场景下，并不能准确地检测到所有的数据竞争，只能检测到一部分。此外，即使检测到了数据竞争，也并不一定会导致 panic 的触发，因为数据竞争的影响是不确定的，可能会导致程序出现意料之外的行为，而不一定是 panic。
