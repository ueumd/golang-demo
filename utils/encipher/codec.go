/*
 * 将数据对称加密再编码，逆操作为解码再解密
 * cipher为go官方标准库采用的包名
 *
 * wencan
 * 2018-08-30
 */

package encipher

import (
	"encoding/base64"
	"log"
)

// Decipher 解码并解密
type Decipher interface {
	Decode(signed string) (string, error)
}

// Encipher 加密并编码
type Encipher interface {
	Encode(src string) (string, error)
}

// Cipher 加解密并编解码
type Cipher interface {
	Encipher
	Decipher
}

// NewAesEncipher 创建AES加解密编码器
func NewAesEncipher(privateKey []byte) *AesCodec {
	// 确保cap(key) == 32
	key := make([]byte, 32)
	copy(key, privateKey)

	return &AesCodec{
		privateKey: key,
	}
}

// AesCodec AES编解码器
type AesCodec struct {
	privateKey []byte
}

// Encode 加密并编码
func (codec *AesCodec) Encode(src string) (string, error) {
	// 加密
	crypted, err := AesEncrypt([]byte(src), codec.privateKey)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// 编码
	encoded := base64.RawURLEncoding.EncodeToString(crypted)
	return encoded, nil
}

// Decode 解码并解密
func (codec *AesCodec) Decode(signed string) (string, error) {
	// 解码
	decoded, err := base64.RawURLEncoding.DecodeString(signed)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// 解密
	src, err := AesDecrypt(decoded, codec.privateKey)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(src), nil
}
