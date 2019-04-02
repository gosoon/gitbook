golang 信号处理

```
    //系统信息注册与资源释放
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-sigs
        fmt.Println()
        os.Exit(0)
    }()
```

k8s 中对信号的处理:

k8s.io/kubernetes/staging/src/k8s.io/apiserver/pkg/server/signal.go

