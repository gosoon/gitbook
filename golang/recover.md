当函数发生panic，函数的后续执行会立即被终止。接着，内部的defer函数会被执行。之后将panic传递给该函数的调用者。重复上述的流程，直到当前goroutine内的函数全部返回。之后程序打印panic传递的信息，紧跟着打印调用栈的信息。最后该goroutine终止。

recover用来阻止panic，恢复程序正常执行。但需要注意以下几点：
1. recover仅仅在defer函数中生效
2. recover只有跟panic同处于相同的goroutine时才工作。
3. revover之后不会打印panic调用栈信息


```
        defer func() {
            if r := recover(); r != nil {
                glog.Errorf("%s \n %s", r, string(debug.Stack()))
            }
        }()
```

4.无法捕获的panic

只有 runtime.panic 抛出的异常可以使用 recover 捕获，其他使用runtime.throw 抛出的 panic 无法进行捕获，比如 map 并发读写的 panic 无法进行捕获

可以捕获的 panic：
1、除0错误
	panic: runtime error: integer divide by zero
2、空指针问题
3、下标越界问题
	panic: runtime error: index out of range
4、...

不可捕获的 panic：
1、栈内存不足
	runtime: goroutine stack exceeds 1000000000-byte limit
	runtime: sp=0xc020161bd8 stack=[0xc020160000, 0xc040160000]
	fatal error: stack overflow
2、尝试使用 goroutine 执行一个 nil function
3、死锁
	fatal error: all goroutines are asleep - deadlock!
4、...


ref: https://www.jianshu.com/p/15c459c85141
