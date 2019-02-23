package main

import (
	"strings"
	"io"
	"fmt"
	"os"
)

/**
src/io/io.go中定义。后边凡是涉及到io相关操作的，基本上都实现了这两个接口
type Reader interface {
    Read(p []byte) (n int, err error) //Read() 方法有两个返回值，一个是读取到的字节数，一个是发生错误时的错误。
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
 */

/**
使用Reader
利用 Reader 可以很容易地进行流式数据传输。
Reader 方法内部是被循环调用的，每次迭代，它会从数据源读取一块数据放入缓冲区 p （即 Read 的参数 p）中，直到返回 io.EOF 错误时停止。
 */

func strReader()  {
	reader := strings.NewReader("Clear is better than clever")
	// make(切换类型， 长度， 容量)
	p := make([]byte, 4) //缓冲区大小

	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF:", n)
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		////从0开始，取n个元素，容量也是n， 常用
		fmt.Println(n, string(p[:n]))
	}


	/**
	4 Clea
	4 r is
	4  bet
	4 ter
	4 than
	4  cle
	3 ver
	EOF: 0
	可以看到，最后一次返回的 n 值有可能小于缓冲区大小。
	 */
}

func strReader2()  {
	str := strings.NewReader("I love China")
	p := make([]byte, 3)

	for {
		n, err := str.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF", n)
				break
			}

			fmt.Println(err)
			os.Exit(1)
		}

		// 切片转字符串
		fmt.Println(n, string(p[:n]))
	}
}

/**
如何自己实现一个。它的功能是从流中过滤掉非字母字符
 */

 type alpahReaderOne struct {
 	// 资源
 	src string

 	//当前位置
 	cur int
 }

 // 创建一个实例
func newAlphaReaderOne(s string) *alpahReaderOne  {
	return &alpahReaderOne{src: s}
}

// 过滤函数
func alpha1(r byte) byte  {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}

	return 0
}

func (a *alpahReaderOne) Read(p []byte)(int, error)  {

	// 当前位置 >= 字符串长度 说明已经读取到结尾 返回 EOF
	if a.cur >= len(a.src) {
		return 0, io.EOF
	}

	// x 是剩余未读取的长度
	x := len(a.src) - a.cur
	n, bound := 0, 0
	if x >= len(p) {
		// 剩余长度超过缓冲区大小，说明本次可完全填满缓冲区
		bound = len(p)
	} else if x < len(p) {
		// 剩余长度小于缓冲区大小，使用剩余长度输出，缓冲区不补满
		bound = x
	}

	buf := make([]byte, bound)
	for n < bound {
		// 每次读取一个字节，执行过滤函数
		if char := alpha1(a.src[a.cur]); char != 0 {
			buf[n] = char
		}
		n++
		a.cur++
	}
	// 将处理后得到的 buf 内容复制到 p 中
	copy(p, buf)
	return n, nil

}

func init()  {
	reader := newAlphaReaderOne("Hello! It's 9am, where is the sun?")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println() //Hello  It s  am  where is the sun
}

func main()  {
	// strReader2()
}
