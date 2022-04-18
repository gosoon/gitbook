对于Go语言中的map, 垃圾回收器在 mark和scan阶段检查map中的每一个元素, 如果缓存中包含数百万的缓存对象，垃圾回收器对这些对象的无意义的检查导致不必要的时间开销。
Go 1.5中一个修复有关([#9477](https://github.com/golang/go/issues/9477)),这个issue还是描述了包含大量对象的map的垃圾回收时的耗时问题，Go的开发者优化了垃圾回收时对于map的处理，如果map对象中的key和value不包含指针，那么垃圾回收器就会对它们进行优化。

在使用 map 时尽量不要在 big map 中保存指针。

以下是针对在 map 中使用指针与非指针进行的测试：

```
const N = 100000000

func TestMapWithPointer(t *testing.T) {
	m := make(map[string]string)

	for i := 0; i < N; i++ {
		t := strconv.Itoa(i)
		m[t] = t
	}
	now := time.Now()
	runtime.GC()
	fmt.Println("time is ", time.Since(now))
	_ = m
}

func TestMapWithnoPointer(t *testing.T) {
	m := make(map[int]int)

	for i := 0; i < N; i++ {
		m[i] = i
	}
	now := time.Now()
	runtime.GC()
	fmt.Println("time is ", time.Since(now))
	_ = m
}

func TestMapWithValuePointer(t *testing.T) {
	m := make(map[int]string)

	for i := 0; i < N; i++ {
		t := strconv.Itoa(i)
		m[i] = t
	}
	now := time.Now()
	runtime.GC()
	fmt.Println("time is ", time.Since(now))
	_ = m
}

func TestMapWithKeyPointer(t *testing.T) {
	m := make(map[string]int)

	for i := 0; i < N; i++ {
		t := strconv.Itoa(i)
		m[t] = i
	}
	now := time.Now()
	runtime.GC()
	fmt.Println("time is ", time.Since(now))
	_ = m
}
```
执行结果如下：

```
=== RUN   TestMapWithPointer
time is  196.453594ms
--- PASS: TestMapWithPointer (91.66s)
=== RUN   TestMapWithnotPointer
time is  19.179509ms
--- PASS: TestMapWithnotPointer (32.26s)
=== RUN   TestMapWithValuePointer
time is  109.085364ms
--- PASS: TestMapWithValuePointer (52.84s)
=== RUN   TestMapWithKeyPointer
time is  188.302601ms
--- PASS: TestMapWithKeyPointer (71.03s)
PASS
ok  	_/Users/feiyu/test/map	248.944s
``` 

参考：
[Go语言使用 map 时尽量不要在 big map 中保存指针](https://www.jianshu.com/p/5903323a7110)
[妙到颠毫: bigcache优化技巧](https://colobu.com/2019/11/18/how-is-the-bigcache-is-fast/)
