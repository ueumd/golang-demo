package main

import "fmt"

type Carer interface {
	GetName() string
	Run()
	DiDi()
}
type Test interface {
	Hello()
}
//如果一个变量含有了多个interface类型的方法，那么这个变量就实现了多个接口
//宝马车
type BMW struct {
	Name string
}

//因此还得实现RUN()
//得实现接口所有的方法，才算实现了该接口
func (p *BMW) GetName() string {
	return p.Name
}

//实现接口的Run()
func (p *BMW) Run() {
	fmt.Printf("%s is running\n", p.Name)
}

//还得实现DiDi()
func (p *BMW) DiDi() {
	fmt.Printf("%s is didi\n", p.Name)
}

func (p *BMW) Hello() {
	fmt.Printf("hello,i'm %s \n", p.Name)
}

//比亚迪
type BYD struct {
	Name string
}

func (p *BYD) GetName() string {
	return p.Name
}

func (p *BYD) Run() {
	fmt.Printf("%s is running\n", p.Name)
}
func (p *BYD) DiDi() {
	fmt.Printf("%s is didi\n", p.Name)
}


func main() {
	var car Carer
	var test Test
	fmt.Println(car)
	//var bmw BMW
	//bmw.Name = "BMW"
	//写法2
	bmw := &BMW{
		Name: "BMW",
	}
	car = bmw
	car.Run()

	test = bmw
	test.Hello()

	byd := &BYD{
		Name: "BYD",
	}
	car = byd
	car.Run()
}