package main

import (
	"fmt"
	"sync"
	"time"
)

var locker = new(sync.Mutex)
var cond =sync.NewCond(locker)

func read5(x int)  {
	cond.L.Lock()  // 获取锁

	cond.Wait()    // 等待通知

	fmt.Println(x)

	time.Sleep(time.Second * 1)

	cond.L.Unlock()   // 释放锁，不释放的话将只会有一次输出
}

func main()  {
	for i := 0; i < 40; i++ {
		go read5(i)
	}

	fmt.Println("start all")
	time.Sleep(time.Second * 1)

	cond.Broadcast() // 下发广播给所有等待的goroutine
	time.Sleep(time.Second * 60)
}