1.时间范围使用
```
where create_time/1000 >= $__unixEpochFrom() and create_time/1000 <= $__unixEpochTo()
```

```
$__timeFrom()  
$__timeTo()
为格式化时间，eg: "2022-11-09 03:53:33"
```

时间字符串使用:
```
updated_at >= $__timeFrom() and updated_at <= $__timeTo()
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

4、变量使用
使用变量的原始类型，避免变量从数字变为字符串

```
${var:raw}
```
ref:https://community.grafana.com/t/possible-to-have-an-int-as-a-template-variable/16593/2

