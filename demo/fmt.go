package main

import "fmt"

/**
https://www.jianshu.com/p/9637c18d5f01
https://www.cnblogs.com/golove/p/3286303.html
Println 可以打印出字符串，和变量
func Println(a ...interface{}) (n int, err error)

Printf : 只可以打印出格式化的字符串,可以输出字符串类型的变量，不可以输出整形变量和整形
func Printf(format string, a ...interface{}) (n int, err error)


将转换结果以字符串形式返回。
func Sprint(a ...interface{}) string
func Sprintln(a ...interface{}) string
func Sprintf(format string, a ...interface{}) string


Scanf 从标准输入中读取数据，并根据格式字符串 format 对数据进行解析，将解析结果存入参数 a 所提供的变量中，变量必须以指针传入。
输入端的换行符必须和 format 中的换行符相对应（如果格式字符串中有换行符，则输入端必须输入相应的换行符）。
占位符 %c 总是匹配下一个字符，包括空白，比如空格符、制表符、换行符。
返回成功解析的参数数量。

Sscanf 用于扫描 str 中的数据，并根据 format 指定的格式，将 扫描出的 数据  填   写到  参数列表 a 中
当 str 中的数据，被全部扫描完毕或者扫描长度超出 format 指定的长度时
则停止扫描（换行符会被当作空格处理）
func Sscanf(str string, format string, a ...interface{}) (n int, err error)
 */

func demo1()  {
	var (
		username = "root"
		password = "123456"
	)
	config := fmt.Sprintf("%s:%s", username, password)
	r, _ := fmt.Println("config: ", config)  // config:  root:123456
	fmt.Printf(config)					// root:123456
	fmt.Printf("\n%s", config)	// root:123456

	fmt.Println("\nr:", r)


	a := 10
	// fmt.Printf(a) // 不可以打印整型
	fmt.Println("\n", a)
}

func demo2()  {

	var name string
	var age int

	// 注意：这里必须传递指针 &name, &age
	// 要获取的数据前后必须有空格

	// func Sscanf(str string, format string, a ...interface{}) (n int, err error)
	s2 := "我的名字叫Golang，今年1岁"
	fmt.Sscanf(s2, "%s, %d", &name, &age)
	fmt.Printf("%s %d", name, age) 		 	// 我的名字叫Golang，今年1岁 0

	fmt.Println()


	var name2 string
	var age2 int

	s := "我的名字叫 Golang ，今年 1 岁"
	fmt.Sscanf(s, "我的名字叫 %s ，今年 %d 岁", &name2, &age2)
	fmt.Printf("%s %d", name2, age2) 		// Golang 1

	fmt.Println()
	var t string
	token := "Bearer asdfsafask201163fasf"
	fmt.Sscanf(token, "Bearer %s", &t)
	fmt.Println(t) 									//asdfsafask201163fasf

}


func main() {
	demo2()
}
