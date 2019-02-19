package main

import (
	"net/http"
	"log"
	"fmt"
	"reflect"
	"bytes"
)

func main()  {

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