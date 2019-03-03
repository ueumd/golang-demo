package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

/**
密码学简介与Golang的加密库Crypto的使用
https://blog.yumaojun.net/2017/02/19/go-crypto/

crypto		ˈkriptō		加密
cipher 		ˈsīfər 	   暗号
encrypt  	enˈkript  加密
decrypt		diˈkript  解密


对称加密算法: 加密和解密均采用同一把秘密钥匙。
非对称加密算法: 有2把密钥,公钥和私钥, 公钥加密, 私钥解密。

对称加密算法：
AES(Advanced Encryption Standard): 高级加密标准，是下一代的加密算法标准，速度快，安全级别高；
 */

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

// 加密
func encrypt(plainText , keyText string) (cipherByte []byte, err error)  {
	// 转换字节数据方便加密
	plainByte := []byte(plainText)
	keyBte :=[]byte(keyText)

	// 创建加密算法aes
	c, err := aes.NewCipher(keyBte)
	if err != nil {
		return  nil, err
	}

	// 加密字符串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	cipherByte = make([]byte, len(plainText))
	cfb.XORKeyStream(cipherByte, plainByte)
	return
}

// 解密
func decrypt(cipherByte []byte, keyText string)(plainText string, err error)  {
	// 转换成字节数据, 方便加密
	keyByte := []byte(keyText)

	c, err := aes.NewCipher(keyByte)

	if err != nil {
		return "", err
	}

	// 解密字符串
	cfbdec  := cipher.NewCFBDecrypter(c, commonIV)
	plainByte := make([]byte, len(cipherByte))

	cfbdec.XORKeyStream(plainByte, cipherByte)

	plainText = string(plainByte)
	return 

}

func main()  {
	plain := "The text need to be encrypt."

	// AES 规定有3种长度的key: 16, 24, 32分别对应AES-128, AES-192, or AES-256
	key := "abcdefgehjhijkmlkjjwwoew"

	// 加密
	cipherByte , err := encrypt(plain, key)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s ==> %x\n", plain, cipherByte)

	// 解密
	plainText , err := decrypt(cipherByte, key)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%x ==> %s\n", cipherByte, plainText)
}