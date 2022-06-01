
get count
```
    sql := "select sum(xx) from xx"
    var counts []int32
    err := m.client.Table(tableName).
        Raw(sql).
        Pluck("SUM(node) as count", &counts).
        Error
```

```
    var counts int64
    db := m.client.Table(xxx).Where("timestamp < ?", startTime).Count(&counts)
```


