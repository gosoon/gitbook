```
        defer func() {
            if err := recover(); err != nil {
                var buf [4096]byte
                n := runtime.Stack(buf[:], false)
                glog.Error(string(buf[:n]))
                debug.PrintStack()
                glog.Error(err)
            }
        }()
```
