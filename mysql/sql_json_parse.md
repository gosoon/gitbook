1.
```
select JSON_UNQUOTE(json_extract(ext, '$.id')) as "id" from xxx
```

ref:https://erebao.com/2020/07/01/20200701_MYSQL%E4%BD%BF%E7%94%A8SQL%E8%AF%AD%E5%8F%A5%E8%A7%A3%E6%9E%90JSON/


2.

```
ROUND(sum(CASE entityType
		WHEN "QTCZ" THEN 
			if(json_extract(quota,'$.cpu') is NULL,0,json_extract(quota,'$.cpu'))
		WHEN "QTTH" THEN 
			if(json_extract(quota,'$.cpu') is NULL,0,-json_extract(quota,'$.cpu'))
	END),2) as "cpu净流入量",
```
