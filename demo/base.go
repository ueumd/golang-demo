package main

import "fmt"

func add()  {
	var (
		a = 1
		b = 100
	)

	c := a + b
	fmt.Printf("a + b = %d", c)
}

func add100()  {
	res := 0
	for i:=1; i<= 100; i ++ {
		res += i
	}
	fmt.Printf("res = %d", res)
}

func add99()  {
	for i := 1; i<10; i++ {
		for j :=1; j<=i; j++ {
			fmt.Printf("%d x %d = %d \t", j, i, i*j)
		}
		fmt.Println("")
	}
}

func main () {
	// add()
	// add100()

	add99()
}