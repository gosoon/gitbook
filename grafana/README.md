1.时间范围使用
```
where create_time/1000 >= $__unixEpochFrom() and create_time/1000 <= $__unixEpochTo()
```

2.跳转链接，__cell:0表示使用表格中第0个字段
```
https://grafana.com/d/rHhmPUAGz/service=${__cell:0}&var-day=All
```
