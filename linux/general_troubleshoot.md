1、查看 oom 进程

```
grep "Out of memory" /var/log/messages | grep envoy
```

2、使用 ulimit -a 以及 ps 查看进程数是否超过最大值

3、使用 strace 查看 D 进程卡住的位置

