1. case 使用
```
SELECT fo.id AS ID, fo.product AS 产品
    , CASE fo.area
        WHEN 'A' THEN 'AA'
        WHEN 'B' THEN 'BB'
    END AS 区域
    ,
```

```
  ROUND(sum(CASE status_name
		WHEN "待交付" THEN total
		WHEN "交付中" THEN total
		WHEN "流程结束" THEN h_total
	END),2) as "cost",
```


2.字符串拼接
```
 concat(region,"-") as "xxx",   

 // group by 之后的字符串拼接
 group_concat(spec separator ' , ') as gpu_spec
```



