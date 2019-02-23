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
如何自己实现一个。它的功能是从流中过滤掉非字母字符
组合多个 Reader，目的是重用和屏蔽下层实现的复杂度
这样做的另一个优点是 alphaReader 能够从任何 Reader 实现中读取。
例如，以下代码展示了 alphaReader 如何与 os.File 结合以过滤掉文件中的非字母字符：
 */

 type alpahReader struct {
	 // alphaReader 里组合了标准库的 io.Reader
	 reader io.Reader
 }

 // 创建一个实例
func newAlphaReader(r io.Reader) *alpahReader  {
	return &alpahReader{reader: r}
}

// 过滤函数
func alpha(r byte) byte  {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}

	return 0
}

func (a *alpahReader) Read(p []byte)(int, error)  {
	// 这行代码调用的就是 io.Reader
	n, err := a.reader.Read(p)
	if err != nil {
		return n, err
	}
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := alpha(p[i]); char != 0 {
			buf[i] = char
		}
	}

	copy(p, buf)
	return n, nil

}

func init1()  {
	//  使用实现了标准库 io.Reader 接口的 strings.Reader 作为实现
	reader := newAlphaReader(strings.NewReader("Hello! It's 9am, where is the sun?"))
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

func init2()  {
	// file 也实现了 io.Reader
	file, err := os.Open("E:\\gocode\\src\\myapiserver\\demo-io\\reader.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// 任何实现了 io.Reader 的类型都可以传入 newAlphaReader
	// 至于具体如何读取文件，那是标准库已经实现了的，我们不用再做一遍，达到了重用的目的
	reader := newAlphaReader(file)
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println()
}

func main()  {
	init2()
}
