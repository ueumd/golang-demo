package main

import (
	"bytes"
	"os"
)

/**
io.Copy() 可以轻松地将数据从一个 Reader 拷贝到另一个 Writer。
它抽象出 for 循环模式（我们上面已经实现了）并正确处理 io.EOF 和 字节计数。
下面是我们之前实现的简化版本：
 */

func main()  {
	proverbs := new(bytes.Buffer)
	proverbs.WriteString("Channels orchestrate mutexes serialize\n")
	proverbs.WriteString("Cgo is not Go\n")
	proverbs.WriteString("Errors are values\n")
	proverbs.WriteString("Don't panic\n")

	file, err := os.Create("./proverbs.txt")
}