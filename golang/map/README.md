golang 内置的 map 没有实现线程安全，其不支持并发读写，官方为了解决此问题推出了 sync.Map，但官方的 sync.Map 应用场景不多，不能替代 map 的功能，想要实现线程安全，可以使用以下两种方法：

（1）加锁读写，使用 RWLock 

（2）分片加锁，效率更高，https://github.com/orcaman/concurrent-map

（3）使用开源的一些本地缓存系统；


参考：

[golang本地缓存(bigcache/freecache/fastcache等)选型对比及原理总结](https://mp.weixin.qq.com/s/YikGX5zj5Qzl4_i91ZZUHg)

[Go 为什么不在语言层面支持 map 并发？](https://eddycjy.com/posts/go/map-con/)
