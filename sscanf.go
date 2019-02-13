package main

import "fmt"

// Sscanf 用于扫描 str 中的数据，并根据 format 指定的格式
// 将扫描出的数据填写到参数列表 a 中
// 当 r 中的数据被全部扫描完毕或者扫描长度超出 format 指定的长度时
// 则停止扫描（换行符会被当作空格处理）
//func Sscanf(str string, format string, a ...interface{}) (n int, err error)

func main() {
	s := "我的名字叫 Golang ，今年 1 岁"
	var name string
	var age int

	// 注意：这里必须传递指针 &name, &age
	// 要获取的数据前后必须有空格

	fmt.Sscanf(s, "我的名字叫 %s ，今年 %d 岁", &name, &age)
	fmt.Printf("%s %d", name, age)
	// Golang 1
}
