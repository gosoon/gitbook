1.
```
SELECT fo.id AS ID, fo.product AS 产品
    , CASE fo.area
        WHEN 'A' THEN 'AA'
        WHEN 'B' THEN 'BB'
    END AS 区域
    ,
```


字符串拼接
```
 concat(region,"-") as "xxx",   
```
