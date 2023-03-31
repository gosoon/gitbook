1.时间范围使用
```
where create_time/1000 >= $__unixEpochFrom() and create_time/1000 <= $__unixEpochTo()
```

```
$__timeFrom()  
$__timeTo()
为格式化时间，eg: "2022-11-09 03:53:33"
```


2.跳转链接，__cell:0表示使用表格中第0个字段
```
https://grafana.com/d/rHhmPUAGz/service=${__cell:0}&var-day=All
```

3.ch 曲线图使用

```
SELECT
  timestamp*1000 as "time"
FROM xxx
WHERE stage = 'xxx'
ORDER by timestamp
```
