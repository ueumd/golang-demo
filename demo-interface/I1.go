package main

import "fmt"

/*
interface类型默认是指针

接口的实现
Golang中的接口，不需要显示的实现。
只需要   一个变量   ， 含有  接口类型中的  所有   方法， 那么这个  变量   就实现这个接口

因为Golang中没有implement类似的关键字
如果一个变量含有了多个interface类型的方法，那么这个变量就实现了多个接口

如果一个变量只含有了1个interface的部分方法，那么这个变量没有实现这个接口..

一种事物的多种形态，都可以按照统一的接口进行操作
*/


// 定义test接口
type Test interface {
	// 接口有2个方法
	Print()
	Sleep()
}

// 定义结构体
type People struct {
	name string
	age int
}


type Student struct {
	name string
	age int
	score int
}

// Student 实现了Test的所有接口
func (p Student) Print()  {
	fmt.Println("name", p.name)
	fmt.Println("age", p.age)
	fmt.Println("score", p.score)
}

// 如果这个方法没有写，就没有实现所Test 接口
//func (p Student) Sleep()  {
//	fmt.Println("student sleep")
//}



// People 实现了Test的所有接口
func (people People) Print() {
	fmt.Println("name:", people.name)
	fmt.Println("age:", people.age)
}
func (p People) Sleep() {
	fmt.Println("People Sleep")
}

func main()  {
	//接口是个地址
	var t Test
	fmt.Println(t)

	var stu Student = Student{
		name:  "stu1",
		age:   20,
		score: 200,
	}

	fmt.Println(stu)
	//t = stu
	t.Print()
	t.Sleep()
}
