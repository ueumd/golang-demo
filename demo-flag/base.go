package main

import "flag"
import "fmt"

// https://www.jianshu.com/p/f9cf46a4de0e golang flag包使用笔记
// https://www.kancloud.cn/itfanr/go-by-example/81692
func main()  {


	/*
	第一种方式
	基础的标记声明适用于string，integer和bool型选项。
	这里我们定义了一个标记`word`，默认值为`foo`和一个简短的描述。
	`flag.String`函数返回一个字符串指 针（而不是一个字符串值），
	我们下面将演示如何使用这个指针
 */
	wordPrt := flag.String("word", "foo", "a string")

	// 这里定义了两个标记，一个`numb`，另一个是`fork`，
	// 使用和上面定义`word`标记相似的方法
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	/**
	第二种方式
	你也可以程序中任意地方定义的变量来定义选项
	只需要把该变量的地址传递给flag声明函数即可
	*/
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")


	/**
	命令行 flag 的语法有如下三种形式：
	base.exe -word=opt -numb=7 -fork -svar=flag hello aa bb
	-flag // 只支持bool类型
	-flag=x
	-flag x // 只支持非bool类型
	 */
	flag.Parse()

	fmt.Println("word", *wordPrt)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

/**
直接运行
go run basee.go

word foo
numb: 42
fork: false
svar: bar
tail: []

编译运行
go build base.go 生成exe文件


E:\gocode\src\myapiserver\demo-flag>base.exe -word=opt -numb=7 -fork -svar=flag
word opt
numb: 7
fork: true
svar: flag
tail: []


注意flag包要求所有的flag都必须出现在尾部位置参数的前面，否则这些flag将被当作位置参数处理
E:\gocode\src\myapiserver\demo-flag>base -word=opt a1 a2 a3
word opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3] 参数处理


E:\gocode\src\myapiserver\demo-flag>base.exe -word=opt -numb=7 -fork -svar=flag hello aa bb
word opt
numb: 7
fork: true
svar: flag
tail: [hello aa bb] 后三位参数处理
 */

}