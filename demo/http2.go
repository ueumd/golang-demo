package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"os"
	"fmt"
)
const url  = "https://broqiang.com"

// get 方式init1 init2
func ErrPrint(err error)  {
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}

func init1()  {
	// 方式一，直接通过 Get 函数
	resp, err := http.Get(url)
	ErrPrint(err)
	defer resp.Body.Close()

	// 拿到数据
	bytes, err := ioutil.ReadAll(resp.Body)
	ErrPrint(err)
	fmt.Printf("%s", bytes)
}

func init2()  {
	client := &http.Client{} // new(http.Client)

	resp, _ := client.Get(url)

	defer  resp.Body.Close()

	res, _  := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", res)
}


/**
post
https://learnku.com/articles/23430/golang-learning-notes-1-http-client-foundation
(url string, contentType string, body io.Reader) (resp *Response, err error)

contentType (4种):
	application/x-www-form-urlencoded 不设置 enctype 属性的原生 form 表单提交方式。
	multipart/form-data 上传文件时的数据提交方式，相当于 form 表单的 enctype 等于 multipart/form-data 。
	application/json 用来告诉服务端消息主体是序列化后的 JSON 字符串。
	text/xml 它是一种使用 HTTP 作为传输协议，XML 作为编码方式的远程调用规范，和 json 作用类型。
 */
func main()  {
	// init1()
	init2()
}
