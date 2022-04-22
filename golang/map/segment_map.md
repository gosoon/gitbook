构建一个高效cache主要要解决两个问题：
1、支持高并发场景：数据分片，降低锁的粒度；
2、实现零GC：
（1）无GC：分配对外内存（Mmap）；
（2）避免GC：

- map非指针优化(map[uint64]uint32)或者采用slice实现一套无指针的map；
- 数据存入[]byte slice(可考虑底层采用环形队列封装循环使用空间);

参考：
[golang本地缓存选型及原理总结](https://docs.google.com/presentation/d/1W2TIVAGFDe7ynjbhyGpc9SKB8Edg6ScUbwJyBWwyuUs/edit?usp=sharing)


```
package main

import (
	"fmt"
	"sync"

	"github.com/cespare/xxhash"
)

// TODO:
/*
	1.add expired time
*/

const (
	segmentCount    = 256
	segmentAndOpVal = 255
)

func NewFreeCache() *FreeCache {
	var f FreeCache
	for i := 0; i < segmentCount; i++ {
		f.CacheValue[i] = make(map[string]interface{})
	}

	return &f
}

type Interface interface {
	Set(key []byte, value interface{})
	Get(key []byte) (interface{}, error)
}

type FreeCache struct {
	SegmentMutex [segmentCount]sync.RWMutex
	CacheValue   [segmentCount]map[string]interface{}
}

// value: interface
// 如果分片数是2的n次方，则计算取余可以通过位运行
func (f *FreeCache) Set(key []byte, value interface{}) {
	hash := hashFunc(key)
	seg := hash & segmentAndOpVal

	f.SegmentMutex[seg].Lock()
	c := f.CacheValue[seg]
	c[string(key)] = value
	f.SegmentMutex[seg].Unlock()
}

func (f *FreeCache) Get(key []byte) (interface{}, error) {
	hash := hashFunc(key)
	seg := hash & segmentAndOpVal
	f.SegmentMutex[seg].RLock()
	c := f.CacheValue[seg]
	value, _ := c[string(key)]
	f.SegmentMutex[seg].RUnlock()
	return value, nil
}

// hashFunc
// 1. the key have a fixed hash64 code,code is number
func hashFunc(key []byte) uint64 {
	return xxhash.Sum64(key)
}

func main() {
	c := NewFreeCache()
	key := []byte("key")
	value := "value"
	c.Set(key, value)

	v, _ := c.Get(key)
	fmt.Println("value:", v)
}
```
