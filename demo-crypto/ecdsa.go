package main

import (
"crypto/ecdsa"
"crypto/rand"
"fmt"
"crypto/elliptic"
"log"
)

/*
生成公钥和私钥的时候，可以选择以下几种
elliptic.P224()
elliptic.P256()
elliptic.P384()
elliptic.P521()
*/
func main() {
	// 生成公钥和私钥
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalln(err)
	}

	// 公钥是存在在私钥中的，从私钥中读取公钥
	publicKey := &privateKey.PublicKey

	fmt.Println(privateKey)
	fmt.Println(publicKey)

	message := []byte("hello,dsa签名")
	// 进入签名操作
	r, s, _ := ecdsa.Sign(rand.Reader, privateKey, message)

	// 进入验证
	flag := ecdsa.Verify(publicKey, message, r, s)
	if flag {
		fmt.Println("数据未被修改")
	} else {
		fmt.Println("数据被修改")
	}

	flag = ecdsa.Verify(publicKey, []byte("hello"), r, s)
	if flag {
		fmt.Println("数据未被修改")
	} else {
		fmt.Println("数据被修改")
	}
}