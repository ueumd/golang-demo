package main

import (
	"fmt"
	"sync"
	"time"
)

func read4(s string) {
	time.Sleep(time.Second * 1)
	fmt.Println(s)
}

func main()  {
	var wg sync.WaitGroup

	var str = []string {
		"Hello, World",
		"Hello, Go",
		"Bye, PHP",
	}

	for _, s := range str {
		// Increment the WaitGroup counter.
		wg.Add(1)

		// Launch a goroutine to read the str.
		go func(s string) {
			defer wg.Done()
			read4(s)
		}(s)
	}

	// Wait for all goroutine to complete.
	wg.Wait()

	fmt.Println("done")
}

// Hello, Go
// Hello, World
// Bye, PHP
// done