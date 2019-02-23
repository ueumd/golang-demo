package main

import (
	"bytes"
	"os"
	"fmt"
)

/**
标准库提供了许多已经实现了 io.Writer 的类型。
下面是一个简单的例子，它使用 bytes.Buffer 类型作为 io.Writer 将数据写入内存缓冲区。
 */
func write1()  {
	proverbs := []string {
		"Channels orchestrate mutexes serialize",
		"Cgo is not Go",
		"Errors are values",
		"Don't panic",
	}

	var writer bytes.Buffer
	for _, p := range proverbs {
		n, err := writer.Write([]byte(p))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if n != len(p) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}
	fmt.Println(writer.String())
}

func main()  {
	write1()
}