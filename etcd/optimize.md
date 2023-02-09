1、使用 tmpfs 内存文件系统作为数据盘，缺点机器重启数据会不存在

2、压缩机制，时间间隔还是版本号

3、集群拆分，events 等数据分集群



阿里的优化方案：

1.自适应历史数据清理 compact 技术

2.基于 raft learner 的只读节点水平扩展能力

3.基于 raft learner 的热备节点

4.etcd 集群 QoS 能力



参考：

[阿里巴巴云原生 etcd 服务集群管控优化实践](https://mp.weixin.qq.com/s/i-sqZ7HbogdxinxhZrt7tw)

[蚂蚁集团万级规模 k8s 集群 etcd 高可用建设之路](https://mp.weixin.qq.com/s/_vLhUWh6dg8_26crQA7NpQ)

[K8s 集群稳定性：LIST 请求源码分析、性能评估与大规模基础服务部署调优](http://arthurchiao.art/blog/k8s-reliability-list-data-zh/)
