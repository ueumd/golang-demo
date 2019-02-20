package main

import "fmt"

type Carer interface {
	GetName() string
	Run()
	DiDi()
}

type BMW struct {
	Name	string
}

//得实现接口所有的方法，才算实现了该接口
func (p *BMW) GetName() string  {
	return p.Name
}

//因此还得实现RUN()
//指针类型*BMW,也可以写值类型 (p BMW)
func (p *BMW) Run()  {
	fmt.Printf("%s is running\n", p.Name)
}

//还得实现DiDi()
func (p *BMW) DiDi()  {
	fmt.Printf("%s is didi\n", p.Name)
}


func main()  {
	/**
	/Golang中的接口，不需要显示的实现。
	只要一个变量，含有接口类型中的所有方法，那么这个变量就实现了这个接口。
	因此Golang中没有implement类似的关键字
	 */
	var car Carer
	bmw := &BMW{
		Name: "BMW",
	}
	bmw.Run()
	car = bmw
	car.Run()

}