package main

import (
	"os"
	"fmt"
	"time"
)

/**
func Create(name string) (file *File, err error)  // Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）
func Open(name string) (file *File, err error)   // Open打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据；对应的文件描述符具有O_RDONLY模式
func (f *File) Stat() (fi FileInfo, err error)   // Stat返回描述文件f的FileInfo类型值
func (f *File) Readdir(n int) (fi []FileInfo, err error)  // Readdir读取目录f的内容，返回一个有n个成员的[]FileInfo，这些FileInfo是被Lstat返回的，采用目录顺序
func (f *File) Read(b []byte) (n int, err error)          // Read方法从f中读取最多len(b)字节数据并写入b
func (f *File) WriteString(s string) (ret int, err error) // 向文件中写入字符串
func (f *File) Sync() (err error)           // Sync递交文件的当前内容进行稳定的存储。一般来说，这表示将文件系统的最近写入的数据在内存中的拷贝刷新到硬盘中稳定保存
func (f *File) Close() error                // Close关闭文件f，使文件不能用于读写
 */
func f1()  {
	// 获取当前目录
	dir, err := os.Getwd()
	fmt.Println(dir, err)

	file := dir +"/demo-io-os/new.txt"
	var fh *os.File

	/**
	FileInfo 用来描述一个文件对象
	type FileInfo interface {
		Name() string       // base name of the file
		Size() int64        // length in bytes for regular files; system-dependent for others
		Mode() FileMode     // file mode bits
		ModTime() time.Time // modification time
		IsDir() bool        // abbreviation for Mode().IsDir()
		Sys() interface{}   // underlying data source (can return nil)
	}
	 */

	fi, _ := os.Stat(file)

	if fi == nil {
		fh, _ = os.Create(file) // 文件不存在就创建
	} else {

		//文件名 nex.txt
		fmt.Println(fi.Name())
		//文件大小 79 byte
		fmt.Println(fi.Size())
		//文件权限 -rw-rw-rw-
		fmt.Println(fi.Mode())
		//文件修改时间 2019-02-24 19:19:20.8650145 +0800 CST
		fmt.Println(fi.ModTime())
		//是否是目录 false
		fmt.Println(fi.IsDir())


		fh, _ = os.OpenFile(file, os.O_RDWR, 0666) //存在就打开
	}

	w := []byte("hello go language " + time.Now().String())

	n , err := fh.Write(w)
	fmt.Println(n, err)

	fh.WriteString(" go go go") //向文文件写入字符串
	// 设置下次读写位置
	/**
	offset 文件指针的位置
	whence
		0 代表相对文件开始的位置，
		1 代表相对当前位置，
	    2 代表相对文件结尾的位置
	ret 返回的是现在指针的位置
	 */
	ret, err := fh.Seek(0, 0) //
	fmt.Println("当前指针位置：", ret, err)

	b := make([]byte, 10)  //设置读取的字节数
	n, err = fh.Read(b)   // 读取并写入到b
	fmt.Println(n, err, string(b))  // 10 <nil> hello go l




	// 当前工程目录下 demo-io-os
	// os.Create("./demo-io-os/hello.txt")
}

func main()  {
	f1()
}
