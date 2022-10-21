分析 gc 的几种方式
1.通过 pprof 中的信息分析 goroutine gc 相关指标
2.通过GODEBUG=gctrace=1可以开启gc日志，查看gc的结果信息
```
GODEBUG=gctrace=1 go run main.go 
```
3.通过 runtime.ReadMemStats() 调用分析，但是该方法开销比较大，需要执行 STW.
