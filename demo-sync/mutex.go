package main

import (
	"fmt"
	"sync"
	"time"
)
/**
Mutex是互斥锁
当一个变量被上了互斥锁后，其他访问该变量的线程会被堵塞，不可对该变量进行读写操作，直到锁被释放。

如果有很多goroutine并发执行的话就会存在一个问题，因为某个线程获得互斥锁后，其他的goroutine被堵塞，导致程序的效率较低
 */
var m *sync.Mutex

func read(i int)  {
	fmt.Println(i, "begin lock")
	m.Lock()

	fmt.Println(i, "in lock")

	m.Unlock()
	fmt.Println(i, "unlock")
}

func main()  {
	m = new(sync.Mutex)

	go read(1)
	go read(2)

	time.Sleep(time.Second)
}

/*

1 begin lock
1 in lock
1 unlock
2 begin lock
2 in lock
2 unlock


1 begin lock
2 begin lock
1 in lock
1 unlock
2 in lock
2 unlock

*/