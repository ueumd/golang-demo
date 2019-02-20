package main

import (
		"io/ioutil"
	"log"
	"os"
	"fmt"
	"strings"
	"io"
	"net/http"
	"reflect"
	"bytes"
)
// https://learnku.com/articles/23430/golang-learning-notes-1-http-client-foundation
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

 func post1 () {
 	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
 		if req.Method == "POST" {
 			fmt.Fprintf(w, "Hello333 %s\n", req.FormValue("name"))
			return
		}
 		http.NotFound(w, req)
	})
 	log.Fatalf("%v", http.ListenAndServe("localhost:8080", nil))
 }



func DataPrint(body io.ReadCloser) {
	// 拿到数据
	bytes, err := ioutil.ReadAll(body)
	ErrPrint(err)

	// 这里要格式化再输出，因为 ReadAll 返回的是字节切片
	fmt.Printf("%s",bytes)
}

const POST_URL  = "http://localhost:8081/hello"

func post2()  {
	resp, _ := http.Post(
		POST_URL,
		"application/x-www-form-urlencoded",
		strings.NewReader("name=Hello World"),
		)

	defer resp.Body.Close()

	DataPrint(resp.Body)
}

func post3()  {
	// 方式二，通过 client 结构体中的 Post 方法
	client := &http.Client{}
	resp, _ := client.Post(
		url,
		"application/x-www-form-urlencoded",
		strings.NewReader("name=New Bro Qiang"),
	)

	defer resp.Body.Close()

	DataPrint(resp.Body)
}

//PostFomr

func postForm()  {

	const url = "http://localhost:8081/form"
	// 方法一：PostForm 函数
	data := map[string][]string {
		"name" : {"John"},
		"gender" : {"male"},
	}

	resp, err := http.PostForm(url, data)
	ErrPrint(err)
	defer resp.Body.Close()

	DataPrint(resp.Body)



	// 方法二：client 结构体的 PostForm 方法
	//client := &http.Client{}
	//resp, err = client.PostForm(url, data)
	//ErrPrint(err)
	//defer resp.Body.Close()
	//
	//DataPrint(resp.Body)
}


func main()  {
	// init1()
	// init2()

	//post1()
	//post2()
	// post3()

	postForm()
}

func httptest()  {
	resp, err := http.Get("http://www.baidu.com")

	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	headers := resp.Header
	for k, v := range headers {
		fmt.Printf("k=%v, v=%v\n", k, v)
	}

	fmt.Printf("resp status %s, statusCode: %d\n", resp.Status, resp.StatusCode) // resp status 200 OK, statusCode: 200
	fmt.Printf("reps Proto %s\n", resp.Proto)									// reps Proto HTTP/1.1
	fmt.Printf("resp content length %d\n", resp.ContentLength)					// resp content length -1
	fmt.Printf("resp transer encodeing %v\n", resp.TransferEncoding)				// resp transer encodeing [chunked]
	fmt.Printf("resp Uncompressed %t\n", resp.Uncompressed)						// resp Uncompressed true
	fmt.Println(reflect.TypeOf(resp.Body))												// *http.gzipReader

	buf := bytes.NewBuffer(make([]byte, 0, 512))

	length, _ := buf.ReadFrom(resp.Body)

	fmt.Println(len(buf.Bytes()))   // 153473
	fmt.Println(length)				// 153473
	fmt.Println(string(buf.Bytes())) // HTML源码

}