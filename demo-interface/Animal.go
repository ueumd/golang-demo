package main

import "fmt"

//接口animal

type Animal interface {
	Speak() string
}
//Dog类实现animal接口
type Dog struct {}

func (d Dog) Speak() string {
	return "Woof!"
}
//Cat类实现animal接口
type Cat struct {}

func (c Cat) Speak() string {
	return "Meow!"
}
//Llama实现animal接口
type Llama struct {}

func (l Llama) Speak() string {
	return "?????"
}

//JavaProgrammer实现animal接口
type JavaProgrammer struct {}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}
//主函数
func main() {
	//利用接口实现多态
	animals := []Animal{
		Dog{},
		Cat{},
		Llama{},
		JavaProgrammer{},
	}
	// 多态
	for _, animal := range animals {
		fmt.Println(animal.Speak())  //打印不同实现该接口的类的方法返回值
	}
}
