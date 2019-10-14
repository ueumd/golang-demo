package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// go 如何通过信号量控制并发量？

var wg08 sync.WaitGroup
var chSem chan int

func main()  {
	//通过管道定义信号量5个 意思代表只能有五个同时并发
	chSem = make(chan int, 5)

	//起100个任务并发处理
	for i:=0; i<100; i++ {
		wg08.Add(1)
		go getSqrt(i)//开辟协程处理
	}

	wg08.Wait() //等待组阻塞主协程
}

func getSqrt(n int)  {
	//规定:所有并发任务都必须注册在信号量管道里
	chSem <- n

	fmt.Printf("%d的平方根是%.2f\n", n, math.Sqrt(float64(n)))
	<- time.After(10 * time.Second)//定时器

	//任务结束后从信号量管道注销,给其它腾出空间
	<-chSem
	wg08.Done()
}

// https://learnku.com/index.php/articles/26151