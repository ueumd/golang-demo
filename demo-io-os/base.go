package main

import (
	"fmt"
	"os"
	"time"
)

/**
 os 常用导出函数
func Hostname() (name string, err error) 	// Hostname返回内核提供的主机名
func Environ() []string 				   // Environ返回表示环境变量的格式为”key=value”的字符串的切片拷贝
func Getenv(key string) string            // Getenv检索并返回名为key的环境变量的值
func Getpid() int                         // Getpid返回调用者所在进程的进程ID
func Exit(code int)                      // Exit让当前程序以给出的状态码code退出。一般来说，状态码0表示成功，非0表示出错。程序会立刻终止，defer的函数不会被执行
func Stat(name string) (fi FileInfo, err error)   // 获取文件信息
func Getwd() (dir string, err error)             // Getwd返回一个对应当前工作目录的根路径
func Mkdir(name string, perm FileMode) error     // 使用指定的权限和名称创建一个目录
func MkdirAll(path string, perm FileMode) error  // 使用指定的权限和名称创建一个目录，包括任何必要的上级目录，并返回nil，否则返回错误
func Remove(name string) error                   // 删除name指定的文件或目录
func TempDir() string                            // 返回一个用于保管临时文件的默认目录
var Args []string Args                           // 保管了命令行参数，第一个是程序名。
 */

// 一些常用操作
func os1()  {
	fmt.Println(os.Args) // [C:\Users\Administrator\AppData\Local\Temp\___go_build_base_go.exe]

	// 获取主机名
	fmt.Println(os.Hostname())   //XL-201807142350

	// 获取用户ID
	fmt.Println(os.Getuid())

	// 获取有效用户ID
	fmt.Println(os.Geteuid())

	// 获取有效组ID
	fmt.Println(os.Getgid())

	// 获取进程ID
	fmt.Println(os.Getpid())

	//获取父进程ID
	fmt.Println(os.Getppid())

	//获取环境变量的值
	fmt.Println(os.Getenv("JAVA_HOME")) //D:\installjava\jdk1.8\jdk1.8.0_91
	fmt.Println(os.Getenv("GOPATH")) // E:\gocode

	//设置环境变量的值
	fmt.Println(os.Setenv("Test", "test"))


	// 获取当前目录
	fmt.Println(os.Getwd())   //E:\gocode\src\myapiserver
	//改变当前工作目录
	//os.Chdir("C:/")
	//fmt.Println(os.Getwd())

	// 获取所有环境变量
	//fmt.Println(strings.Join(os.Environ(), "\r\n"))

	//把字符串中带${var}或$var替换成指定指符串

	fmt.Println(os.Expand("${1} ${2} ${3}", func(s string) string {
		mapp := map[string]string{
			"1":"111",
			"2":"222",
			"3":"333",
		}
		return mapp[s]
	})) //111 222 333
}

func osprocess()  {
	//设置新进程的属性
	attr := &os.ProcAttr{
		//files指定新进程继承的活动文件对象
		//前三个分别为，标准输入、标准输出、标准错误输出
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		//新进程的环境变量
		Env: os.Environ(),
	}

	//win下通过记事本打开1.txt
	//开始一个新进程
	p, err := os.StartProcess("C:\\Windows\\System32\\notepad.exe", []string{"C:\\Windows\\System32\\notepad.exe", "D:\\1.txt"}, attr)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(p)
	fmt.Println("进程ID：", p.Pid)

	//通过进程ID查找进程
	p2, _ := os.FindProcess(p.Pid)
	fmt.Println(p2)

	//等待10秒，执行函数
	time.AfterFunc(time.Second*10, func() {
		//向p进程发送退出信号
		p.Signal(os.Kill)
	})

	//等待进程p的退出，返回进程状态
	ps, _ := p.Wait()
	fmt.Println(ps.String())

	fmt.Println(os.Environ())
}

func main()  {
	// os1()

	osprocess()
}