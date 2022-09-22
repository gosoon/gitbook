1. create table 
```
-- example 本地表
CREATE TABLE IF NOT EXISTS local_example (
	service_id 			       LowCardinality(String), 		-- comment 服务ID
	service_name 			   LowCardinality(String), 		-- comment 服务名称
	product_line 			   LowCardinality(String), 		-- comment 服务所在业务线
	stage 			           LowCardinality(String), 		-- comment 服务环境类型
	total_instance_count 	   UInt64, 					    -- comment 总实例数
	success_instance_count 	   UInt64, 				        -- comment 正常实例数
	success_rate               Float32,                     -- comment 可用率指标
	timestamp 		           UInt64, 					    -- comment 打点时间戳
	create_time	               DateTime,                    -- comment 打点时间
	date 				       Date 						-- comment 记录创建日期
)ENGINE=MergeTree()
ORDER BY (service_id,service_name,product_line,stage,timestamp,success_rate)
PARTITION BY date;

-- example 分布式表
CREATE TABLE IF NOT EXISTS example(
	service_id 			       LowCardinality(String), 		-- comment 服务ID
	service_name 			   LowCardinality(String), 		-- comment 服务名称
	product_line 			   LowCardinality(String), 		-- comment 服务所在业务线
	stage 			           LowCardinality(String), 		-- comment 服务环境类型
	total_instance_count 	   UInt64, 					    -- comment 总实例数
	success_instance_count 	   UInt64, 				        -- comment 正常实例数
	success_rate               Float32,                     -- comment 可用率指标
	timestamp 		           UInt64, 					    -- comment 打点时间戳
	create_time	               DateTime,                    -- comment 打点时间
	date 				       Date 						-- comment 记录创建日期
)ENGINE=Distributed('xxx', 'l_gifshow', 'local_example', rand());
```
2. timestamp convert
```
select toDateTime(timestamp/1000) from xxx;
```

3. argMin 与 argMax 的使用
存在多列时，可以获取某一列的最大或最小值对应另一列的值
