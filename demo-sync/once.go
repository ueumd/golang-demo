package main

import (
	"fmt"
	"sync"
	"time"
)

func read3() {
	fmt.Println(1)
}

func main()  {
	var once sync.Once

	for i :=0; i<10; i++ {
		go func() {
			once.Do(read3)
		}()
	}

	time.Sleep(time.Second)
}

// 1
// 最终只会打印出一次1。