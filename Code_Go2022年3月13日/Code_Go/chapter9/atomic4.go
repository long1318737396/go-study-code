package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var count int32
var locker sync.Mutex
var waiter sync.WaitGroup

func main() {
	start := time.Now()
	for i := 1; i <= 100000; i++ {
		waiter.Add(1)
		go addUseLock()
	}
	waiter.Wait()
	fmt.Println("加锁操作耗时：", time.Now().Sub(start))
	start = time.Now()
	for i := 1; i <= 100000; i++ {
		waiter.Add(1)
		go addUseAtomic()
	}
	waiter.Wait()
	fmt.Println("原子操作耗时：", time.Now().Sub(start))
	fmt.Println("count值为：", count)
}
func addUseLock() {
	locker.Lock()
	count++
	locker.Unlock()
	waiter.Done()
}
func addUseAtomic() {
	atomic.AddInt32(&count, 1)
	waiter.Done()
}
