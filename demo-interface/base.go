package main

import "fmt"

// File类
type File struct {

}

// File类方法
func (f *File) Read(buf []byte) (n int, err error)  {
	return
}

func (f *File) Write(buf []byte) (n int, err error)  {
	return
}

func (f *File) Seek(off int64, whence int) (pos int64, err error)  {
	return
}

func (f *File) Close() error  {
	return nil
}

//接口1：IFile
type IFile interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Seek(off int64, whence int) (pos int64, err error)
	Close() error
}
type IReader interface {
	Read(buf []byte) (n int, err error)
}


func main()  {
	//接口赋值,File类实现了IFile和IReader接口，即接口所包含的所有方法

	var f1 IFile = new(File)
	var f2 IReader = new(File)

	fmt.Println(f1)
	fmt.Println(f2)


}