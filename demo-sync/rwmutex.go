package main

import (
	"fmt"
	"sync"
	"time"
)

var m2 *sync.RWMutex
var val = 0

func read2(i int)  {
	fmt.Println(i, "begin read")
	m2.RLock()  // 读锁  数据可以被多个goroutine并发访问但不可写
	time.Sleep(1 * time.Second)

	fmt.Println(i, "val: ", val)
	time.Sleep(1 * time.Second)

	m2.RUnlock()
	fmt.Println(i, "end read")
}

func write2(i int)  {
	fmt.Println(i, "begin write")
	m2.Lock() // 写锁 写锁时，数据不可被其他goroutine读或写
	val = 10
	fmt.Println(i, "val: ", val)
	time.Sleep(1 * time.Second)
	m2.Unlock()
	fmt.Println(i, "end write")
}

func main()  {
	m2 = new(sync.RWMutex)
	go read2(1)

	go write2(2)

	go read2(3)

	time.Sleep(5 * time.Second)

}

// 1 begin read
// 2 begin write
// 3 begin read
// 1 val:  0
// 1 end read
// 2 val:  10
// 2 end write
// 3 val:  10