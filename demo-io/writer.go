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

//自己实现一个 Writer
type chanWriter struct {
	// ch 实际上就是目标资源
	ch chan byte
}

func newChanWriter() *chanWriter  {
	return &chanWriter{make(chan byte, 1024)}
}

func (w *chanWriter) Chan() <-chan byte {
	return w.ch
}

func (w *chanWriter) Write(p []byte) (int, error)  {
	n := 0
	for _, b := range p {
		w.ch <- b
		n ++
	}
	return n, nil
}

func (w *chanWriter) Close() error  {
	close(w.ch)
	return nil
}

func writer2()  {
	writer := newChanWriter()
	go func() {
		defer writer.Close()
		writer.Write([]byte("Stream"))
		writer.Write([]byte("me!"))
	}()

	for c := range writer.Chan() {
		fmt.Printf("%c", c)
	}
	fmt.Println()
}

func main()  {
//	write1()
	writer2()
}