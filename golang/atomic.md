原子操作由**底层硬件**支持，而锁则由操作系统的**调度器**实现。锁应当用来保护一段逻辑，对于一个变量更新的保护，原子操作通常会更有效率，并且更能利用计算机多核的优势，如果要更新的是一个复合对象，则应当使用`atomic.Value`封装好的实现。

原子操作是无锁的，常常直接通过CPU指令直接实现。

使用互斥锁和通道的思路都是在线程获得到访问权后阻塞其他线程对共享内存的访问，而使用原子操作解决数据竞争问题则是利用了其不可被打断的特性。

```
	var t uint32
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddUint32(&t, 1)
		}()
	}
	wg.Wait()
	fmt.Println(atomic.LoadUint32(&t))
```


```
var ApproveWindows atomic.Value

// validateApproveWindows
func (cm *constraintManager) startApproveWindowsTicker() {
    ticker := time.NewTicker(approveWindowsTicker)
    defer ticker.Stop()

    ctx := context.Background()
    defer ctx.Done()

    for {
        select {
        case <-ctx.Done():
            return
        case <-ticker.C:
            approveWindows, err := cm.opt.WhiteListStatus.ListApproveWindows(ctx)
            if err != nil {
                glog.Error("list approve windows failed with:%v", err)
                break
            }
            // atomic 
            cm.ApproveWindows.Store(approveWindows)
        }
    }
}


// use cache
var approveWindows []storagetypes.ApproveWindow
if cm.ApproveWindows.Load() != nil {
    approveWindows = cm.ApproveWindows.Load().([]storagetypes.ApproveWindow)
} else {
    var err error
    approveWindows, err = cm.opt.WhiteListStatus.ListApproveWindows(ctx)
    if err != nil {
        return false, "", err
    }
}
```
