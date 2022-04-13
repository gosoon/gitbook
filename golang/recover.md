当函数发生panic，函数的后续执行会立即被终止。接着，内部的defer函数会被执行。之后将panic传递给该函数的调用者。重复上述的流程，直到当前gorouter内的函数全部返回。之后程序打印panic传递的信息，紧跟着打印调用栈的信息。最后该gorouter终止。

recover用来阻止panic，恢复程序正常执行。但需要注意以下几点：
1. recover仅仅在defer函数中生效
2. recover只有跟panic同处于相同的gorouter时才工作。
3. revover之后不会打印panic调用栈信息


```
        defer func() {
            if r := recover(); r != nil {
                glog.Errorf("%s \n %s", r, string(debug.Stack()))
            }
        }()
```
